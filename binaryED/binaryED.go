/*
	binary Encode Decode implementation
  	(c) 2013 Cergoo
	under terms of ISC license
*/

package binaryED

import (
	"encoding/binary"
	"math"
	"time"
)

type (

	// Buffer with support reserve
	IBuffer interface {
		Reserve(int) []byte
		WriteByte(byte) error
	}
)

// Words size in bytes.
const (
	WORD16 = 2
	WORD32 = 4
	WORD64 = 8
)

var Pack = binary.LittleEndian

/* Encoders */

func PutUint8(buf IBuffer, val uint8) {
	buf.WriteByte(val)
}

func PutUint16(buf IBuffer, val uint16) {
	Pack.PutUint16(buf.Reserve(WORD16), val)
}

func PutUint32(buf IBuffer, val uint32) {
	Pack.PutUint32(buf.Reserve(WORD32), val)
}

func PutUint64(buf IBuffer, val uint64) {
	Pack.PutUint64(buf.Reserve(WORD64), val)
}

func PutFloat64(buf IBuffer, val float64) {
	Pack.PutUint64(buf.Reserve(WORD64), math.Float64bits(val))
}

func PutFloat32(buf IBuffer, val float32) {
	Pack.PutUint32(buf.Reserve(WORD32), math.Float32bits(val))
}

func PutInt8(buf IBuffer, val int8) {
	buf.WriteByte(uint8(val))
}

func PutInt16(buf IBuffer, val int16) {
	Pack.PutUint16(buf.Reserve(WORD16), uint16(val))
}

func PutInt32(buf IBuffer, val int32) {
	Pack.PutUint32(buf.Reserve(WORD32), uint32(val))
}

func PutInt64(buf IBuffer, val int64) {
	Pack.PutUint64(buf.Reserve(WORD64), uint64(val))
}

func PutBool(buf IBuffer, val bool) {
	if val {
		buf.Reserve(1)[0] = 1
		return
	}
	buf.Reserve(1)[0] = 0

}

func PutTime(buf IBuffer, val time.Time) {
	Pack.PutUint64(buf.Reserve(WORD64), uint64(val.UnixNano()/1e6))
}

/* Decoders */

func Uint16(b []byte) uint16 {
	return Pack.Uint16(b)
}

func Uint32(b []byte) uint32 {
	return Pack.Uint32(b)
}

func Uint64(b []byte) uint64 {
	return Pack.Uint64(b)
}

func Int16(b []byte) int16 {
	return int16(Pack.Uint16(b))
}

func Int32(b []byte) int32 {
	return int32(Pack.Uint32(b))
}

func Int64(b []byte) int64 {
	return int64(Pack.Uint64(b))
}

func Float32(b []byte) float32 {
	return math.Float32frombits(Pack.Uint32(b))
}

func Float64(b []byte) float64 {
	return math.Float64frombits(Pack.Uint64(b))
}

func Bool(b byte) bool {
	return b != 0
}

func Time(b []byte) time.Time {
	return time.Unix(0, int64(Pack.Uint64(b))*1e6).UTC()
}
