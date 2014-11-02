// (c) 2014 Cergoo
// under terms of ISC license

/*
Package mrswd it's a dispatcher of a multi reade single write controller for resources with ID a type string
*/
package mrswd

import (
	"github.com/Cergoo/gol/sync/mrswString/mrsw"
	"time"
)

type (
	TDispatcher struct {
		chThread chan uint16
		mrsw.TControl
	}
)

// New construct new dispatcher
func New(threadcount uint16, timeOnSleep time.Duration) (t *TDispatcher) {
	t = &TDispatcher{chThread: make(chan uint16, threadcount), TControl: mrsw.New(threadcount, timeOnSleep)}
	for i := uint16(0); i < threadcount; i++ {
		t.chThread <- i
	}
	return
}

// RLock readlock resurs from thread
func (t *TDispatcher) RLock(resursId string) (threadid uint16) {
	threadid = <-t.chThread
	t.TControl.RLock(threadid, resursId)
	return
}

// RUnlock readunlock resurs from thread
func (t *TDispatcher) RUnlock(threadid uint16) {
	t.TControl.RUnlock(threadid)
	t.chThread <- threadid
}
