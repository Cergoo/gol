// (c) 2013-2014 Cergoo
// under terms of ISC license

package err

// Editable error struct.
type OpenErr struct {
	Text string
	Code int
}

// Create new error.
func New(e string, code int) *OpenErr {
	return &OpenErr{e, code}
}

// Interface error metod.
func (t *OpenErr) Error() string {
	return t.Text
}

// Panic gen.
func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
