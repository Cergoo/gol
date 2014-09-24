// (c) 2014 Cergoo
// under terms of ISC license

// Package stack it's a easy stack implementation. No thread safe.
package stack

type (
	// TStack it's a main struct
	TStack []interface{}
)

// Push push item into stack
func (t *TStack) Push(val interface{}) {
	*t = append(*t, val)
}

// Pop pop item from stack, ok = false if stack empty
func (t *TStack) Pop() (val interface{}, ok bool) {
	if len(*t) < 1 {
		return
	}
	val, ok = (*t)[len(*t)-1], true
	*t = (*t)[:len(*t)-1]
	return
}

// Peek peek item from stack, ok = false if stack empty
func (t TStack) Peek() (val interface{}, ok bool) {
	if len(t) < 1 {
		return
	}
	val, ok = t[len(t)-1], true
	return
}
