// (c) 2014 Cergoo
// under terms of ISC license

// Package mrsw it's a simple multi reade single write dispatcher
package mrsw

import (
	"math"
	"runtime"
	"sync/atomic"
)

type TDispatcher struct {
	readers []uint64
	writer  uint64
}

// New construct new dispatcher
func New(readersCount int) *TDispatcher {
	t := &TDispatcher{writer: math.MaxUint64, readers: make([]uint64, readersCount)}
	for i := range t.readers {
		t.readers[i] = math.MaxUint64
	}
	return t
}

// RLock readlock resurs from thread
// uses double check
func (t *TDispatcher) RLock(threadId int, resursId uint64) {
	var wlock uint64
	for {
		wlock = atomic.LoadUint64(&t.writer)
		if wlock != resursId {
			atomic.StoreUint64(&t.readers[threadId], resursId)
			wlock = atomic.LoadUint64(&t.writer)
			if wlock != resursId {
				return
			}
			atomic.StoreUint64(&t.readers[threadId], math.MaxUint64)
		}
		runtime.Gosched()
	}
}

// RUnlock readunlock resurs from thread
func (t *TDispatcher) RUnlock(threadId int) {
	atomic.StoreUint64(&t.readers[threadId], math.MaxUint64)
}

// Lock resurs
func (t *TDispatcher) Lock(resursId uint64) {
	var rlock uint64
	atomic.StoreUint64(&t.writer, resursId)
	for i := range t.readers {
		for {
			rlock = atomic.LoadUint64(&t.readers[i])
			if rlock != resursId {
				break
			}
			runtime.Gosched()
		}
	}
}

// Unlock resurs
func (t *TDispatcher) Unlock() {
	atomic.StoreUint64(&t.writer, math.MaxUint64)
}
