/*
	Fast buffer
  (c) 2013 Cergoo
	under terms of ISC license
*/

package fastbuf

type (
	Buf []byte
)

// write to buf
func (t *Buf) Write(p []byte) (n int, err error) {
	*t = append(*t, p...)
	return
}

// get all buf and clear buf
func (t *Buf) Flush() (r []byte) {
	r = *t
	*t = (*t)[:0]
	return
}
