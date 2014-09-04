// (c) 2013 Cergoo
// under terms of ISC license

// Package lookup it's a lookup reflection functions
package lookup

import (
	"bytes"
	"reflect"
	"strconv"
)

// LookupI interface{}
func LookupI(v interface{}, path []byte) (vi interface{}, fullpath bool) {
	var r reflect.Value
	r, fullpath = Lookup(reflect.ValueOf(v), path)
	if r.IsValid() {
		vi = r.Interface()
	}
	return
}

// Lookup reflect.Value
func Lookup(v reflect.Value, path []byte) (r reflect.Value, fullpath bool) {
	var (
		e error
		i int64
		n string
	)
	part := bytes.Split(path, []byte("/"))

	for _, p := range part {
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
			v = v.Index(int(i))
		default:
			return
		}

		if !v.IsValid() {
			return
		}
		r = v
	}

	fullpath = true
	return
}
