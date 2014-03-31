/*
	pkg a cookie based session engin implementation
	(c) 2014 Cergoo
	under terms of ISC license
*/
package session

import (
	"gol/cache"
	"gol/cookie"
	"gol/err"
	"gol/genid"
	"math"
	"net/http"
	"time"
)

type (
	TSession struct {
		cache cache.Cache     // local memory cache
		gen   genid.HTTPGenID // id generator
		stor  TStor           // store interface implementation
	}

	TStor interface {
		Get(string) interface{}
		Set(id string, data interface{})
		Del(string)
	}
)

const (
	sid = "sid"
)

// constructor
func NewSessionEngin(timeLiveInCache, lenID uint8, stor TStor) *TSession {
	t := new(TSession)
	t.cache = cache.New(1000, true, time.Minute*time.Duration(timeLiveInCache), nil)
	t.gen = genid.NewHTTPGen(lenID)
	t.stor = stor
	return t
}

// Create new session
func (t *TSession) New(w http.ResponseWriter, data interface{}) (id string) {
	id = t.gen.NewID()
	t.cache.Set(id, data)
	if t.stor != nil {
		t.stor.Set(id, data)
	}
	cookie.SetCookie(w, sid, id, &cookie.Options{Path: "/", MaxAge: math.MaxInt32, HttpOnly: true})
	return
}

// Delete session
func (t *TSession) Del(w http.ResponseWriter, r *http.Request) {
	vcoockie, e := r.Cookie(sid)
	err.Panic(e)
	t.cache.Del(vcoockie.Value)
	if t.stor != nil {
		t.stor.Del(vcoockie.Value)
	}
	cookie.DelCookie(w, sid)
}

// Get session
func (t *TSession) Get(w http.ResponseWriter, r *http.Request) interface{} {
	var val interface{}
	vcoockie, e := r.Cookie(sid)
	if e != nil {
		return nil
	}

	val = t.cache.Get(vcoockie.Value)
	// if cache miss
	if val == nil {
		if t.stor != nil {
			val = t.stor.Get(vcoockie.Value)
		}
		// if session delet then del cookie
		if val == nil {
			cookie.DelCookie(w, sid)
			return nil
		} else {
			t.cache.Set(vcoockie.Value, val)
		}
	}
	return val
}
