// (c) 2014 Cergoo
// under terms of ISC license

// Package tmpName it's a temporary name generator. Thread-safe.
package tmpName

import (
	"github.com/Cergoo/gol/counter"
	"strconv"
)

const (
	tmpname = "tmp"
)

type (
	// TtmpName it's a main structure
	TtmpName struct {
		count counter.TCounter
	}
)

// New create new tmp name generator
func New() *TtmpName {
	return &TtmpName{count: counter.TCounter{}}
}

// Get get new tmp name
func (t *TtmpName) Get() string {
	return tmpname + strconv.FormatUint(t.count.Inc(), 10)
}

// Clear clear generator
func (t *TtmpName) Clear() {
	t.count.Set(0)
}
