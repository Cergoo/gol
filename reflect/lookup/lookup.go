// (c) 2013 Cergoo
// under terms of ISC license

// Package lookup it's a lookup reflection functions
package lookup

import (
	"reflect"
	"strconv"
)

// LookupI interface{}
func LookupI(v interface{}, path []string) (vi interface{}, ok bool) {
	var r reflect.Value
	r, ok = Lookup(reflect.ValueOf(v), path)
	if r.IsValid() {
		vi = r.Interface()
	}
	return
}

// Lookup reflect.Value
func Lookup(v reflect.Value, path []string) (r reflect.Value, ok bool) {
	var (
		e error
		i int64
		n string
	)

	for _, p := range path {
		n = string(p)
		v = reflect.Indirect(v)

		switch v.Kind() {
		case reflect.Struct:
			v = v.FieldByName(n)
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(n))
		case reflect.Slice, reflect.Array:
			i, e = strconv.ParseInt(n, 10, 64)
			if e != nil {
				return
			}
			if !(v.Len() > int(i)) {
				return
			}
			v = v.Index(int(i))
		default:
			return
		}

		if !v.IsValid() {
			return
		}
		r = v
	}

	ok = true
	return
}
