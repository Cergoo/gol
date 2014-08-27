// (c) 2013-2014 Cergoo
// under terms of ISC license

// Package err it's a editable error implementation.
package err

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
