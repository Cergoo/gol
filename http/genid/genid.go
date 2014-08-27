// (c) 2014 Cergoo
// under terms of ISC license

// Package genid it's generate http compatible ID implementation
package genid

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Cergoo/gol/err"
	"math"
	"strings"
)

type (
  // HTTPGenID type 
	HTTPGenID uint8
)

/*
	NewHTTPGen ID creator, resize to base64 encoding, len(id) = 4*length/3
	the actual length can be less per unit
*/
func NewHTTPGen(length uint8) HTTPGenID {
	return HTTPGenID(math.Ceil(float64(length / 4 * 3)))
}

// NewID generate random strind http compatible
func (t HTTPGenID) NewID() string {
	val := make([]byte, t)
	_, e := rand.Read(val)
	err.Panic(e)
	return strings.TrimRight(base64.URLEncoding.EncodeToString(val), "=")
}
