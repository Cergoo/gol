package binaryED

import (
	"math"
	"time"
)

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
