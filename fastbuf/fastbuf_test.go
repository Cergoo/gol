package fastbuf

import (
	"testing"
)

var fb = New(nil)

func Test1(t *testing.T) {
	v := []byte("12")
	fb.Write(v)
	if string(fb.Flush()) != string(v) {
		t.Errorf("Error:")
	}
	if fb.Len() > 0 {
		t.Errorf("Error:")
	}

}
