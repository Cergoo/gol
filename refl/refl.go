/*
	additional reflection functions pack
	(c) 2013 Cergoo
	under terms of ISC license
*/
package refl

import (
	"github.com/Cergoo/gol/err"
	"reflect"
	"strconv"
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

/*
	If "v" is struct copy fields to "m" map[string]interface{} and return true else return false.
	If "unexported" true copy all fields.
*/
func StructToMap(v interface{}, m map[string]interface{}, unexported bool, prefix string) bool {
	objVal := reflect.Indirect(reflect.ValueOf(v))
	if objVal.Kind() != reflect.Struct {
		return false
	}
	objType := objVal.Type()
	for i := 0; i < objType.NumField(); i++ {
		// access the value of unexported fields
		if !unexported && objType.Field(i).PkgPath != "" {
			continue
		}
		m[prefix+objType.Field(i).Name] = objVal.Field(i).Interface()
	}
	return true
}

const (
	NoInt8 = iota
	NoInt16
	NoInt32
	NoInt64
	NoInt
	NoUint8
	NoUint16
	NoUint32
	NoUint64
	NoUint
	NoFloat32
	NoFloat64
)

func StrTo(from string, to *reflect.Value) error {
	var (
		e error
	)

	switch to.Kind() {

	case reflect.String:
		to.SetString(from)

	case reflect.Int8:
		v, e := strconv.ParseInt(from, 10, 8)
		if e != nil {
			e = err.New(e.Error(), NoInt8)
			return e
		}
		to.SetInt(v)

	case reflect.Int16:
		v, e := strconv.ParseInt(from, 10, 16)
		if e != nil {
			e = err.New(e.Error(), NoInt16)
			return e
		}
		to.SetInt(v)

	case reflect.Int32:
		v, e := strconv.ParseInt(from, 10, 32)
		if e != nil {
			e = err.New(e.Error(), NoInt32)
			return e
		}
		to.SetInt(v)

	case reflect.Int64:
		v, e := strconv.ParseInt(from, 10, 64)
		if e != nil {
			e = err.New(e.Error(), NoInt64)
			return e
		}
		to.SetInt(v)

	case reflect.Int:
		v, e := strconv.ParseInt(from, 10, 0)
		if e != nil {
			e = err.New(e.Error(), NoInt)
			return e
		}
		to.SetInt(v)

	case reflect.Uint8:
		v, e := strconv.ParseUint(from, 10, 8)
		if e != nil {
			e = err.New(e.Error(), NoUint8)
			return e
		}
		to.SetUint(v)

	case reflect.Uint16:
		v, e := strconv.ParseUint(from, 10, 16)
		if e != nil {
			e = err.New(e.Error(), NoUint16)
			return e
		}
		to.SetUint(v)

	case reflect.Uint32:
		v, e := strconv.ParseUint(from, 10, 32)
		if e != nil {
			e = err.New(e.Error(), NoUint32)
			return e
		}
		to.SetUint(v)

	case reflect.Uint64:
		v, e := strconv.ParseUint(from, 10, 64)
		if e != nil {
			e = err.New(e.Error(), NoUint64)
			return e
		}
		to.SetUint(v)

	case reflect.Uint:
		v, e := strconv.ParseUint(from, 10, 0)
		if e != nil {
			e = err.New(e.Error(), NoUint)
			return e
		}
		to.SetUint(v)

	case reflect.Float32:
		v, e := strconv.ParseFloat(from, 32)
		if e != nil {
			e = err.New(e.Error(), NoFloat32)
			return e
		}
		to.SetFloat(v)

	case reflect.Float64:
		v, e := strconv.ParseFloat(from, 64)
		if e != nil {
			e = err.New(e.Error(), NoFloat64)
			return e
		}
		to.SetFloat(v)

	case reflect.Bool:
		if from == "1" || from == "true" {
			to.SetBool(true)
		} else {
			to.SetBool(false)
		}

	}
}

/*
   Convert a map[fieldName]value type map[string]string to structure

func MapStrToStruct(m map[string]string, pointToObj interface{}) {
	objVal := reflect.ValueOf(pointToObj).Elem()
	nfield := objVal.NumField()
	for i:=0; i!=nfield; i++ {
		objVal.Field(i).Set()
	}
}
*/

/*
	IsStruct returns true if the given variable is a struct or a pointer to struct.
*/
func IsStruct(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}

/*
	Return true if v is chan, func, interface, map, pointer, or slice and v is nil
*/
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Interface, reflect.Ptr, reflect.Slice:
		return val.IsNil()
	}

	return false
}

/*
	Return true if v is nil or empty
*/
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Map, reflect.Slice, reflect.Chan:
		return (val.IsNil() || val.Len() == 0)
	case reflect.Func, reflect.Ptr, reflect.Interface:
		return val.IsNil()
	case reflect.Array, reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return val.Bool() == false
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Complex64, reflect.Complex128:
		return val.Complex() == 0
	}

	return false
}
