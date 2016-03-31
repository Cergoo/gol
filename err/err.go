// (c) 2013-2014 Cergoo
// under terms of ISC license

// Package err it's a editable error implementation.
package err

import (
	"errors"
	"fmt"
	"log"
)

// OpenErr editable error struct.
type OpenErr struct {
	Text string
	Code int
}

// New create new error.
func New(e string, code int) *OpenErr {
	return &OpenErr{e, code}
}

// Error it's interface error metod.
func (t *OpenErr) Error() string {
	return t.Text
}

// Panic gen.
func Panic(e error) {
	if e != nil {
		panic(e)
	}
}

// Panic gen.
func LogPanic(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// Panic gen.
func PanicBool(ok bool, e string, code int) {
	if !ok {
		panic(New(e, code))
	}
}

// Nopanic
func Nopanic(e *error) {
	if v := recover(); v != nil && e != nil {
		*e = errors.New(fmt.Sprintln(v))
	}
}
