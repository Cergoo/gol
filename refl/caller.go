/*
	additional reflection functions pack
	(c) 2013 Cergoo
	under terms of ISC license
*/
package refl

import (
	"gol/err"
	"reflect"
)

// Function caller

type (
	FuncMap   map[string]reflect.Value
	FuncSlice []reflect.Value
)

// Add to function map
func (t FuncMap) Add(name string, f interface{}) {
	t[name] = reflect.ValueOf(f)
}

// Add to function slice
func (t *FuncSlice) Add(id int, f interface{}) {
	*t = append(*t, reflect.ValueOf(f))
}

// Call and return interface{}
func (t FuncMap) Calli(name string, params ...interface{}) []interface{} {
	var result []interface{}
	r := t.Call(name, params...)
	if len(r) > 0 {
		result = make([]interface{}, 0, len(r))
		for i := range r {
			result = append(result, r[i].Interface())
		}
	}
	return result
}

// Call and return interface{}
func (t FuncSlice) Calli(id int, params ...interface{}) []interface{} {
	var result []interface{}
	r := t.Call(id, params...)
	if len(r) > 0 {
		result = make([]interface{}, len(r))
		for i := range r {
			result = append(result, r[i].Interface())
		}
	}
	return result
}

// Call function from a function map
func (t FuncMap) Call(name string, params ...interface{}) []reflect.Value {
	f, e := t[name]
	if !e {
		err.Panic(err.New("The function not found", 0))
	}
	return call(f, params...)
}

// Call function from a function slice
func (t FuncSlice) Call(id int, params ...interface{}) []reflect.Value {
	if len(t) <= id {
		err.Panic(err.New("The function not found", 0))
	}
	return call(t[id], params...)
}

func call(f reflect.Value, params ...interface{}) []reflect.Value {
	if len(params) != f.Type().NumIn() {
		err.Panic(err.New("The number of params is not adapted.", 0))
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in)
}
