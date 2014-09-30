// (c) 2014 Cergoo
// under terms of ISC license

package encodebinary

import (
	"errors"
	. "github.com/Cergoo/gol/encode/binary/primitive"
	"math"
	"reflect"
	"time"
)

type (
	// TDecoder it's a decoder main structure
	TDecoder struct {
		typeBase map[string]reflect.Type
		buf      IBuf
	}
)

// New construct new decoder
func NewDecoder(buf IBuf) *TDecoder {
	t := &TDecoder{buf: buf, typeBase: make(map[string]reflect.Type)}
	t.register()
	return t
}

// Register register type for decode from a interface{} type receiver
func (t *TDecoder) Register(v ...interface{}) {
	for i := range v {
		t.typeBase[reflect.TypeOf(v[i]).String()] = reflect.TypeOf(v[i])
	}
}

// Preregister elementary
func (t *TDecoder) register() {
	t.Register(uint8(0), uint16(0), uint32(0), uint64(0), uint(0))
	t.Register(int8(0), int16(0), int32(0), int64(0), int(0))
	t.Register(float32(0), float64(0))
	t.Register(complex64(0i), complex128(0i))
	t.Register(string(""), time.Time{}, false)
	t.Register([]uint8{}, []uint16{}, []uint32{}, []uint64{}, []uint{})
	t.Register([]int8{}, []int16{}, []int32{}, []int64{}, []int{})
	t.Register([]float32{}, []float64{})
	t.Register([]complex64{}, []complex128{})
	t.Register([]string{}, []time.Time{}, []bool(nil))
}

// Decode decode value fom binary to receiver val
func (t *TDecoder) Decode(val interface{}) error {
	return t.decode(reflect.ValueOf(val).Elem())
}

