// (c) 2013 Cergoo
// under terms of ISC license

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
	tItem struct {
		r   uint8
		Key string
		Val interface{}
	}

	tBucket struct {
		items []*tItem
		sync.RWMutex
	}

	tHash struct {
		ht   []*tBucket          // hash table
		hash func([]byte) uint32 // hash function
	}

	tCache struct {
		t                     *tHash           // hash table
		janitorDuration       time.Duration    // janitorDuration janitor
		janitorIfReadThenLive bool             // janitorIfReadThenLive enable
		count                 counter.TCounter // limit items count in cache, 0 - unlimit
		growlock              uint32
		janitorCh             chan<- *TCortege
		stopCh                chan bool
	}

	// for finalizer
	tfinalize struct {
		*tCache
	}
)

// keyToID Key to ID function
func (t *tHash) keyToID(key []byte) uint32 {
	return t.hash(key) % uint32(len(t.ht))
}

/*
New constructor of a new cache:
    hash                   - hash function;
    janitorIfReadThenLive - if item read then item live;
    janitorDuration       - time to clear items, if 0 then never;
    janitorCh             - chanel to send removed items.
*/
func New(
	hash func([]byte) uint32,
	janitorIfReadThenLive bool,
	janitorDuration time.Duration,
	janitorCh chan<- *TCortege) Cache {

	ht := new(tHash)
	ht.ht = make([]*tBucket, 1024)
	ht.hash = hash
	for i := range ht.ht {
		ht.ht[i] = &tBucket{
			items: make([]*tItem, 0, initbucketscap),
		}
	}

	t := &tCache{
		janitorDuration:       janitorDuration,
		janitorIfReadThenLive: janitorIfReadThenLive,
		janitorCh:             janitorCh,
		t:                     ht,
	}

	if janitorDuration > 0 {
		finalized := &tfinalize{t}
		t.stopCh = make(chan bool)
		go t.janitor(t.stopCh)
		runtime.SetFinalizer(finalized, stop)
		return finalized
	}

	return t
}

func stop(t *tfinalize) {
	close(t.stopCh)
}

// Len get counter cache
func (t *tCache) Len() ICounter {
	return &t.count
}

// GetBucketsStat get statistics of a buckets
func (t *tCache) GetBucketsStat() (countitem uint64, countbucket uint32, stat [][2]int) {
	var i int
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
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

// Get get item value or nil
func (t *tCache) Get(key string) (val interface{}) {
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht.ht[ht.keyToID([]byte(key))]
	bucket.RLock()
	for _, v := range bucket.items {
		if v.Key == key {
			if t.janitorIfReadThenLive {
				v.r = 1
			}
			val = v.Val
			break
		}
	}
	bucket.RUnlock()
	return
}

/*
Set:
    key  - record key
    val  - record val
    live - record time live, if 255 then no janitor remov, removed only when the user
    mode - set mode: onlyUpdate, onlyInsert, updateOrInsert
*/
func (t *tCache) Set(key string, val interface{}, live uint8, mode uint8) (rval interface{}, actionResult uint8) {
	var (
		v *tItem
	)
	rval = val
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
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
			v.r = live
			bucket.Unlock()
			return
		}
	}

	// Add
	if mode != OnlyUpdate && t.count.Check() {
		lenbucket := len(bucket.items)
		bucket.items = append(bucket.items, &tItem{Key: key, Val: val, r: live})
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

// Del get and delete item
func (t *tCache) Del(key string) (val interface{}) {
	var endi int
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
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

// Inc incremet/decrement any numeric type, return modified value or nil
func (t *tCache) Inc(key string, n float64) interface{} {
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
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

// DelAll delete all items
func (t *tCache) DelAll() {
	var i int
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)))).ht
	for _, bucket := range ht {
		bucket.Lock()
		t.count.Add(int64(-(len(bucket.items) - 1)))
		for i = range bucket.items {
			bucket.items[i] = nil
		}
		bucket.Unlock()
	}
}

// Range interface range function, for breack range close a channel
func (t *tCache) Range(ch chan<- *TCortege) {
	go t.rangeCache(ch)
}

/*
    non interface range function,
	for breack range close a channel

	Use buffer for unlock bucket
*/
func (t *tCache) rangeCache(ch chan<- *TCortege) {
	var (
		i   int
		buf []*TCortege
		v   *TCortege
	)
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)))).ht
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

// Save write the cache's items (using Gob) to an io.Writer.
func (t *tCache) Save(w io.Writer) (err error) {
	var (
		item   *tItem
		bucket *tBucket
	)
	enc := gob.NewEncoder(w)
	ht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))

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

// SaveFile save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
func (t *tCache) SaveFile(fname string) error {
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

// Load add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
func (t *tCache) Load(r io.Reader) error {
	var (
		err  error
		item tItem
	)
	dec := gob.NewDecoder(r)
	for err = dec.Decode(&item); err == nil; err = dec.Decode(&item) {
		t.Set(item.Key, item.Val, item.r, UpdateOrInsert)
	}
	return err
}

// LoadFile load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
func (t *tCache) LoadFile(fname string) error {
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
func (t *tCache) grow() {
	var (
		i, j int
		val  *tItem
	)

	oldht := (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	oldlen := len(oldht.ht)
	newlen := oldlen << 1

	if uint64(newlen) >= math.MaxUint32 {
		// grow only singleton runed
		atomic.StoreUint32(&t.growlock, 0)
		return
	}

	newht := new(tHash)
	newht.ht = make([]*tBucket, newlen)
	copy(newht.ht, oldht.ht) // possible since it links
	for i = oldlen; i < newlen; i++ {
		newht.ht[i] = &tBucket{}
		newht.ht[i].Lock()
		j++
	}
	newht.hash = oldht.hash
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)), unsafe.Pointer(newht))
	oldht = nil
	// rehash

	j = oldlen
	for i = 0; i < oldlen; i++ {
		itemsold := make([]*tItem, 0, initbucketscap)
		itemsnew := make([]*tItem, 0, initbucketscap)
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

func (t *tCache) janitor(stop <-chan bool) {
	var (
		i, lenbucket int
		ht           *tHash
		bucket       *tBucket
		count_del    uint32
		buf          []*TCortege
		v            *TCortege
	)
	tick := time.Tick(t.janitorDuration)
	for {
		select {
		case <-stop:
			return
		case <-tick:
			count_del = 0
			ht = (*tHash)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
			for _, bucket = range ht.ht {
				bucket.Lock()
				lenbucket = len(bucket.items)
				for i = 0; i < lenbucket; {
					if bucket.items[i].r == 0 {
						if t.janitorCh != nil {
							buf = append(buf, &TCortege{Key: bucket.items[i].Key, Val: bucket.items[i].Val})
						}
						lenbucket--
						bucket.items[i], bucket.items[lenbucket] = bucket.items[lenbucket], nil
						count_del++
					} else {
						if bucket.items[i].r < 255 {
							bucket.items[i].r--
						}
						i++
					}
				}
				bucket.items = bucket.items[:lenbucket]
				bucket.Unlock()
				for _, v = range buf {
					t.janitorCh <- v
				}
				buf = buf[:0]
			}
			t.count.Add(int64(-count_del))
		}
	}
}
