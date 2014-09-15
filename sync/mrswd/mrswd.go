// (c) 2014 Cergoo
// under terms of ISC license

// Package mrswd it's a multi reade single write control dispatcher
package mrswd

import (
	"github.com/Cergoo/gol/lockfree/stack"
	"github.com/Cergoo/gol/sync/mrsw"
)

type (
	TDispatcher struct {
		threadstack stack.Tstack
		control     mrsw.TControl
	}
)

// New construct new dispatcher
func New(threadcount uint16) (t TDispatcher) {
	t = TDispatcher{threadstack: stack.Tstack{}, control: mrsw.New(threadcount)}
	for i := uint16(0); i < threadcount; i++ {
		t.threadstack.Push(i)
	}
	return
}

// RLock readlock resurs from thread
func (t *TDispatcher) RLock(resursId uint64) (threadid uint16) {
	threadid = t.threadstack.PopWait().(uint16)
	t.control.RLock(threadid, resursId)
	return
}

// RUnlock readunlock resurs from thread
func (t *TDispatcher) RUnlock(threadid uint16) {
	t.control.RUnlock(threadid)
	t.threadstack.Push(threadid)
}

// Lock writer
func (t *TDispatcher) Lock(resursId uint64) {
	t.control.Lock(resursId)
}

// Unlock writer
func (t *TDispatcher) Unlock() {
	t.control.Unlock()
}
