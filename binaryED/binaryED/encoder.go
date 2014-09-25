// (c) 2014 Cergoo
// under terms of ISC license

package binaryED

import (
	//"fmt"
	. "github.com/Cergoo/gol/binaryED/primitive"
	"math"
	"reflect"
	"time"
)

// Encode encode value to binary
func Encode(buf IBuf, val interface{}) {
	encodeField(buf, reflect.ValueOf(val))
}

func encodeField(buf IBuf, val reflect.Value) {
	switch val.Kind() {
	case reflect.Uint8:
		buf.WriteByte(uint8(val.Uint()))
	case reflect.Uint16:
		Pack.PutUint16(buf.Reserve(WORD16), uint16(val.Uint()))
	case reflect.Uint32:
		Pack.PutUint32(buf.Reserve(WORD32), uint32(val.Uint()))
	case reflect.Uint64, reflect.Uint:
		Pack.PutUint64(buf.Reserve(WORD64), val.Uint())
	case reflect.Int8:
		buf.WriteByte(uint8(val.Int()))
	case reflect.Int16:
		Pack.PutUint16(buf.Reserve(WORD16), uint16(val.Int()))
	case reflect.Int32:
		Pack.PutUint32(buf.Reserve(WORD32), uint32(val.Int()))
	case reflect.Int64, reflect.Int:
		Pack.PutUint64(buf.Reserve(WORD64), uint64(val.Int()))
	case reflect.Float32:
		Pack.PutUint32(buf.Reserve(WORD32), math.Float32bits(float32(val.Float())))
	case reflect.Float64:
		Pack.PutUint64(buf.Reserve(WORD64), math.Float64bits(val.Float()))
	case reflect.Complex64:
		PutComplex64(buf, complex64(val.Complex()))
	case reflect.Complex128:
		PutComplex128(buf, val.Complex())
	case reflect.Bool:
		PutBool(buf, val.Bool())
	case reflect.String:
		Pack.PutUint32(buf.Reserve(WORD32), uint32(val.Len()))
		buf.Write([]byte(val.String()))
	case reflect.Slice:
		if val.IsNil() {
			buf.WriteByte(0)
			return
		}
		buf.WriteByte(1)
		vLen := val.Len()
		Pack.PutUint32(buf.Reserve(WORD32), uint32(vLen))
		if val.Type().Elem().Kind() == reflect.Uint8 {
			buf.Write(val.Bytes())
		} else {
			for i := 0; i < vLen; i++ {
				encodeField(buf, val.Index(i))
			}
		}
	case reflect.Array:
		vLen := val.Len()
		Pack.PutUint32(buf.Reserve(WORD32), uint32(vLen))
		for i := 0; i < vLen; i++ {
			encodeField(buf, val.Index(i))
		}
	case reflect.Ptr:
		if val.IsNil() {
			buf.WriteByte(0)
			return
		}
		buf.WriteByte(1)
		val = val.Elem()
		encodeField(buf, val)
	case reflect.Struct:
		vType := val.Type()
		if vType == TimeType {
			PutTime(buf, val.Interface().(time.Time))
		} else {
			ln := val.NumField()
			for i := 0; i < ln; i++ {
				// Ignore private fields
				if vType.Field(i).PkgPath != "" {
					continue
				}
				encodeField(buf, val.Field(i))
			}
		}
	case reflect.Map:
		if val.IsNil() {
			buf.WriteByte(0)
			return
		}
		buf.WriteByte(1)
		vLen := val.Len()
		Pack.PutUint32(buf.Reserve(WORD32), uint32(vLen))
		keys := val.MapKeys()
		for _, k := range keys {
			encodeField(buf, k)
			encodeField(buf, val.MapIndex(k))
		}
	case reflect.Interface:
		if val.IsNil() {
			buf.WriteByte(0)
			return
		}
		buf.WriteByte(1)
		val = val.Elem()
		tpName := val.Type().String()
		buf.WriteByte(uint8(len(tpName)))
		buf.Write([]byte(tpName))
		encodeField(buf, val)
	}
}
