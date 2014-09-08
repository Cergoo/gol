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

/* Encoders */

// Encode encode value to binary
func Encode(buf IBuf, val interface{}) {
	encodeField(buf, reflect.ValueOf(val))
}

func encodeField(buf IBuf, val reflect.Value) {
	val = reflect.Indirect(val)
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
	case reflect.Bool:
		PutBool(buf, val.Bool())
	case reflect.String:
		vLen := val.Len()
		Pack.PutUint32(buf.Reserve(WORD32), uint32(vLen))
		buf.Write([]byte(val.String()))
	case reflect.Slice, reflect.Array:
		vLen := val.Len()
		Pack.PutUint32(buf.Reserve(WORD32), uint32(vLen))
		if val.Type().Elem().Kind() == reflect.Uint8 {
			buf.Write(val.Interface().([]byte))
		} else {
			for i := 0; i < vLen; i++ {
				encodeField(buf, val.Index(i))
			}
		}
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
		vLen := val.Len()
		Pack.PutUint32(buf.Reserve(WORD32), uint32(vLen))
		keys := val.MapKeys()
		for _, k := range keys {
			encodeField(buf, k)
			encodeField(buf, val.MapIndex(k))
		}
	}
}

/* Decoders */

// Decode decode value to binary
func Decode(buf IBuf, val interface{}) {
	decodeField(buf, reflect.ValueOf(val))
}

func decodeField(buf IBuf, val reflect.Value) (e error) {
	var (
		part []byte
		bt   byte
	)
	val = reflect.Indirect(val)
	switch val.Kind() {
	case reflect.Uint8:
		bt, e = buf.ReadByte()
		if e == nil {
			*(*uint8)(val.Ptr()) = uint8(bt)
		}
	case reflect.Uint16:
		part, e = buf.ReadNext(WORD16)
		if e == nil {
			*(*uint16)(val.Ptr()) = Pack.Uint16(part)
		}
	case reflect.Uint32:
		part, e = buf.ReadNext(WORD32)
		if e == nil {
			*(*uint32)(val.Ptr()) = Pack.Uint32(part)
		}
	case reflect.Uint64:
		part, e = buf.ReadNext(WORD64)
		if e == nil {
			*(*uint64)(val.Ptr()) = Pack.Uint64(part)
		}
	case reflect.Uint:
		part, e = buf.ReadNext(WORD64)
		if e == nil {
			*(*uint)(val.Ptr()) = uint(Pack.Uint64(part))
		}
	case reflect.Int8:
		bt, e = buf.ReadByte()
		if e == nil {
			*(*int8)(val.Ptr()) = int8(bt)
		}
	case reflect.Int16:
		part, e = buf.ReadNext(WORD16)
		if e == nil {
			*(*int16)(val.Ptr()) = int16(Pack.Uint16(part))
		}
	case reflect.Int32:
		part, e = buf.ReadNext(WORD32)
		if e == nil {
			*(*int32)(val.Ptr()) = int32(Pack.Uint32(part))
		}
	case reflect.Int64:
		part, e = buf.ReadNext(WORD64)
		if e == nil {
			*(*int64)(val.Ptr()) = int64(Pack.Uint64(part))
		}
	case reflect.Int:
		part, e = buf.ReadNext(WORD64)
		if e == nil {
			*(*int)(val.Ptr()) = int(Pack.Uint64(part))
		}
	case reflect.Float32:
		part, e = buf.ReadNext(WORD32)
		if e == nil {
			*(*float32)(val.Ptr()) = math.Float32frombits(Pack.Uint32(part))
		}
	case reflect.Bool:
		bt, e = buf.ReadByte()
		if e == nil {
			if bt != 0 {
				*(*bool)(val.Ptr()) = true
			}
		}
	case reflect.String:
		part, e = buf.ReadNext(WORD32)
		if e == nil {
			ln := int(Pack.Uint32(part))
			if ln == 0 {
				*(*string)(val.Ptr()) = ""
				return
			}
			part, e = buf.ReadNext(ln)
			if e == nil {
				*(*string)(val.Ptr()) = string(part)
			}
		}
	case reflect.Slice:
		part, e = buf.ReadNext(WORD32)
		if e != nil {
			return
		}
		ln := int(Pack.Uint32(part))
		if val.Type().Elem().Kind() == reflect.Uint8 {
			part, e = buf.ReadNext(ln)
			if e != nil {
				return
			}
			*(*[]byte)(val.Ptr()) = part
		} else {
			if val.IsNil() {
				val.Set(reflect.MakeSlice(val.Type(), ln, ln))
			}
			for i := 0; i < ln; i++ {
				decodeField(buf, val.Index(i))
			}
		}
	case reflect.Struct:
		vType := val.Type()
		if vType == TimeType {
			part, e = buf.ReadNext(WORD64)
			if e != nil {
				return
			}
			val.Set(reflect.ValueOf(Time(part)))
		} else {
			ln := val.NumField()
			for i := 0; i < ln; i++ {
				if vType.Field(i).PkgPath != "" {
					continue
				}
				decodeField(buf, val.Field(i))
			}
		}
	case reflect.Map:
		part, e = buf.ReadNext(WORD32)
		if e != nil {
			return
		}
		ln := int(Pack.Uint32(part))
		vtype := val.Type()
		if val.IsNil() {
			val.Set(reflect.MakeMap(vtype))
		}
		var (
			mapKey, mapVal reflect.Value
			ptrVal         bool
		)
		mapValType := vtype.Elem()
		if mapValType.Kind() == reflect.Ptr {
			ptrVal = true
			mapValType = mapValType.Elem()
		}
		for i := 0; i < ln; i++ {
			mapVal = reflect.New(mapValType)
			if !ptrVal {
				mapVal = mapVal.Elem()
			}
			mapKey = reflect.New(vtype.Key()).Elem()
			decodeField(buf, mapKey)
			decodeField(buf, mapVal)
			val.SetMapIndex(mapKey, mapVal)
		}
	}
	return
}
