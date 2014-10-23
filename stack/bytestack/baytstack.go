// (c) 2014 Cergoo
// under terms of ISC license

// Package bytestack it's a simple fixed length slice stack implementation. No thread safe.
package bytestack

import (
	"fmt"
)

type (
	// TStack it's a main struct
	TStack struct {
		Stack      []byte
		LenElement int
	}
)

// New create new bytes stack
func New(lenElement int) *TStack {
	return &TStack{LenElement: lenElement}
}

// Push push item into stack
func (t *TStack) Push(val []byte) error {
	if len(val) != t.LenElement {
		return fmt.Errorf("Mismatch length stack element %d, and length value %d", t.LenElement, len(val))
	}
	t.Stack = append(t.Stack, val...)
	return nil
}

// PopPoint pop slice as pointer to value item from stack, nil if stack empty
func (t *TStack) PopPoint() (val []byte) {
	n := len(t.Stack) - t.LenElement
	if n >= 0 {
		val = t.Stack[n:]
		t.Stack = t.Stack[:n]
	}
	return
}

// PopVal pop slice as copy value item from stack, nil if stack empty
func (t *TStack) PopVal() (val []byte) {
	n := len(t.Stack) - t.LenElement
	if n >= 0 {
		val = make([]byte, t.LenElement)
		copy(val, t.Stack[n:])
		t.Stack = t.Stack[:n]
	}
	return
}

// Clear clear stack
func (t *TStack) Clear() {
	t.Stack = t.Stack[:0]
}
