package mrswd

import (
	"testing"
)

func Test1(tt *testing.T) {
	d := New(10, 1000)
	d.Lock("n1")
	d.Unlock()
	i := d.RLock("n1")
	d.RUnlock(i)
}
