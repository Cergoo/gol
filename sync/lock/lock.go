// (c) 2014 Cergoo
// under terms of ISC license

// Package lock it's a simple spin lock structure implementation
package lock

import (
	"runtime"
	"sync/atomic"
)

const (
	lockUNLOCKED = 0
	lockLOCKED   = 1
)

type SpinLock int32

func (t *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(t), lockUNLOCKED, lockLOCKED)
}

func (t *SpinLock) Lock() {
	for {
		if t.TryLock() {
			return
		}
		runtime.Gosched()
	}
}

func (t *SpinLock) Unlock() {
	*t = lockUNLOCKED
}
