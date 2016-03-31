// (c) 2016 Cergoo
// under terms of ISC license

// Package enigma its a crypto encripter/decripter with periodically changes key
package enigma

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
	"time"
)

var (
	l      = 141
	tested = make([][]byte, 20, 20)
	box    *TBox
)

func init() {
	for i := range tested {
		tested[i] = make([]byte, l)
		io.ReadFull(rand.Reader, tested[i])
	}
	box, _ = New(time.Duration(1) * time.Minute)
}

func Test1(t *testing.T) {
	var (
		src []byte
		e   error
	)

	for i := range tested {
		src, e = box.Decrypt(box.Encrypt(tested[i]))
		if e != nil {
			t.Errorf("Error:", i, e)
		}
		if bytes.Compare(src, tested[i]) != 0 {
			t.Errorf("Error1:", i)
		}

	}
}

func Test2(t *testing.T) {
	var (
		src []byte
		enc []string
		e   error
	)

	for i := range tested {
		enc = append(enc, box.Encrypt(tested[i]))
	}

	<-time.Tick(time.Duration(1) * time.Minute)

	for i := range enc {
		src, e = box.Decrypt(enc[i])
		if e != nil {
			t.Errorf("Error:", i, e)
		}
		if bytes.Compare(src, tested[i]) != 0 {
			t.Errorf("Error1:", i)
		}

	}

}

func Test3(t *testing.T) {
	var (
		src []byte
		enc []string
		e   error
	)

	for i := range tested {
		enc = append(enc, box.Encrypt(tested[i]))
	}

	<-time.Tick(time.Duration(2) * time.Minute)

	for i := range enc {
		src, e = box.Decrypt(enc[i])
		if e == nil {
			t.Errorf("Error:", i, e)
		}
		if bytes.Compare(src, tested[i]) == 0 {
			t.Errorf("Error1:", i)
		}

	}

}
