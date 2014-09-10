// (c) 2013 Cergoo
// under terms of ISC license

// Package fastbuf it's io.Writer implementation
package fastbuf

import (
	"io"
)

type (
	// Buf struct of a buf
	Buf struct {
		buf      []byte //
		writeoff int    // writeoff - index buffer is filled
		readeoff int    // readeoff - index subtracts buffer
	}
)

// New it's constructor buf
func New(buf []byte) (b *Buf) {
	b = new(Buf)
	if buf == nil {
		buf = make([]byte, 64)
	} else {
		b.writeoff = len(buf)
		buf = buf[:cap(buf)]
	}
	b.buf = buf
	return
}

/* Writer functions */

// Grow resize cap
func (t *Buf) Grow(n int) {
	newbuf := make([]byte, cap(t.buf)+n)
	copy(newbuf, t.buf)
	t.buf = newbuf
}

func (t *Buf) grow(n int) {
	if cap(t.buf) < t.writeoff+n {
		if cap(t.buf) < n {
			t.Grow(n)
		}
		t.Grow(cap(t.buf))
	}
}

// Write write slice into buf
func (t *Buf) Write(p []byte) (n int, err error) {
	t.grow(len(p))
	copy(t.buf[t.writeoff:], p)
	t.writeoff += len(p)
	return
}

// WriteByte write byte into buf
func (t *Buf) WriteByte(p byte) (err error) {
	t.grow(1)
	t.buf[t.writeoff] = p
	t.writeoff++
	return
}

// FlushP get slice of all buf and clear buf
func (t *Buf) FlushP() (r []byte) {
	r = t.buf[:t.writeoff]
	t.writeoff = 0
	return
}

// Show get show current buffer
func (t *Buf) Show() []byte {
	return t.buf[:t.writeoff]
}

// Flush get value all buf and clear buf
func (t *Buf) Flush() (r []byte) {
	r = append(r, t.buf[:t.writeoff]...)
	t.writeoff = 0
	return
}

// Reserve reserve slice size n
func (t *Buf) Reserve(n int) []byte {
	t.grow(n)
	oldoff := t.writeoff
	t.writeoff += n
	return t.buf[oldoff:t.writeoff]
}

// Len get buffer length
func (t *Buf) Len() int {
	return t.writeoff
}

// Cap get buffer capacity
func (t *Buf) Cap() int {
	return cap(t.buf)
}

/* Reader functions */

// ReadNext read next n byte
func (t *Buf) ReadNext(n int) (data []byte, e error) {
	if t.readeoff >= t.writeoff {
		e = io.EOF
		return
	}
	n += t.readeoff
	if n > t.writeoff {
		n = t.writeoff
		e = io.EOF
	}
	data = t.buf[t.readeoff:n]
	t.readeoff = n
	return
}

// ReadByte read next byte
func (t *Buf) ReadByte() (b byte, e error) {
	if t.readeoff < t.writeoff {
		b = t.buf[t.readeoff]
		t.readeoff++
		return
	}
	e = io.EOF
	return
}

// ReadReset reset reader index
func (t *Buf) ReadReset() {
	t.readeoff = 0
}

// ReadWriteReset reset reader and writer indexs
func (t *Buf) ReadWriteReset() {
	t.readeoff = 0
	t.writeoff = 0
}
