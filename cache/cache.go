/*
	in-memory key-value store
	(c) 2013 Cergoo
	under terms of ISC license
*/
package cache

/*
TODO :
 - memory limiter implement
*/

import (
	"encoding/gob"
	"fmt"
	"github.com/Cergoo/gol/counter"
	"io"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	growcountgo    = 16
	initbucketscap = 10
)

type (
	t_item struct {
		r   uint8
		Key string
		Val interface{}
	}

	t_bucket struct {
		items []*t_item
		sync.RWMutex
	}

	t_hash struct {
		ht   []*t_bucket         // hash table
		hash func([]byte) uint32 // hash function
	}

	t_cache struct {
		t                      *t_hash           // hash table
		janitor_duration       time.Duration     // janitor_duration janitor
		janitor_ifreadthenlive bool              // janitor_ifreadthenlive enable
		count                  counter.T_counter // limit items count in cache, 0 - unlimit
		growlock               uint32
		janitor_ch             chan<- *TCortege
		chan_stop              chan bool
	}

	// for finalizer
	tfinalize struct {
		*t_cache
	}
)

// Key to ID function
func (t *t_hash) keyToID(key []byte) uint32 {
	return t.hash(key) % uint32(len(t.ht))
}

/*
	Constructor new cache:
	hash                   - hash function
	janitor_ifreadthenlive - if item read then item live
	janitor_duration       - time to clear items, if 0 then never
	janitor_ch             - chanel to send removed items
*/
func New(
	hash func([]byte) uint32,
	janitor_ifreadthenlive bool,
	janitor_duration time.Duration,
	janitor_ch chan<- *TCortege) Cache {

	ht := new(t_hash)
	ht.ht = make([]*t_bucket, 1024)
	ht.hash = hash
	for i := range ht.ht {
		ht.ht[i] = &t_bucket{
			items: make([]*t_item, 0, initbucketscap),
		}
	}

	t := &t_cache{
		janitor_duration:       janitor_duration,
		janitor_ifreadthenlive: janitor_ifreadthenlive,
		t:          ht,
		janitor_ch: janitor_ch,
	}

	if janitor_duration > 0 {
		finalized := &tfinalize{t}
		t.chan_stop = make(chan bool)
		go t.janitor(t.chan_stop)
		runtime.SetFinalizer(finalized, stop)
		return finalized
	}

	return t
}

func stop(t *tfinalize) {
	close(t.chan_stop)
}

// Point to countern
func (t *t_cache) Len() I_counter {
	return &t.count
}

// Get statistics records Bucket
func (t *t_cache) GetBucketsStat() (countitem uint64, countbucket uint32, stat [][2]int) {
	var i int
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	tmp1 := make(map[int]int)
	for _, bucket := range ht.ht {
		bucket.RLock()
		tmp1[len(bucket.items)]++
		bucket.RUnlock()
		countbucket++
	}

	// sort
	tmp2 := make([]int, 0, len(tmp1))
	for i = range tmp1 {
		tmp2 = append(tmp2, i)
	}
	sort.Ints(tmp2)
	stat = make([][2]int, 0, len(tmp1))
	for i = range tmp2 {
		stat = append(stat, [2]int{tmp2[i], tmp1[tmp2[i]]})
	}
	countitem = t.count.Get()
	return
}

// Get item value or nil
func (t *t_cache) Get(key string) (val interface{}) {
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht.ht[ht.keyToID([]byte(key))]
	bucket.RLock()
	for _, v := range bucket.items {
		if v.Key == key {
			if t.janitor_ifreadthenlive {
				v.r = 1
			}
			val = v.Val
			break
		}
	}
	bucket.RUnlock()
	return
}

