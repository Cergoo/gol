package fastbuf

import (
	"testing"
)

var fb Buf

func Test1(t *testing.T) {
	v := []byte("12")
	fb.Write(v)
	if string(fb.Flush()) != string(v) {
		t.Errorf("Error:")
	}
	if len(fb) > 0 {
		t.Errorf("Error:")
	}

}
