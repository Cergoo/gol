// (c) 2014 Cergoo
// under terms of ISC license

package encodeBinaryV1

import (
	//"fmt"
	"reflect"
	"time"
	"unsafe"
)

type (
	iEndianedEncoder interface {
		PutUint16(buf []byte, val unsafe.Pointer) []byte
		PutUint32(buf []byte, val unsafe.Pointer) []byte
		PutUint64(buf []byte, val unsafe.Pointer) []byte
		PutUint16Array(buf []byte, val []uint16) []byte
		PutUint32Array(buf []byte, val []uint32) []byte
		PutUint64Array(buf []byte, val []uint64) []byte
	}

	// iEndianedEncoder implementation
	tBigEndianEncoder bool
	tLitEndianEncoder bool

	TEncoder struct {
		iEndianedEncoder
	}
)

var (
	// TimeType reflection type time
	TimeType = reflect.TypeOf(time.Time{})
	ByteType = reflect.TypeOf([]byte(nil))
)

// EndianBig get current endian, if Big return true
func EndianBig() bool {
	var x uint16 = 0x0102
	if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
		return true // big
	}
	return false // litle
}

// NewEncoder create new encoder
func NewEncoder() TEncoder {
	if EndianBig() {
		return TEncoder{tBigEndianEncoder(false)}
	}
	return TEncoder{tLitEndianEncoder(false)}
}

/* Encoders tBigEndian */

func (t tBigEndianEncoder) PutUint16(buf []byte, val unsafe.Pointer) []byte {
	b := *(*[2]byte)(val)
	return append(buf, b[1], b[0])
}

func (t tBigEndianEncoder) PutUint32(buf []byte, val unsafe.Pointer) []byte {
	b := *(*[4]byte)(val)
	return append(buf, b[3], b[2], b[1], b[0])
}

func (t tBigEndianEncoder) PutUint64(buf []byte, val unsafe.Pointer) []byte {
	b := *(*[8]byte)(val)
	return append(buf, b[7], b[6], b[5], b[4], b[3], b[2], b[1], b[0])
}

func (t tBigEndianEncoder) PutUint16Array(buf []byte, val []uint16) []byte {
	var b [2]byte
	for i := range val {
		b = *(*[2]byte)(unsafe.Pointer(&val[i]))
		buf = append(buf, b[1], b[0])
	}
	return buf
}

func (t tBigEndianEncoder) PutUint32Array(buf []byte, val []uint32) []byte {
	var b [4]byte
	for i := range val {
		b = *(*[4]byte)(unsafe.Pointer(&val[i]))
		buf = append(buf, b[3], b[2], b[1], b[0])
	}
	return buf
}

func (t tBigEndianEncoder) PutUint64Array(buf []byte, val []uint64) []byte {
	var b [8]byte
	for i := range val {
		b = *(*[8]byte)(unsafe.Pointer(&val[i]))
		buf = append(buf, b[7], b[6], b[5], b[4], b[3], b[2], b[1], b[0])
	}
	return buf
}

/* Encoders tLitEndian */

func (t TEncoder) PutUint(buf []byte, val unsafe.Pointer) []byte {
	n := uint64(*(*uint)(val))
	b := *(*[8]byte)(unsafe.Pointer(&n))
	return append(buf, b[:]...)
}

func (t tLitEndianEncoder) PutUint16(buf []byte, val unsafe.Pointer) []byte {
	b := *(*[2]byte)(val)
	return append(buf, b[:]...)
}

func (t tLitEndianEncoder) PutUint32(buf []byte, val unsafe.Pointer) []byte {
	b := *(*[4]byte)(val)
	return append(buf, b[:]...)
}

func (t tLitEndianEncoder) PutUint64(buf []byte, val unsafe.Pointer) []byte {
	b := *(*[8]byte)(val)
	return append(buf, b[:]...)
}

func (t tLitEndianEncoder) PutUint16Array(buf []byte, val []uint16) []byte {
	var b [2]byte
	for i := range val {
		b = *(*[2]byte)(unsafe.Pointer(&val[i]))
		buf = append(buf, b[:]...)
	}
	return buf
}

func (t tLitEndianEncoder) PutUint32Array(buf []byte, val []uint32) []byte {
	var b [4]byte
	for i := range val {
		b = *(*[4]byte)(unsafe.Pointer(&val[i]))
		buf = append(buf, b[:]...)
	}
	return buf
}

func (t tLitEndianEncoder) PutUint64Array(buf []byte, val []uint64) []byte {
	var b [8]byte
	for i := range val {
		b = *(*[8]byte)(unsafe.Pointer(&val[i]))
		buf = append(buf, b[:]...)
	}
	return buf
}

/* Encoders TEncoder */

