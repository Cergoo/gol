// (c) 2014 Cergoo
// under terms of ISC license

/*
Package mrsw it's a simple multi reades single write controller for resources with ID a string type
*/
package mrsw

import (
	"runtime"
	"sync/atomic"
	"time"
	"unsafe"
)

type (
	TControl struct {
		readers []*string
		writer  *string
		sleep   func()
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
	t := TControl{readers: make([]*string, readersCount)}
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
func (t *TControl) RLock(threadId uint16, resursId string) {
	var wlock *string
	for {
		wlock = (*string)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.writer))))
		if wlock == nil || *wlock != resursId {
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.readers[threadId])), unsafe.Pointer(&resursId))
			wlock = (*string)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.writer))))
			if wlock == nil || *wlock != resursId {
				return
			}
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.readers[threadId])), nil)
		}
		t.sleep()
	}
}

// RUnlock readunlock resurs from thread
func (t *TControl) RUnlock(threadId uint16) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.readers[threadId])), nil)
}

// Lock lock resurs
func (t *TControl) Lock(resursId string) {
	// lock writer
	for !atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&t.writer)), nil, unsafe.Pointer(&resursId)) {
		t.sleep()
	}
	// wait readers
	var rlock *string
	for i := range t.readers {
		rlock = (*string)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.readers[i]))))
		for rlock != nil && *rlock == resursId {
			t.sleep()
			rlock = (*string)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.readers[i]))))
		}
	}
}

// Unlock resurs
func (t *TControl) Unlock() {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.writer)), nil)
}
