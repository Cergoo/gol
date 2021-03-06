// (c) 2014 Cergoo
// under terms of ISC license

/*
Package mrsw it's a simple multi reade single write controller for resources with
ID in the range of 0, maxUint64-1
*/
package mrsw

import (
	"math"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	unlocked = math.MaxUint64
)

type (
	TControl struct {
		readers []uint64
		sleep   func()
		writer  uint64
	}
)

func spinlock() func() {
	return runtime.Gosched
}

func sleep(n time.Duration) func() {
	return func() {
		time.Sleep(n)
	}
}

// New construct new dispatcher
// readersCount - count of a threads reader;
// timetosleep  - time a microsecond on wait of lock, zero - spinlock;
func New(readersCount uint16, timeOnSleep time.Duration) TControl {
	t := TControl{writer: unlocked, readers: make([]uint64, readersCount)}
	for i := range t.readers {
		t.readers[i] = unlocked
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
	for {
		if atomic.LoadUint64(&t.writer) != resursId {
			atomic.StoreUint64(&t.readers[threadId], resursId)
			if atomic.LoadUint64(&t.writer) != resursId {
				return
			}
			atomic.StoreUint64(&t.readers[threadId], unlocked)
		}
		t.sleep()
	}
}

// RUnlock readunlock resurs from thread
func (t *TControl) RUnlock(threadId uint16) {
	atomic.StoreUint64(&t.readers[threadId], unlocked)
}

// Lock resurs
func (t *TControl) Lock(resursId uint64) {
	for !atomic.CompareAndSwapUint64(&t.writer, unlocked, resursId) {
		t.sleep()
	}
	for i := range t.readers {
		for atomic.LoadUint64(&t.readers[i]) == resursId {
			t.sleep()
		}
	}
}

// Unlock resurs
func (t *TControl) Unlock() {
	atomic.StoreUint64(&t.writer, unlocked)
}
