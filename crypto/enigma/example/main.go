// Example use pkg.
package main

import (
	"fmt"
	"github.com/Cergoo/gol/crypto/enigma"
	"time"
)

func main() {
	b := []byte("Test str")
	box, _ := enigma.New(time.Duration(1) * time.Minute)
	enc := box.Encrypt(b)
	fmt.Println(enc, len(enc))
	dec, e := box.Decrypt(enc)
	fmt.Println(string(dec), len(dec), e)

	b = []byte("Test str new Test str new")
	box, _ = enigma.New(time.Duration(1) * time.Minute)
	enc = box.Encrypt(b)
	fmt.Println(enc, len(enc))
	dec, e = box.Decrypt(enc)
	fmt.Println(string(dec), len(dec), e)

}
