/*
	io.Writer implementation
  (c) 2013 Cergoo
	under terms of ISC license
*/

package fastbuf

import (
	"io"
)

type (
	Buf struct {
		buf      []byte
		writeoff int
		readeoff int
	}
)

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

// Resize cap
func (t *Buf) grow(n int) {
	if cap(t.buf) < t.writeoff+n {
		if cap(t.buf) < n {
			t.Grow(n)
		}
		t.Grow(cap(t.buf))
	}
}

// write to buf
func (t *Buf) Write(p []byte) (n int, err error) {
	t.grow(len(p))
	copy(t.buf[t.writeoff:], p)
	t.writeoff += len(p)
	return
}

// get all buf and clear buf
func (t *Buf) FlushP() (r []byte) {
	r = t.buf[:t.writeoff]
	t.writeoff = 0
	return
}

// get all buf and clear buf
func (t *Buf) Flush() (r []byte) {
	r = append(r, t.buf[:t.writeoff]...)
	t.writeoff = 0
	return
}

//
func (t *Buf) Reserve(n int) []byte {
	t.grow(n)
	oldoff := t.writeoff
	t.writeoff += n
	return t.buf[oldoff:t.writeoff]
}

func (t *Buf) Len() int {
	return t.writeoff
}

/* Reader functions */

func (t *Buf) ReadNext(n int) []byte {
	n += t.readeoff
	if n > t.writeoff {
		n = t.writeoff
	}
	data := t.buf[t.readeoff:n]
	t.readeoff = n
	return data
}

func (t *Buf) ReadByte() (b byte, e error) {
	if t.readeoff < t.writeoff {
		b = t.ReadNext(1)[0]
		return
	}
	e = io.EOF
	return
}

func (t *Buf) ReadReset() {
	t.readeoff = 0
}
