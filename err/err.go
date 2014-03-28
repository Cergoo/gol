/*
	error pkg
	(c) 2013 Cergoo
	under terms of ISC license
*/
package err

// editable error struct
type OpenErr struct {
	Text string
	Code int
}

// create new error
func New(e string, code int) *OpenErr {
	return &OpenErr{e, code}
}

// interface error metod
func (t *OpenErr) Error() string {
	return t.Text
}

// panic gen
func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
