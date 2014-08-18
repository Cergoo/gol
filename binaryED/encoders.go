/*
	binary Encode Decode implementation
  	(c) 2013 Cergoo
	under terms of ISC license
*/

package binaryED

import (
	"encoding/binary"
	"github.com/Cergoo/gol/fastbuf"
	"math"
	"time"
)

// Words size in bytes.
const (
	WORD16 = 2
	WORD32 = 4
	WORD64 = 8
)

var Pack = binary.LittleEndian

func PutUint16(buf *fastbuf.Buf, val uint16) {
	Pack.PutUint16(buf.Reserve(WORD16), val)
}

func PutUint32(buf *fastbuf.Buf, val uint32) {
	Pack.PutUint32(buf.Reserve(WORD32), val)
}

func PutUint64(buf *fastbuf.Buf, val uint64) {
	Pack.PutUint64(buf.Reserve(WORD64), val)
}

func PutFloat64(buf *fastbuf.Buf, val float64) {
	Pack.PutUint64(buf.Reserve(WORD64), math.Float64bits(val))
}

func PutFloat32(buf *fastbuf.Buf, val float32) {
	Pack.PutUint32(buf.Reserve(WORD32), math.Float32bits(val))
}

func PutInt16(buf *fastbuf.Buf, val int16) {
	Pack.PutUint16(buf.Reserve(WORD16), uint16(val))
}

func PutInt32(buf *fastbuf.Buf, val int32) {
	Pack.PutUint32(buf.Reserve(WORD32), uint32(val))
}

func PutInt64(buf *fastbuf.Buf, val int64) {
	Pack.PutUint64(buf.Reserve(WORD64), uint64(val))
}

func PutBool(buf *fastbuf.Buf, val bool) {
	if val {
		buf.Reserve(1)[0] = 1
		return
	}
	buf.Reserve(1)[0] = 0

}

func PutTime(buf *fastbuf.Buf, val time.Time) {
	Pack.PutUint64(buf.Reserve(WORD64), uint64(val.UnixNano()/1e6))
}