func (t TEncoder) PutComplex64(buf []byte, val complex64) []byte {
	r := real(val)
	i := imag(val)
	buf = t.PutUint32(buf, unsafe.Pointer(&r))
	buf = t.PutUint32(buf, unsafe.Pointer(&i))
	return buf
}

func (t TEncoder) PutComplex128(buf []byte, val complex128) []byte {
	r := real(val)
	i := imag(val)
	buf = t.PutUint64(buf, unsafe.Pointer(&r))
	buf = t.PutUint64(buf, unsafe.Pointer(&i))
	return buf
}

// PutBool encode a bool into buf
func PutBool(buf []byte, val bool) []byte {
	if val {
		return append(buf, 1)
	}
	return append(buf, 0)
}

// PutTime encode a time into buf
func (t TEncoder) PutTime(buf []byte, val time.Time) []byte {
	n := int64(val.UnixNano())
	return t.PutUint64(buf, unsafe.Pointer(&n))
}

// Encode encode value to binary
func (t TEncoder) Encode(buf []byte, val interface{}) []byte {
	return t.encode(buf, reflect.ValueOf(val))
}

func (t TEncoder) encode(buf []byte, val reflect.Value) []byte {
	switch val.Kind() {
	case reflect.Uint8, reflect.Int8:
		buf = append(buf, *(*uint8)(val.GetPtr()))
	case reflect.Uint16, reflect.Int16:
		buf = t.PutUint16(buf, val.GetPtr())
	case reflect.Uint32, reflect.Int32, reflect.Float32:
		buf = t.PutUint32(buf, val.GetPtr())
	case reflect.Uint64, reflect.Int64, reflect.Float64:
		buf = t.PutUint64(buf, val.GetPtr())
	case reflect.Uint, reflect.Int:
		buf = t.PutUint(buf, val.GetPtr())
	case reflect.Complex64:
		buf = t.PutComplex64(buf, complex64(val.Complex()))
	case reflect.Complex128:
		buf = t.PutComplex128(buf, val.Complex())
	case reflect.Bool:
		buf = PutBool(buf, val.Bool())
	case reflect.String:
		ln := uint32(val.Len())
		buf = t.PutUint32(buf, unsafe.Pointer(&ln))
		buf = append(buf, []byte(val.String())...)
	case reflect.Slice:
		if val.IsNil() {
			return append(buf, uint8(0))
		}
		buf = append(buf, uint8(1))

		vLen := uint32(val.Len())
		buf = t.PutUint32(buf, unsafe.Pointer(&vLen))
		switch val.Type().Elem().Kind() {
		case reflect.Uint8, reflect.Int8:
			buf = append(buf, *(*[]byte)(val.GetPtr())...)
		case reflect.Uint16, reflect.Int16:
			buf = t.PutUint16Array(buf, *(*[]uint16)(val.GetPtr()))
		case reflect.Uint32, reflect.Int32, reflect.Float32:
			buf = t.PutUint32Array(buf, *(*[]uint32)(val.GetPtr()))
		case reflect.Uint64, reflect.Int64, reflect.Float64:
			buf = t.PutUint64Array(buf, *(*[]uint64)(val.GetPtr()))
		default:
			for i := 0; i < int(vLen); i++ {
				buf = t.encode(buf, val.Index(i))
			}
		}
	case reflect.Array:
		vLen := uint32(val.Len())
		buf = t.PutUint32(buf, unsafe.Pointer(&vLen))
		for i := 0; i < int(vLen); i++ {
			buf = t.encode(buf, val.Index(i))
		}
	case reflect.Ptr:
		if val.IsNil() {
			return append(buf, uint8(0))
		}
		buf = append(buf, uint8(1))
		buf = t.encode(buf, val.Elem())
	case reflect.Struct:
		vType := val.Type()
		if vType == TimeType {
			buf = t.PutTime(buf, val.Interface().(time.Time))
		} else {
			ln := val.NumField()
			for i := 0; i < ln; i++ {
				// Ignore private fields
				if vType.Field(i).PkgPath != "" {
					continue
				}
				buf = t.encode(buf, val.Field(i))
			}
		}
	case reflect.Map:
		if val.IsNil() {
			return append(buf, uint8(0))
		}
		buf = append(buf, uint8(1))
		ln := uint32(val.Len())
		buf = t.PutUint32(buf, unsafe.Pointer(&ln))
		keys := val.MapKeys()
		for _, k := range keys {
			buf = t.encode(buf, k)
			buf = t.encode(buf, val.MapIndex(k))
		}
	case reflect.Interface:
		if val.IsNil() {
			return append(buf, uint8(0))
		}
		buf = append(buf, uint8(1))
		val = val.Elem()
		tpName := val.Type().String()
		buf = append(buf, uint8(len(tpName)))
		buf = append(buf, []byte(tpName)...)
		buf = t.encode(buf, val)
	}
	return buf
}
