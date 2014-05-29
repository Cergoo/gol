/*
  extension []bytes pkg
  (c)	2014 Cergoo
  under terms of ISC license
*/

package golbytes

import (
	"bytes"
)

//  equal head
func HeadEqual(source []byte, head []byte) bool {
	return len(source) >= len(head) && bytes.Equal(source[:len(head)], head)
}

//  equal tail
func TailEqual(source []byte, tail []byte) bool {
	i := len(source) - len(tail)
	return i >= 0 && bytes.Equal(source[i:], tail)
}