// Set mode: onlyUpdate, onlyInsert, updateOrInsert
func (t *t_cache) Set(key string, val interface{}, mode uint8) (rval interface{}, actionResult uint8) {
	var (
		v *t_item
	)
	rval = val
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht.ht[ht.keyToID([]byte(key))]
	bucket.Lock()

	// Update
	for _, v = range bucket.items {
		if v.Key == key {
			actionResult = ResultExist
			if mode == OnlyInsert {
				rval = v.Val
				bucket.Unlock()
				return
			}
			v.Val = val
			bucket.Unlock()
			return
		}
	}

	// Add
	if mode != OnlyUpdate && t.count.Check() {
		lenbucket := len(bucket.items)
		bucket.items = append(bucket.items, &t_item{Key: key, Val: val, r: 1})
		bucket.Unlock()
		if lenbucket > growcountgo && atomic.CompareAndSwapUint32(&t.growlock, 0, 2) {
			go t.grow()
		}
		actionResult = ResultAdd
		t.count.Inc()
		return
	}

	bucket.Unlock()
	actionResult = ResultNoExistNoAdd
	return
}

// Get and Delete item key
func (t *t_cache) Del(key string) (val interface{}) {
	var endi int
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht.ht[ht.keyToID([]byte(key))]
	bucket.Lock()
	for i, v := range bucket.items {
		if v.Key == key {
			val = v.Val
			endi = len(bucket.items) - 1
			bucket.items[i], bucket.items[endi] = bucket.items[endi], nil
			bucket.items = bucket.items[:endi]
			t.count.Dec()
			break
		}
	}
	bucket.Unlock()
	return
}

// Incremet and Decrement any type, return modified value or nil
func (t *t_cache) Inc(key string, n float64) interface{} {
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht.ht[ht.keyToID([]byte(key))]

	bucket.Lock()
	for _, v := range bucket.items {
		if v.Key == key {
			switch value := v.Val.(type) {
			case int:
				v.Val = value + int(n)
			case int8:
				v.Val = value + int8(n)
			case int16:
				v.Val = value + int16(n)
			case int32:
				v.Val = value + int32(n)
			case int64:
				v.Val = value + int64(n)
			case uint:
				v.Val = value + uint(n)
			case uint8:
				v.Val = value + uint8(n)
			case uint16:
				v.Val = value + uint16(n)
			case uint32:
				v.Val = value + uint32(n)
			case uint64:
				v.Val = value + uint64(n)
			case float32:
				v.Val = value + float32(n)
			case float64:
				v.Val = value + n
			}

			v.r = 1
			bucket.Unlock()
			return v.Val
		}
	}
	bucket.Unlock()
	return nil
}

// Delete all items
func (t *t_cache) DelAll() {
	var i int
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)))).ht
	for _, bucket := range ht {
		bucket.Lock()
		t.count.Add(int64(-(len(bucket.items) - 1)))
		for i = range bucket.items {
			bucket.items[i] = nil
		}
		bucket.Unlock()
	}
}

// Interface range function, for breack range close a channel
func (t *t_cache) Range(ch chan<- *TCortege) {
	go t.rangeCache(ch)
}

/*
    non interface range function,
	for breack range close a channel
	Use buffer for unlock bucket
*/
func (t *t_cache) rangeCache(ch chan<- *TCortege) {
	var (
		i   int
		buf []*TCortege
		v   *TCortege
	)
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)))).ht
	for _, bucket := range ht {
		bucket.RLock()
		for i = range bucket.items {
			buf = append(buf, &TCortege{Key: bucket.items[i].Key, Val: bucket.items[i].Val})
		}
		bucket.RUnlock()
		for _, v = range buf {
			ch <- v
		}
		buf = buf[:0]
	}
	defer nopanic()
}

func nopanic() {
	recover()
}

