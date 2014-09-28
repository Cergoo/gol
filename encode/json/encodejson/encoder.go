// (c) 2014 Cergoo
// under terms of ISC license

package encodejson

import (
	"encoding/json"
	"reflect"
	"strconv"
)

var (
	marshalerType = reflect.TypeOf(new(json.Marshaler)).Elem()
	null          = []byte("null")
)

// Encode encode into buf
func Encode(buf []byte, val interface{}) []byte {
	return encode(reflect.ValueOf(val), buf)
}

func encode(val reflect.Value, buf []byte) []byte {
	// use json.Marshaler implement
	if val.Type().Implements(marshalerType) {
		b, _ := val.Interface().(json.Marshaler).MarshalJSON()
		return append(buf, []byte(b)...)
	}
	switch val.Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		buf = strconv.AppendUint(buf, val.Uint(), 10)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		buf = strconv.AppendInt(buf, val.Int(), 10)
	case reflect.Float32, reflect.Float64:
		buf = strconv.AppendFloat(buf, val.Float(), 'd', -1, 64)
	case reflect.Bool:
		if val.Bool() {
			buf = append(buf, []byte("true")...)
		} else {
			buf = append(buf, []byte("false")...)
		}
	case reflect.String:
		buf = append(buf, '"')
		buf = append(buf, []byte(val.String())...)
		buf = append(buf, '"')
	case reflect.Array:
		buf = append(buf, '[')
		vLen := val.Len()
		if vLen > 0 {
			for i := 0; i < vLen; i++ {
				buf = encode(val.Index(i), buf)
				buf = append(buf, ',')
			}
			buf[len(buf)-1] = ']'
		} else {
			buf = append(buf, ']')
		}
	case reflect.Slice:
		if val.IsNil() {
			return append(buf, null...)
		}
		if val.Type().Elem().Kind() == reflect.Uint8 {
			buf = append(buf, '"')
			buf = append(buf, val.Bytes()...)
			buf = append(buf, '"')
			return buf
		}
		buf = append(buf, '[')
		vLen := val.Len()
		if vLen > 0 {
			for i := 0; i < vLen; i++ {
				buf = encode(val.Index(i), buf)
				buf = append(buf, ',')
			}
			buf[len(buf)-1] = ']'
		} else {
			buf = append(buf, ']')
		}
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return append(buf, null...)
		}
		buf = encode(val.Elem(), buf)
	case reflect.Struct:
		buf = append(buf, '{')
		ln := val.NumField()
		vType := val.Type()
		for i := 0; i < ln; i++ {
			// Ignore private fields
			if vType.Field(i).PkgPath != "" {
				continue
			}
			buf = append(buf, '"')
			buf = append(buf, []byte(vType.Field(i).Name)...)
			buf = append(buf, '"', ':')
			buf = encode(val.Field(i), buf)
			buf = append(buf, ',')
		}
		if buf[len(buf)-1] == ',' {
			buf[len(buf)-1] = '}'
		} else {
			buf = append(buf, '}')
		}
	case reflect.Map:
		if val.IsNil() {
			return append(buf, null...)
		}
		keys := val.MapKeys()
		if len(keys) == 0 {
			return append(buf, '{', '}')
		}
		// as object
		if val.Type().Key().Kind() == reflect.String {
			buf = append(buf, '{')
			for _, k := range keys {
				buf = append(buf, '"')
				buf = append(buf, k.Bytes()...)
				buf = append(buf, '"', ':')
				buf = encode(val.MapIndex(k), buf)
				buf = append(buf, ',')
			}
			// remove last ','
			buf[len(buf)-1] = '}'
			return buf
		}
		// as array
		buf = append(buf, '[')
		for _, k := range keys {
			buf = encode(k, buf)
			buf = append(buf, ',')
			buf = encode(val.MapIndex(k), buf)
			buf = append(buf, ',')
		}
		// remove last ','
		buf[len(buf)-1] = ']'
	case reflect.Invalid:
		buf = append(buf, null...)
	}
	return buf
}
