// (c) 2014 Cergoo
// under terms of ISC license

// Package human it's a formatters for units to human friendly sizes
package human

import (
	"math"
	"strconv"
)

var (
	vbyten  = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	log1024 = math.Log(1024)
)

// HumanByten1 return human format and range value from i18n
func Byten1(v uint64) (string, uint8) {
	if v < 1024 {
		return strconv.FormatUint(v, 10), 0
	}
	i := math.Floor(math.Log(float64(v)) / log1024)
	r := float64(v) / math.Pow(1024, i)
	var f int
	if r < 10 {
		f = 1
	}
	return strconv.FormatFloat(r, 'f', f, 64), uint8(i)
}

// HumanByten return human format and default description (english)
func Byten(v uint64) string {
	r, i := Byten1(v)
	return r + vbyten[i]
}
