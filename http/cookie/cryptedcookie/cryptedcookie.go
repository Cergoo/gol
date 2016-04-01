// (c) 2016 Cergoo
// under terms of MIT license

// Package cryptedcookie
package cryptedcookie

import (
	"errors"
	"github.com/Cergoo/gol/crypto/enigma"
	"github.com/Cergoo/gol/http/cookie/cookie"
	"net/http"
	"time"
)

type (
	TBoxCookie struct {
		box    *enigma.TBox
		maxAge int
	}
)

// New create cryptedcookie box
func New(maxAge int) (*TBoxCookie, error) {
	if maxAge < 10 {
		return nil, errors.New("maxAge must be > 10 seconds")
	}
	box, e := enigma.New(time.Duration(maxAge/2) * time.Second)
	if e != nil {
		return nil, e
	}
	return &TBoxCookie{
		box:    box,
		maxAge: maxAge,
	}, e
}

// NewCookie create new *http.Cookie
func (t *TBoxCookie) NewCookie(name, value string, options *cookie.Options) *http.Cookie {
	options.MaxAge = t.maxAge
	return cookie.NewCookie(name, t.box.Encrypt([]byte(value)), options)
}

// SetCookie set cookie
func (t *TBoxCookie) SetCookie(w http.ResponseWriter, name, value string, options *cookie.Options) {
	http.SetCookie(w, t.NewCookie(name, value, options))
}

// DelCookie del cookie
func (t *TBoxCookie) DelCookie(w http.ResponseWriter, name string) {
	cookie.DelCookie(w, name)
}

// GetCookie get cookie
func (t *TBoxCookie) GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	var dec []byte

	co, e := r.Cookie(name)
	if e != nil {
		return nil, e
	}

	dec, e = t.box.Decrypt(co.Value)
	if e != nil {
		return nil, e
	}

	co.Value = string(dec)

	return co, e
}
