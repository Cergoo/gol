package stack

import (
	"github.com/Cergoo/gol/test"
	"testing"
)

var (
	v  = new(Tstack)
	t1 *test.TT
)

func Test_MapKeysEq_ok1(t *testing.T) {
	t1 = test.New(t)
	obj := "nnn"

	v.Push(obj)
	a, b := v.Pop()
	t1.Eq(a, obj)
	t1.Eq(b, true)

	v.Push(nil)
	a, b = v.Pop()
	t1.Eq(a, nil)
	t1.Eq(b, true)

	a, b = v.Pop()
	t1.Eq(a, nil)
	t1.Eq(b, false)
}
