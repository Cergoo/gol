// (c) 2013 Cergoo
// under terms of ISC license

package cacheUint

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

// Returned result code
const (
	ResultExist = iota
	ResultAdd
	ResultNoExistNoAdd
)

// Mode operation set
const (
	OnlyUpdate = iota
	OnlyInsert
	UpdateOrInsert
)

type (
	TItem struct {
		r   uint8
		Key uint64
		Val interface{}
	}

	tBucket struct {
		items []*TItem
		sync.RWMutex
	}

	tCache struct {
		t                     *[]*tBucket      // hash table
		janitorDuration       time.Duration    // janitorDuration janitor
		janitorIfReadThenLive bool             // janitorIfReadThenLive enable
		count                 counter.TCounter // limit items count in cache, 0 - unlimit
		growlock              uint32
		janitorFn             func(*TItem) bool
		stopCh                chan bool
	}

	// for finalizer
	tfinalize struct {
		*tCache
	}

	// Cache interface
	Cache interface {
		GetBucketsStat() (counTItem uint64, countbucket uint32, stat [][2]int)
		Get(uint64) interface{}
		Func(key uint64, f func(*interface{})) interface{}
		Set(key uint64, val interface{}, live, mode uint8) (rval interface{}, actionResult uint8)
		Inc(key uint64, n interface{}, mode uint8) interface{}
		Del(uint64) (val interface{})
		DelAll()
		Range(chan<- *TItem)
		Save(io.Writer) error
		SaveFile(string) error
		Load(io.Reader) error
		LoadFile(string) error
		Len() counter.ICounter
	}
)

