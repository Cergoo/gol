// (c) 2014 Cergoo
// under terms of ISC license

// Package mrsw it's a simple multi reade single write controller
package mrsw

import (
	"math"
	"runtime"
	"sync/atomic"
	"time"
)

type TControl struct {
	readers []uint64
	writer  uint64
	sleep   func()
}

func spinlock() func() {
	return runtime.Gosched
}

func sleep(n time.Duration) func() {
	n *= time.Microsecond
	return func() {
		time.Sleep(n)
	}
}

// New construct new dispatcher
// readersCount - count of a threads reader;
// timetosleep  - time a microsecond on wait of lock, zero - spinlock;
func New(readersCount uint16, timeOnSleep uint16) TControl {
	t := TControl{writer: math.MaxUint64, readers: make([]uint64, readersCount)}
	for i := range t.readers {
		t.readers[i] = math.MaxUint64
	}

	// resolution of collision
	if timeOnSleep == 0 {
		t.sleep = spinlock()
	} else {
		t.sleep = sleep(time.Duration(timeOnSleep))
	}

	return t
}

// RLock readlock resurs from thread
// uses double check
func (t *TControl) RLock(threadId uint16, resursId uint64) {
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
		t.sleep()
	}
}

// RUnlock readunlock resurs from thread
func (t *TControl) RUnlock(threadId uint16) {
	atomic.StoreUint64(&t.readers[threadId], math.MaxUint64)
}

// Lock resurs
func (t *TControl) Lock(resursId uint64) {
	var rlock uint64
	atomic.StoreUint64(&t.writer, resursId)
	for i := range t.readers {
		for {
			rlock = atomic.LoadUint64(&t.readers[i])
			if rlock != resursId {
				break
			}
			t.sleep()
		}
	}
}

// Unlock resurs
func (t *TControl) Unlock() {
	atomic.StoreUint64(&t.writer, math.MaxUint64)
}
