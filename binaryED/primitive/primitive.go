// (c) 2013 Cergoo
// under terms of ISC license

// Package primitive it's a binary encode/decode primitive elementary implementation
package primitive

import (
	"encoding/binary"
	"math"
	"reflect"
	"time"
)

type (
	// IBuf interface of a buffer with support reserve
	IBuf interface {
		Reserve(int) []byte
		WriteByte(byte) error
		Write([]byte) (int, error)
		ReadNext(n int) ([]byte, error)
		ReadByte() (byte, error)
		ReadWriteReset()
	}
)

// Words size in bytes.
const (
	WORD16 = 2
	WORD32 = 4
	WORD64 = 8
)

var (
	// Pack it's a main coder
	Pack = binary.LittleEndian
	// TimeType reflection type time
	TimeType = reflect.TypeOf(time.Time{})
)

/* Encoders */

// PutUint8 encode a uint8 into buf
func PutUint8(buf IBuf, val uint8) {
	buf.WriteByte(val)
}

// PutUint16 encode a uint16 into buf
func PutUint16(buf IBuf, val uint16) {
	Pack.PutUint16(buf.Reserve(WORD16), val)
}

// PutUint32 encode a uint32 into buf
func PutUint32(buf IBuf, val uint32) {
	Pack.PutUint32(buf.Reserve(WORD32), val)
}

// PutUint64 encode a uint64 into buf
func PutUint64(buf IBuf, val uint64) {
	Pack.PutUint64(buf.Reserve(WORD64), val)
}

// PutFloat64 encode a float64 into buf
func PutFloat64(buf IBuf, val float64) {
	Pack.PutUint64(buf.Reserve(WORD64), math.Float64bits(val))
}

// PutFloat32 encode a float32 into buf
func PutFloat32(buf IBuf, val float32) {
	Pack.PutUint32(buf.Reserve(WORD32), math.Float32bits(val))
}

// PutInt8 encode a int8 into buf
func PutInt8(buf IBuf, val int8) {
	buf.WriteByte(uint8(val))
}

// PutInt16 encode a int16 into buf
func PutInt16(buf IBuf, val int16) {
	Pack.PutUint16(buf.Reserve(WORD16), uint16(val))
}

// PutInt32 encode a int32 into buf
func PutInt32(buf IBuf, val int32) {
	Pack.PutUint32(buf.Reserve(WORD32), uint32(val))
}

// PutInt64 encode a int64 into buf
func PutInt64(buf IBuf, val int64) {
	Pack.PutUint64(buf.Reserve(WORD64), uint64(val))
}

// PutBool encode a bool into buf
func PutBool(buf IBuf, val bool) {
	if val {
		buf.Reserve(1)[0] = 1
		return
	}
	buf.Reserve(1)[0] = 0

}

// PutString encode a bool into buf
func PutString(buf IBuf, val string) {
	Pack.PutUint32(buf.Reserve(WORD32), uint32(val.Len()))
	buf.Write([]byte(val.String()))
}

// PutTime encode a time into buf
func PutTime(buf IBuf, val time.Time) {
	Pack.PutUint64(buf.Reserve(WORD64), uint64(val.UnixNano()))
}

/* Decoders */

// Uint16 decode a uint16 from []byte
func Uint16(b []byte) uint16 {
	return Pack.Uint16(b)
}

// Uint32 decode a uint32 from []byte
func Uint32(b []byte) uint32 {
	return Pack.Uint32(b)
}

// Uint64 decode a uint64 from []byte
func Uint64(b []byte) uint64 {
	return Pack.Uint64(b)
}

// Int16 decode a int16 from []byte
func Int16(b []byte) int16 {
	return int16(Pack.Uint16(b))
}

// Int32 decode a int32 from []byte
func Int32(b []byte) int32 {
	return int32(Pack.Uint32(b))
}

// Int64 decode a int64 from []byte
func Int64(b []byte) int64 {
	return int64(Pack.Uint64(b))
}

// Float32 decode a float32 from []byte
func Float32(b []byte) float32 {
	return math.Float32frombits(Pack.Uint32(b))
}

// Float64 decode a float64 from []byte
func Float64(b []byte) float64 {
	return math.Float64frombits(Pack.Uint64(b))
}

// Bool decode a bool from byte
func Bool(b byte) bool {
	return b != 0
}

// Time decode a time from []byte
func Time(b []byte) time.Time {
	return time.Unix(0, int64(Pack.Uint64(b))).UTC()
}
