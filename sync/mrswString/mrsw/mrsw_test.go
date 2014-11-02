package mrsw

import (
	"testing"
)

func Test1(tt *testing.T) {
	d := New(10, 1000)
	d.Lock("n1")
	d.Unlock()
	i := uint16(1)
	d.RLock(i, "n1")
	d.RUnlock(i)
}
