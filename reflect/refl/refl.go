// (c) 2013 Cergoo
// under terms of ISC license

// Package refl it's additional reflection functions
package refl

import (
	"reflect"
)

// SliceResize a resize to slice all types. It panics if v's Kind is not slice.
func SliceResize(pointToSlice interface{}, newCap int) {
	slice := reflect.ValueOf(pointToSlice).Elem()
	newslice := reflect.MakeSlice(slice.Type(), newCap, newCap)
	reflect.Copy(newslice, slice)
	slice.Set(newslice)
}

// MapKeysEq return true if keys map1 == keys map2. It panics if v's Kind is not map.
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

//StructToMap if "v" is struct copy fields to "m" map[string]interface{} and return true else return false.
func StructToMap(v interface{}, m map[string]interface{}, prefix string) bool {
	objVal := reflect.Indirect(reflect.ValueOf(v))
	if objVal.Kind() != reflect.Struct {
		return false
	}
	objType := objVal.Type()
	for i := 0; i < objType.NumField(); i++ {
		// access the value of unexported fields objVal.Field(i).Interface() panic generate
		if objType.Field(i).PkgPath != "" {
			continue
		}
		m[prefix+objType.Field(i).Name] = objVal.Field(i).Interface()
	}
	return true
}

// IsStruct returns true if the given variable is a struct or a pointer to struct.
func IsStruct(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}

// IsNil return true if v is chan, func, interface, map, pointer, or slice and v is nil
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

// IsEmpty return true if v is nil or empty
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

type (
	// TField struct about field sruct
	TField struct {
		Fname     string
		Fexported bool
		Ftype     string
		Fkind     reflect.Kind
		Ftags     reflect.StructTag
	}
)

// StructFields return v type and v fields description
func StructFields(v interface{}) (vType reflect.Type, vFields []*TField) {
	var (
		field reflect.StructField
	)
	vType = reflect.Indirect(reflect.ValueOf(v)).Type()
	numFields := vType.NumField()
	vFields = make([]*TField, 0, numFields)
	for i := 0; i < numFields; i++ {
		field = vType.Field(i)
		fieldInfo := &TField{
			Fname: field.Name,
			Ftype: field.Type.String(),
			Fkind: field.Type.Kind(),
			Ftags: field.Tag,
		}
		if field.PkgPath == "" {
			fieldInfo.Fexported = true
		}
		vFields = append(vFields, fieldInfo)
	}
	return
}
