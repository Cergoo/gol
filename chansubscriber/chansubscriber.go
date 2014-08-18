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
		closeSubscribers bool
		sendStrict       bool
		in               <-chan interface{}
		out              []chan<- interface{}
		sync.Mutex
	}
)

// Constructor
func New(ch <-chan interface{}, sendStrict, closeSubscribers bool) *TChanSubscriber {
	t := &TChanSubscriber{
		in:               ch,
		sendStrict:       sendStrict,
		closeSubscribers: closeSubscribers,
	}
	chan_stop := make(chan bool)
	stopRun := func(t *TChanSubscriber) {
		close(chan_stop)
	}
	go t.send(chan_stop)
	runtime.SetFinalizer(t, stopRun)
	return t
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
	if t.closeSubscribers {
		t.Unlock()
		t.Lock()
		for i := range t.out {
			close(t.out[i])
		}
		t.Unlock()
	}
}

func (t *TChanSubscriber) send(stop <-chan bool) {
	var (
		outChan chan<- interface{}
		v       interface{}
		ok      bool
		f       func()
	)

	if t.sendStrict {
		f = func() {
			for _, outChan = range t.out {
				outChan <- v
			}
		}
	} else {
		f = func() {
			for _, outChan = range t.out {
				select {
				case outChan <- v:
				default:
				}
			}
		}
	}

	v, ok = <-t.in
	for ok {

		t.Lock()
		f()
		t.Unlock()

		select {
		case <-stop:
			return
		case v, ok = <-t.in:
		}
	}
	defer t.close()
}
