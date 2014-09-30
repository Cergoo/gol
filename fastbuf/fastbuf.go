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
		readeoff int    // readeoff - index subtracts buffer
		limit    int
		w        io.Writer
	}
)

// New it's constructor buf. If will set linit > 0 then necessarily set w io.Writer
func New(buf []byte, limit int, w io.Writer) (b *Buf) {
	if buf == nil {
		buf = make([]byte, 0, 64)
	}
	return &Buf{limit: limit, w: w, buf: buf}
}

/* Writer functions */

// Grow resize cap
func (t *Buf) Grow(n int) {
	t.grow(cap(t.buf) + n)
}

func (t *Buf) grow(n int) {
	newbuf := make([]byte, n)
	copy(newbuf, t.buf)
	t.buf = newbuf
}

// Write write slice into buf
func (t *Buf) Write(p []byte) (n int, err error) {
	t.buf = append(t.buf, p...)
	if t.limit > 0 && len(t.buf) > t.limit {
		_, err = t.w.Write(t.FlushP())
	}
	return
}

// WriteByte write byte into buf
func (t *Buf) WriteByte(p byte) (err error) {
	t.buf = append(t.buf, p)
	if t.limit > 0 && len(t.buf) > t.limit {
		_, err = t.w.Write(t.FlushP())
	}
	return
}

// FlushP get slice of all buf and clear buf
func (t *Buf) FlushP() (r []byte) {
	r = t.buf
	t.ReadWriteReset()
	return
}

// FlushToWriter get slice of all buf to buf io.Writer and clear buf
func (t *Buf) FlushToWriter() (err error) {
	_, err = t.w.Write(t.FlushP())
	return
}

// GetBuf get current buffer slice
func (t *Buf) GetBuf() []byte {
	return t.buf
}

// SetBuf set a slice into current buffer
func (t *Buf) SetBuf(val []byte) {
	t.buf = val
	if t.readeoff > len(val) {
		t.readeoff = len(val)
	}
}

// Flush get value all buf and clear buf
func (t *Buf) Flush() (r []byte) {
	r = append(r, t.buf...)
	t.ReadWriteReset()
	return
}

// Reserve reserve slice size n
func (t *Buf) Reserve(n int) []byte {
	oldln := len(t.buf)
	n += oldln
	if n > cap(t.buf) {
		t.grow(n + int(n/3))
	}
	t.buf = t.buf[:n]
	return t.buf[oldln:n]
}

// Len get buffer length
func (t *Buf) Len() int {
	return len(t.buf)
}

// Cap get buffer capacity
func (t *Buf) Cap() int {
	return cap(t.buf)
}

/* Reader functions */

// Read it's a io.Reader implementation, read into data
func (t *Buf) Read(data []byte) (n int, e error) {
	var b []byte
	b, e = t.ReadNext(len(data))
	n = copy(data, b)
	return
}

// ReadNext read next n byte
func (t *Buf) ReadNext(n int) (data []byte, e error) {
	n += t.readeoff
	if n > len(t.buf) {
		n = len(t.buf)
		e = io.EOF
	}
	data = t.buf[t.readeoff:n]
	t.readeoff = n
	return
}

// ReadByte read next byte
func (t *Buf) ReadByte() (b byte, e error) {
	if t.readeoff < len(t.buf) {
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
	t.buf = t.buf[:0]
}
