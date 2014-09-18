package mrsw

import (
	"testing"
)

func Test1(tt *testing.T) {
	d := New(10, 1000)
	d.Lock(1)
	d.Unlock()
	d.RLock(1, 12)
	d.RUnlock(1)
}