/*
New constructor of a new cache:
    janitorIfReadThenLive - if item read then item live;
    janitorDuration       - time to clear items, if 0 then never;
    janitorCh             - chanel to send removed items.
*/
func New(
	janitorIfReadThenLive bool,
	janitorDuration time.Duration,
	janitorFn func(*TItem) bool) Cache {

	ht := make([]*tBucket, 1024)
	for i := range ht {
		ht[i] = &tBucket{
			items: make([]*TItem, 0, initbucketscap),
		}
	}

	t := &tCache{
		janitorDuration:       janitorDuration,
		janitorIfReadThenLive: janitorIfReadThenLive,
		janitorFn:             janitorFn,
		t:                     &ht,
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
func (t *tCache) Len() counter.ICounter {
	return &t.count
}

// GetBucketsStat get statistics of a buckets
func (t *tCache) GetBucketsStat() (counTItem uint64, countbucket uint32, stat [][2]int) {
	var i int
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	tmp1 := make(map[int]int)
	for _, bucket := range ht {
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
	counTItem = t.count.Get()
	return
}

// Get get item value or nil
func (t *tCache) Get(key uint64) (val interface{}) {
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht[key%uint64(len(ht))]
	bucket.RLock()
	for _, v := range bucket.items {
		if v.Key == key {
			if t.janitorIfReadThenLive && v.r < 1 {
				v.r = 1
			}
			val = v.Val
			break
		}
	}
	bucket.RUnlock()
	return
}

// Func accept function to item value
func (t *tCache) Func(key uint64, f func(val *interface{})) (val interface{}) {
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht[key%uint64(len(ht))]
	bucket.RLock()
	for _, v := range bucket.items {
		if v.Key == key {
			if t.janitorIfReadThenLive && v.r < 1 {
				v.r = 1
			}
			if f != nil {
				f(&v.Val)
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
func (t *tCache) Set(key uint64, val interface{}, live uint8, mode uint8) (rval interface{}, actionResult uint8) {
	var (
		v *TItem
	)
	rval = val
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht[key%uint64(len(ht))]
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
		bucket.items = append(bucket.items, &TItem{Key: key, Val: val, r: live})
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
func (t *tCache) Del(key uint64) (val interface{}) {
	var endi int
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht[key%uint64(len(ht))]
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
func (t *tCache) Inc(key uint64, n interface{}, mode uint8) interface{} {
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	bucket := ht[key%uint64(len(ht))]

	bucket.Lock()
	for _, v := range bucket.items {
		if v.Key == key {
			switch value := v.Val.(type) {
			case int:
				v.Val = value + n.(int)
			case int8:
				v.Val = value + n.(int8)
			case int16:
				v.Val = value + n.(int16)
			case int32:
				v.Val = value + n.(int32)
			case int64:
				v.Val = value + n.(int64)
			case uint:
				v.Val = value + n.(uint)
			case uint8:
				v.Val = value + n.(uint8)
			case uint16:
				v.Val = value + n.(uint16)
			case uint32:
				v.Val = value + n.(uint32)
			case uint64:
				v.Val = value + n.(uint64)
			case float32:
				v.Val = value + n.(float32)
			case float64:
				v.Val = value + n.(float64)
			}

			v.r = 1
			bucket.Unlock()
			return v.Val
		}
	}
	if mode == UpdateOrInsert && t.count.Check() {
		lenbucket := len(bucket.items)
		bucket.items = append(bucket.items, &TItem{Key: key, Val: n, r: 1})
		bucket.Unlock()
		if lenbucket > growcountgo && atomic.CompareAndSwapUint32(&t.growlock, 0, 2) {
			go t.grow()
		}
		t.count.Inc()
		return n
	}
	bucket.Unlock()
	return nil
}

// DelAll delete all items
func (t *tCache) DelAll() {
	var i int
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
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
func (t *tCache) Range(ch chan<- *TItem) {
	go t.rangeCache(ch)
}

/*
    non interface range function,
	for breack range close a channel

	Use buffer for unlock bucket
*/
func (t *tCache) rangeCache(ch chan<- *TItem) {
	var (
		buf []*TItem
		v   *TItem
	)
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	for _, bucket := range ht {
		bucket.RLock()
		for _, v = range bucket.items {
			buf = append(buf, v)
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
		item   *TItem
		bucket *tBucket
	)
	enc := gob.NewEncoder(w)
	ht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	defer func() {
		if x := recover(); x != nil {
			for _, bucket = range ht {
				bucket.RUnlock()
			}
			err = fmt.Errorf("error registering item types with Gob library")
		}
	}()
	for _, bucket = range ht {
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
		item TItem
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
		val  *TItem
	)
	oldht := *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
	oldlen := len(oldht)
	newlen := oldlen << 1

	if uint64(newlen) >= math.MaxUint32 {
		// grow only singleton runed
		atomic.StoreUint32(&t.growlock, 0)
		return
	}

	newht := make([]*tBucket, newlen)
	copy(newht, oldht) // possible since it links
	for i = oldlen; i < newlen; i++ {
		newht[i] = &tBucket{}
		newht[i].Lock()
		j++
	}
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.t)), unsafe.Pointer(&newht))
	oldht = nil

	// rehash
	j = oldlen
	for i = 0; i < oldlen; i++ {
		itemsold := make([]*TItem, 0, initbucketscap)
		itemsnew := make([]*TItem, 0, initbucketscap)
		newht[i].Lock()
		for _, val = range newht[i].items {
			if val.Key%uint64(len(newht)) == uint64(i) {
				itemsold = append(itemsold, val)
			} else {
				itemsnew = append(itemsnew, val)
			}
		}

		newht[j].items = itemsnew
		newht[j].Unlock()
		newht[i].items = itemsold
		newht[i].Unlock()
		j++
	}

	// grow only singleton runed
	atomic.StoreUint32(&t.growlock, 0)
}

func (t *tCache) janitor(stop <-chan bool) {
	var (
		i, lenbucket int
		ht           []*tBucket
		bucket       *tBucket
		countDel     uint32
	)
	tick := time.Tick(t.janitorDuration)
	for {
		select {
		case <-stop:
			return
		case <-tick:
			countDel = 0
			ht = *(*[]*tBucket)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.t))))
			for _, bucket = range ht {
				bucket.Lock()
				lenbucket = len(bucket.items)
				for i = 0; i < lenbucket; {
					if bucket.items[i].r == 0 {
						if t.janitorFn == nil || t.janitorFn(bucket.items[i]) {
							lenbucket--
							bucket.items[i], bucket.items[lenbucket] = bucket.items[lenbucket], nil
							countDel++
						}
					} else {
						if bucket.items[i].r < 255 {
							bucket.items[i].r--
						}
						i++
					}
				}
				bucket.items = bucket.items[:lenbucket]
				bucket.Unlock()
			}
			t.count.Add(int64(-countDel))
		}
	}
}
