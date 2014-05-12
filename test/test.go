/*
	test pkg is a simple assertion wrapper for Go's built in "testing" package
	fork jmervine/GoT
	(c) 2013 Cergoo
	under terms of ISC license
*/

package test

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
)

type TT struct {
	t *testing.T
}

// constructor
func New(t *testing.T) *TT {
	return &TT{t: t}
}

// helper error generator
func (t *TT) error(args ...interface{}) {
	m := args[0].(string)

	depth := 2
	if len(args) == 2 {
		depth = args[1].(int)
	}

	var err string
	if _, file, line, ok := runtime.Caller(depth); ok {
		err = fmt.Sprintf("> %s:%d: %s", path.Base(file), line, m)
	} else {
		err = fmt.Sprintf("> ???:-1: %s", m)
	}
	t.t.Error(err)
}

// Equivalent test
func (t *TT) Eq(id string, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.error(fmt.Sprintf("eq: %s %d %d", id, a, b))
	}
}

// No Equivalent test
func (t *TT) NoEq(id string, a, b interface{}) {
	if reflect.DeepEqual(a, b) {
		t.error(fmt.Sprintf("noeq: %s %d %d", id, a, b))
	}
}
