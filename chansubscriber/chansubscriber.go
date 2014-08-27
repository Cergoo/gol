// (c) 2014 Cergoo
// under terms of ISC license

/*
Package chansubscriber it's a implements send messages of writer to a each subscribers.
Features:
  - thread safe;
  - protect of a double subscribe;
  - close a input channel for ending send messages.
*/
package chansubscriber

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type (
	// TChanSubscriber main struct
	TChanSubscriber struct {
		closeSubscribers bool
		sendStrict       bool
		in               <-chan interface{}
		out              *[]chan<- interface{}
		mu               sync.Mutex
	}
)

// New constructor of a new chansubscriber:
//  ch               - channel writer;
//  sendStrict       - if true not drop packets;
//  closeSubscribers - close all reader channel after close writer.
func New(ch <-chan interface{}, sendStrict, closeSubscribers bool) *TChanSubscriber {
	out := make([]chan<- interface{}, 0)
	t := &TChanSubscriber{
		in:               ch,
		sendStrict:       sendStrict,
		closeSubscribers: closeSubscribers,
		out:              &out,
	}
	go t.send()
	return t
}

// Subscribe subscribe channel
func (t *TChanSubscriber) Subscribe(ch chan<- interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if ch == nil {
		return
	}
	outslice := *(*[]chan<- interface{})(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.out))))
	for _, v := range outslice {
		if v == ch {
			return
		}
	}
	newoutslice := make([]chan<- interface{}, len(outslice), len(outslice)+1)
	copy(newoutslice, outslice)
	newoutslice = append(newoutslice, ch)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.out)), unsafe.Pointer(&newoutslice))
}

// Unsubscribe unsubscribe channel
func (t *TChanSubscriber) Unsubscribe(ch chan<- interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	outslice := *(*[]chan<- interface{})(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.out))))
	for id, v := range outslice {
		if v == ch {
			newoutslice := make([]chan<- interface{}, len(outslice)-1)
			copy(newoutslice, outslice[:id])
			copy(newoutslice[id:], outslice[id+1:])
			atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&t.out)), unsafe.Pointer(&newoutslice))
			break
		}
	}
}

// Len get count subscribers
func (t *TChanSubscriber) Len() int {
	return len(*(*[]chan<- interface{})(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.out)))))
}

// helper function, Close all subscribers
func (t *TChanSubscriber) close() {
	if t.closeSubscribers {
		t.mu.Lock()
		defer t.mu.Unlock()
		outslice := *t.out
		for i := range outslice {
			close(outslice[i])
		}
	}
}

func (t *TChanSubscriber) send() {
	var (
		outChan  chan<- interface{}
		v        interface{}
		f        func()
		outslice []chan<- interface{}
	)

	if t.sendStrict {
		f = func() {
			for _, outChan = range outslice {
				outChan <- v
			}
		}
	} else {
		f = func() {
			for _, outChan = range outslice {
				select {
				case outChan <- v:
				default:
				}
			}
		}
	}

	defer t.close()

	for v = range t.in {
		outslice = *(*[]chan<- interface{})(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&t.out))))
		f()
	}
}