// Write the cache's items (using Gob) to an io.Writer.
func (t *t_cache) Save(w io.Writer) (err error) {
	var (
		item   *t_item
		bucket *t_bucket
	)
	enc := gob.NewEncoder(w)
	ht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))

	defer func() {
		if x := recover(); x != nil {
			for _, bucket = range ht.ht {
				bucket.RUnlock()
			}
			err = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	for _, bucket = range ht.ht {
		bucket.RLock()
		for _, item = range bucket.items {
			gob.Register(item.Val)
			err = enc.Encode(*item)
			if err != nil {
				return
			}
		}
		bucket.RUnlock()
	}
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
func (t *t_cache) SaveFile(fname string) error {
	fp, err := os.Create(fname)
	if err != nil {
		return err
	}
	err = t.Save(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
func (t *t_cache) Load(r io.Reader) error {
	var (
		err  error
		item t_item
	)
	dec := gob.NewDecoder(r)
	for err = dec.Decode(&item); err == nil; err = dec.Decode(&item) {
		t.Set(item.Key, item.Val, UpdateOrInsert)
	}
	return err
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
func (t *t_cache) LoadFile(fname string) error {
	fp, err := os.Open(fname)
	if err != nil {
		return err
	}
	err = t.Load(fp)
	if err != nil && err.Error() != "EOF" {
		fp.Close()
		return err
	}
	return fp.Close()
}

/*
	Grow hash table procedure
	old hash table (n buckets)	->	new hash table	(n*2 buckets)
	example:
	0    0       0Lock   0       0       0       0
	1    1       1       1Lock   1       1       1
	2    2       2       2       2Lock   2       2
	3    3       3       3       3       3Lock   3
	     4Lock   4Lock   4       4       4       4
	     5Lock   5Lock   5Lock   5       5       5
	     6Lock   6Lock   6Lock   6Lock   6       6
	     7Lock   7Lock   7Lock   7Lock   7Lock   7
*/
func (t *t_cache) grow() {
	var (
		i, j int
		val  *t_item
	)

	oldht := (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	oldlen := len(oldht.ht)
	newlen := oldlen << 1

	if uint64(newlen) >= math.MaxUint32 {
		// grow only singleton runed
		atomic.StoreUint32(&t.growlock, 0)
		return
	}

	newht := new(t_hash)
	newht.ht = make([]*t_bucket, newlen)
	copy(newht.ht, oldht.ht) // possible since it links
	for i = oldlen; i < newlen; i++ {
		newht.ht[i] = &t_bucket{}
		newht.ht[i].Lock()
		j++
	}
	newht.hash = oldht.hash
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)), unsafe.Pointer(newht))
	oldht = nil
	// rehash

	j = oldlen
	for i = 0; i < oldlen; i++ {
		itemsold := make([]*t_item, 0, initbucketscap)
		itemsnew := make([]*t_item, 0, initbucketscap)
		newht.ht[i].Lock()
		for _, val = range newht.ht[i].items {
			if newht.keyToID([]byte(val.Key)) == uint32(i) {
				itemsold = append(itemsold, val)
			} else {
				itemsnew = append(itemsnew, val)
			}
		}

		newht.ht[j].items = itemsnew
		newht.ht[j].Unlock()
		newht.ht[i].items = itemsold
		newht.ht[i].Unlock()
		j++
	}

	// grow only singleton runed
	atomic.StoreUint32(&t.growlock, 0)
}

func (t *t_cache) janitor(stop <-chan bool) {
	var (
		i, lenbucket int
		ht           *t_hash
		bucket       *t_bucket
		count_del    uint32
		buf          []*TCortege
		v            *TCortege
	)
	tick := time.Tick(t.janitor_duration)
	for {
		select {
		case <-stop:
			return
		case <-tick:
			count_del = 0
			ht = (*t_hash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
			for _, bucket = range ht.ht {
				bucket.Lock()
				lenbucket = len(bucket.items)
				for i = 0; i < lenbucket; {
					if bucket.items[i].r == 0 {
						if t.janitor_ch != nil {
							buf = append(buf, &TCortege{Key: bucket.items[i].Key, Val: bucket.items[i].Val})
						}
						lenbucket--
						bucket.items[i], bucket.items[lenbucket] = bucket.items[lenbucket], nil
						count_del++
					} else {
						bucket.items[i].r--
						i++
					}
				}
				bucket.items = bucket.items[:lenbucket]
				bucket.Unlock()
				for _, v = range buf {
					t.janitor_ch <- v
				}
				buf = buf[:0]
			}
			t.count.Add(int64(-count_del))
		}
	}
}
