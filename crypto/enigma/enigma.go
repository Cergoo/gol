// (c) 2016 Cergoo
// under terms of ISC license

// Package enigma its a crypto encripter/decripter, base64.URL safety from web, with periodically changes key
package enigma

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	cip "golang.org/x/crypto/twofish"
	"io"
	"runtime"
	"sync"
	"time"
)

const cipherlen = 16

type (
	tendecr struct {
		id        byte
		encrypter cipher.BlockMode
		decrypter cipher.BlockMode
	}

	TBox struct {
		rwmu   sync.RWMutex
		label  []byte
		stopCh chan bool
		b      cipher.Block
		endecr [2]*tendecr
	}
)

// New create new enigma box
func New(periodDuration time.Duration) (*TBox, error) {
	var e error
	t := &TBox{
		stopCh: make(chan bool),
		label:  make([]byte, 4),
	}
	t.endecr[0] = &tendecr{}
	t.endecr[1] = &tendecr{}

	io.ReadFull(rand.Reader, t.label)

	key := make([]byte, cipherlen)
	io.ReadFull(rand.Reader, key)
	t.b, e = cip.NewCipher(key)

	t.create()

	go t.changState(periodDuration)
	runtime.SetFinalizer(t, stop)

	return t, e
}

func stop(t *TBox) {
	close(t.stopCh)
}

func (t *TBox) create() {
	iv := make([]byte, cipherlen)
	io.ReadFull(rand.Reader, iv)
	t.endecr[1] = t.endecr[0]
	t.endecr[0] = &tendecr{
		id:        t.endecr[1].id + 1,
		encrypter: cipher.NewCBCEncrypter(t.b, iv),
		decrypter: cipher.NewCBCDecrypter(t.b, iv),
	}

}

/*
save old encripter&decripter
create new encripter&decripter
create init vector
*/
func (t *TBox) changState(periodDuration time.Duration) {
	for {
		select {
		case <-t.stopCh:
			return
		case <-time.Tick(periodDuration):
			t.rwmu.Lock()
			t.create()
			t.rwmu.Unlock()
		}
	}
}

// Encryp []byte to string
func (t *TBox) Encrypt(b []byte) string {

	l := len(t.label) + int(1) + len(b)

	remainder := uint8(t.b.BlockSize() - l%t.b.BlockSize())
	l += int(remainder)

	crypt := make([]byte, 0, l+1)
	crypt = append(crypt, t.label...)
	crypt = append(crypt, remainder)
	crypt = append(crypt, b...)
	crypt = crypt[:cap(crypt)-1]

	t.rwmu.RLock()
	t.endecr[0].encrypter.CryptBlocks(crypt, crypt)
	crypt = append(crypt, t.endecr[0].id)
	t.rwmu.RUnlock()

	return base64.URLEncoding.EncodeToString(crypt)
}

// Decrypt string to []byte
func (t *TBox) Decrypt(b string) ([]byte, error) {
	dec, e := base64.URLEncoding.DecodeString(b)
	if e != nil {
		return nil, e
	}
	var id byte

	t.rwmu.RLock()
	switch dec[len(dec)-1] {
	case t.endecr[0].id:
		id = 0
	case t.endecr[1].id:
		id = 1
	default:
		t.rwmu.RUnlock()
		return nil, errors.New("not found decrypter")
	}

	dec = dec[:len(dec)-1]
	t.endecr[id].decrypter.CryptBlocks(dec, dec)
	t.rwmu.RUnlock()

	if bytes.Compare(t.label, dec[:len(t.label)]) != 0 {
		return nil, errors.New("decrypte not correct")
	}

	return dec[len(t.label)+1 : len(dec)-int(dec[len(t.label)])], e
}