func (t *TDecoder) decode(val reflect.Value) (e error) {
	var (
		part []byte
		bt   byte
	)
	switch val.Kind() {
	case reflect.Uint8:
		bt, e = t.buf.ReadByte()
		if e == nil {
			*(*uint8)(val.Ptr()) = uint8(bt)
		}
	case reflect.Uint16:
		part, e = t.buf.ReadNext(WORD16)
		if e == nil {
			*(*uint16)(val.Ptr()) = Pack.Uint16(part)
		}
	case reflect.Uint32:
		part, e = t.buf.ReadNext(WORD32)
		if e == nil {
			*(*uint32)(val.Ptr()) = Pack.Uint32(part)
		}
	case reflect.Uint64:
		part, e = t.buf.ReadNext(WORD64)
		if e == nil {
			*(*uint64)(val.Ptr()) = Pack.Uint64(part)
		}
	case reflect.Uint:
		part, e = t.buf.ReadNext(WORD64)
		if e == nil {
			*(*uint)(val.Ptr()) = uint(Pack.Uint64(part))
		}
	case reflect.Int8:
		bt, e = t.buf.ReadByte()
		if e == nil {
			*(*int8)(val.Ptr()) = int8(bt)
		}
	case reflect.Int16:
		part, e = t.buf.ReadNext(WORD16)
		if e == nil {
			*(*int16)(val.Ptr()) = int16(Pack.Uint16(part))
		}
	case reflect.Int32:
		part, e = t.buf.ReadNext(WORD32)
		if e == nil {
			*(*int32)(val.Ptr()) = int32(Pack.Uint32(part))
		}
	case reflect.Int64:
		part, e = t.buf.ReadNext(WORD64)
		if e == nil {
			*(*int64)(val.Ptr()) = int64(Pack.Uint64(part))
		}
	case reflect.Int:
		part, e = t.buf.ReadNext(WORD64)
		if e == nil {
			*(*int)(val.Ptr()) = int(Pack.Uint64(part))
		}
	case reflect.Float32:
		part, e = t.buf.ReadNext(WORD32)
		if e == nil {
			*(*float32)(val.Ptr()) = math.Float32frombits(Pack.Uint32(part))
		}
	case reflect.Float64:
		part, e = t.buf.ReadNext(WORD64)
		if e == nil {
			*(*float64)(val.Ptr()) = math.Float64frombits(Pack.Uint64(part))
		}
	case reflect.Complex64:
		part, e = t.buf.ReadNext(WORD64)
		if e == nil {
			*(*complex64)(val.Ptr()) = Complex64(part)
		}
	case reflect.Complex128:
		part, e = t.buf.ReadNext(16)
		if e == nil {
			*(*complex128)(val.Ptr()) = Complex128(part)
		}
	case reflect.Bool:
		bt, e = t.buf.ReadByte()
		if bt != 0 {
			*(*bool)(val.Ptr()) = true
		}
	case reflect.String:
		part, e = t.buf.ReadNext(WORD32)
		if e != nil {
			return
		}
		ln := int(Pack.Uint32(part))

		if ln == 0 {
			*(*string)(val.Ptr()) = ""
			return
		}
		part, e = t.buf.ReadNext(ln)
		if e == nil {
			*(*string)(val.Ptr()) = string(part)
		}

	case reflect.Array:
		// get length
		part, e = t.buf.ReadNext(WORD32)
		if e != nil {
			return
		}
		ln := int(Pack.Uint32(part))
		for i := 0; i < ln; i++ {
			e = t.decode(val.Index(i))
			if e != nil {
				return
			}
		}
	case reflect.Slice:
		// check nil
		bt, e = t.buf.ReadByte()
		if e != nil || bt == 0 {
			return
		}
		// get length
		part, e = t.buf.ReadNext(WORD32)
		if e != nil {
			return
		}
		ln := int(Pack.Uint32(part))

		if val.Type().Elem().Kind() == reflect.Uint8 {
			part, e = t.buf.ReadNext(ln)
			if e != nil {
				return
			}
			*(*[]byte)(val.Ptr()) = part
		} else {
			if val.IsNil() || val.Cap() < ln {
				val.Set(reflect.MakeSlice(val.Type(), ln, ln))
			} else {
				val.SetLen(ln)
			}
			for i := 0; i < ln; i++ {
				e = t.decode(val.Index(i))
				if e != nil {
					return
				}
			}
		}
	case reflect.Ptr:
		bt, e = t.buf.ReadByte()
		if e != nil {
			return
		}
		// if nil
		if bt == 0 {
			val.Set(reflect.Zero(val.Type()))
			return
		}
		// interface
		if val.Elem().Kind() == reflect.Invalid {
			val.Set(reflect.New(val.Type().Elem()))
		}

		val = val.Elem()
		e = t.decode(val)
	case reflect.Struct:
		vType := val.Type()
		if vType == TimeType {
			part, e = t.buf.ReadNext(WORD64)
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
				e = t.decode(val.Field(i))
				if e != nil {
					return
				}
			}
		}
	case reflect.Map:
		bt, e = t.buf.ReadByte()
		// if nil
		if e != nil || bt == 0 {
			return
		}
		// get len
		part, e = t.buf.ReadNext(WORD32)
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
		)

		for i := 0; i < ln; i++ {
			mapVal = reflect.New(vtype.Elem()).Elem()
			mapKey = reflect.New(vtype.Key()).Elem()
			e = t.decode(mapKey)
			if e != nil {
				return
			}
			e = t.decode(mapVal)
			if e != nil {
				return
			}
			val.SetMapIndex(mapKey, mapVal)
		}
	case reflect.Interface:
		// check nil
		bt, e = t.buf.ReadByte()
		if e != nil {
			return
		}
		if bt == 0 {
			val.Set(reflect.Zero(val.Type()))
			return
		}

		// get type name
		bt, e = t.buf.ReadByte()
		if e != nil {
			return
		}
		part, e = t.buf.ReadNext(int(bt))
		if e != nil {
			return
		}

		vType, ok := t.typeBase[string(part)]
		if !ok {
			e = errors.New("type '" + string(part) + "' not found, you need to register.")
			return
		}
		valInner := reflect.New(vType).Elem()
		e = t.decode(valInner)
		if e != nil {
			return
		}
		val.Set(valInner)
	}
	return
}
