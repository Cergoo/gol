// (c) 2014 Cergoo
// under terms of ISC license

// Package human it's a formatters for units to human friendly sizes
package human

import (
	"math"
	"strconv"
)

var (
	log1024 = math.Log(1024)
)

// GetBytenHumanize return function from humanize byten value
func GetBytenHumanize(names []string) func(v uint64) string {
	return func(v uint64) string {
		r, i, _ := Byten(v)
		return r + names[i]
	}
}

// Byten return human format and range value from i18n
func Byten(v uint64) (string, uint8, float64) {
	if v < 1024 {
		return strconv.FormatUint(v, 10), 0, float64(v)
	}
	i := math.Floor(math.Log(float64(v)) / log1024)
	r := float64(v) / math.Pow(1024, i)
	var f int
	if r < 10 {
		f = 1
	}
	return strconv.FormatFloat(r, 'f', f, 64), uint8(i), r
}
