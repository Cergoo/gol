/*
	additional reflection functions pack
	(c) 2013 Cergoo
	under terms of ISC license
*/
package refl

import (
	"reflect"
)

// A resize to slice all types. It panics if v's Kind is not slice.
func SliceResize(pointToSlice interface{}, newCap int) {
	slice := reflect.ValueOf(pointToSlice).Elem()
	newslice := reflect.MakeSlice(slice.Type(), newCap, newCap)
	reflect.Copy(newslice, slice)
	slice.Set(newslice)
}

// Return true if keys map1 == keys map2. It panics if v's Kind is not map.
func MapKeysEq(map1, map2 interface{}) bool {
	rv1 := reflect.ValueOf(map1)
	rv2 := reflect.ValueOf(map2)
	if rv1.Len() != rv2.Len() {
		return false
	}
	r1keys := rv1.MapKeys()	
	for _, val := range r1keys {
		if !rv2.MapIndex(val).IsValid() {
			return false
		}
	}
	return true
}
