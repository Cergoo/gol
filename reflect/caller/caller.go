// (c) 2013 Cergoo
// under terms of ISC license

// Package caller it's universal caller of functions
package caller

import (
	"github.com/Cergoo/gol/err"
	"reflect"
)

const (
	errNotFunction            = "It's not function."
	errFunctionNotFound       = "The function not found."
	errNumberParamsNotAdapted = "The number of params is not adapted."
)

type (
  // FuncMap it's universal caller of functions map contain
  FuncMap   map[string]reflect.Value
  // FuncSlice it's universal caller of functions slice contain
	FuncSlice []reflect.Value
)

// Add add to function map
func (t FuncMap) Add(name string, f interface{}) {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		err.Panic(err.New(errNotFunction, 0))
	}
	t[name] = v
}

// Add add to function slice, return element id
func (t *FuncSlice) Add(f interface{}) int {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		err.Panic(err.New(errNotFunction, 0))
	}
	*t = append(*t, v)
	return len(*t) - 1
}

// Calli call and return interface{}
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

// Calli call and return interface{}
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

// Call call function from a function map
func (t FuncMap) Call(name string, params ...interface{}) []reflect.Value {
	f, e := t[name]
	if !e {
		err.Panic(err.New(errFunctionNotFound, 0))
	}
	return call(f, params...)
}

// Call call function from a function slice
func (t FuncSlice) Call(id int, params ...interface{}) []reflect.Value {
	if len(t) <= id {
		err.Panic(err.New(errFunctionNotFound, 0))
	}
	return call(t[id], params...)
}

func call(f reflect.Value, params ...interface{}) []reflect.Value {
	if len(params) != f.Type().NumIn() {
		err.Panic(err.New(errNumberParamsNotAdapted, 0))
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in)
}
