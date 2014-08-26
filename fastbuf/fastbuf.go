// (c) 2013 Cergoo
// under terms of ISC license

package fastbuf

import (
	"io"
)

type (
	/*
		writeoff - index buffer is filled
		readeoff - index subtracts buffer
	*/
	Buf struct {
		buf      []byte
		writeoff int
		readeoff int
	}
)

// Constructor
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

// Resize cap
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

// Write slice to buf
func (t *Buf) Write(p []byte) (n int, err error) {
	t.grow(len(p))
	copy(t.buf[t.writeoff:], p)
	t.writeoff += len(p)
	return
}

// Write byte to buf
func (t *Buf) WriteByte(p byte) (err error) {
	t.grow(1)
	t.writeoff++
	t.buf[t.writeoff] = p
	return
}

// Get poiner to all buf and clear buf
func (t *Buf) FlushP() (r []byte) {
	r = t.buf[:t.writeoff]
	t.writeoff = 0
	return
}

// Get all buf and clear buf
func (t *Buf) Flush() (r []byte) {
	r = append(r, t.buf[:t.writeoff]...)
	t.writeoff = 0
	return
}

// Reserve slice size n
func (t *Buf) Reserve(n int) []byte {
	t.grow(n)
	oldoff := t.writeoff
	t.writeoff += n
	return t.buf[oldoff:t.writeoff]
}

// Get buffer length
func (t *Buf) Len() int {
	return t.writeoff
}

// Get buffer capacity
func (t *Buf) Cap() int {
	return cap(t.buf)
}

/* Reader functions */

// Read next slice
func (t *Buf) ReadNext(n int) []byte {
	if t.readeoff >= t.writeoff {
		return nil
	}
	n += t.readeoff
	if n > t.writeoff {
		n = t.writeoff
	}
	data := t.buf[t.readeoff:n]
	t.readeoff = n
	return data
}

// Read next byte
func (t *Buf) ReadByte() (b byte, e error) {
	if t.readeoff < t.writeoff {
		b = t.buf[t.readeoff]
		t.readeoff++
		return
	}
	e = io.EOF
	return
}

// Reset reader index
func (t *Buf) ReadReset() {
	t.readeoff = 0
}
