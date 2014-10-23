package bytestack

import (
	"github.com/Cergoo/gol/test"
	"testing"
)

var (
	stack = New(7)
	elem  = []byte{0, 1, 2, 3, 4, 5, 6}
	t1    *test.TT
	a     []byte
)

func Test_1(t *testing.T) {
	t1 = test.New(t)
	stack.Push(elem)
	stack.Push(elem)
	stack.Push(elem)

	a = stack.PopVal()
	t1.Eq(a, elem)

	stack.Push(elem)

	a = stack.PopVal()
	t1.Eq(a, elem)
	a = stack.PopPoint()
	t1.Eq(a, elem)
	a = stack.PopPoint()
	t1.Eq(a, elem)
	a = stack.PopVal()
	t1.Eq(a, []byte(nil))
	a = stack.PopPoint()
	t1.Eq(a, []byte(nil))
}
