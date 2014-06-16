/*
	subscribe channel pack
	(c) 2014 Cergoo
	under terms of ISC license
*/

package chansubscriber

import (
	"runtime"
	"sync"
)

type (
	TChanSubscriber struct {
		closesubscribers bool
		sendStrict       bool
		in               <-chan interface{}
		out              []chan<- interface{}
		sync.Mutex
	}
)

// Constructor
func New(ch <-chan interface{}, closesubscribers bool) *TChanSubscriber {
	t := new(TChanSubscriber)
	t.in = ch
	t.closesubscribers = closesubscribers
	chan_stop := make(chan bool)
	stopRun := func(t *TChanSubscriber) {
		close(chan_stop)
	}
	go t.run(chan_stop)
	runtime.SetFinalizer(t, stopRun)
	return t
}

// Set strict send or not strict
func (t *TChanSubscriber) StrictSet(v bool) {
	t.Lock()
	t.sendStrict = v
	t.Unlock()
}

// Add subscribe
func (t *TChanSubscriber) Subscribe(ch chan<- interface{}) {
	if ch == nil {
		return
	}
	t.Lock()
	for _, v := range t.out {
		if v == ch {
			return
		}
	}
	t.out = append(t.out, ch)
	t.Unlock()
}

// Unsubscribe
func (t *TChanSubscriber) Unsubscribe(ch chan<- interface{}) {
	t.Lock()
	for id, v := range t.out {
		if v == ch {
			vLen := len(t.out) - 1
			t.out[id], t.out[vLen] = t.out[vLen], nil
			t.out = t.out[:vLen-1]
		}
	}
	t.Unlock()
}

// Close all subscribers
func (t *TChanSubscriber) close() {
	if t.closesubscribers {
		t.Unlock()
		t.Lock()
		for i := range t.out {
			close(t.out[i])
		}
		t.Unlock()
	}
}

func (t *TChanSubscriber) run(stop <-chan bool) {
	var (
		outChan chan<- interface{}
		v       interface{}
		ok      bool = true
	)
	for ok {
		select {
		case <-stop:
			return
		case v, ok = <-t.in:
			t.Lock()
			for _, outChan = range t.out {
				if t.sendStrict {
					outChan <- v
				} else {
					select {
					case outChan <- v:
					default:
					}
				}
			}
			t.Unlock()
		}
	}
	defer t.close()
}
