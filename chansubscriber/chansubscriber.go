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
		keyid            map[string]uint
		sync.Mutex
	}
)

// Constructor
func New(ch <-chan interface{}, closesubscribers bool) *TChanSubscriber {
	t := new(TChanSubscriber)
	t.in = ch
	t.closesubscribers = closesubscribers
	t.keyid = make(map[string]uint)
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
func (t *TChanSubscriber) Subscribe(name string, ch chan<- interface{}) {
	t.Lock()
	defer t.Unlock()

	id, ok := t.keyid[name]
	if ok {
		t.out[id] = ch
		return
	}

	for i := range t.out {
		if t.out[i] == nil {
			t.out[i] = ch
			t.keyid[name] = uint(i)
			return
		}
	}

	t.keyid[name] = uint(len(t.out))
	t.out = append(t.out, ch)
}

// Unsubscribe
func (t *TChanSubscriber) Unsubscribe(name string) {
	t.Lock()
	delete(t.keyid, name)
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
		case v, ok = <-t.in:
			t.Lock()
			for _, outChan = range t.out {
				if outChan != nil {
					if t.sendStrict {
						outChan <- v
					} else {
						select {
						case outChan <- v:
						default:
						}
					}
				}
			}
			t.Unlock()
		}
	}
	defer t.close()
}
