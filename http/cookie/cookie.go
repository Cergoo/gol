// (c) 2013 Cergoo
// under terms of ISC license

// Cookie pkg
package cookie

import (
	"net/http"
	"time"
)

type (
	// Cookie options struct
	Options struct {
		Path     string
		Domain   string
		MaxAge   int
		Secure   bool
		HttpOnly bool
	}
)

// Create new *http.Cookie
func NewCookie(name, value string, options *Options) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     options.Path,
		Domain:   options.Domain,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
	if options.MaxAge > 0 {
		cookie.Expires = time.Now().Add(time.Duration(options.MaxAge) * time.Second)
		cookie.MaxAge = options.MaxAge
	} else if options.MaxAge < 0 {
		cookie.Expires = time.Unix(1, 0)
	}
	return cookie
}

// Set cookie
func SetCookie(w http.ResponseWriter, name, value string, options *Options) {
	http.SetCookie(w, NewCookie(name, value, options))
}

// Del cookie
func DelCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, NewCookie(name, "", &Options{Path: "/", MaxAge: -1, HttpOnly: true}))
}
