package binaryED

import (
	"github.com/Cergoo/gol/fastbuf"
	"github.com/Cergoo/gol/test"
	"testing"
)

type (
	t1 struct {
		F1 int
		F2 string
		f3 int // unexported field no encode no decode
		F4 []int
		F  *t2
	}
	t2 struct {
		F1 string
		F2 string
	}
)

var (
	buf = fastbuf.New(nil)

	inInt    = int(17)
	inStr    = "tested string"
	inStr1   = ""
	inBoolt  = true
	inBoolf  = false
	inSlice  = []string{"1", "2", "3", "nnnn1", "nn2", "", "n1"}
	inStruct = &t1{
		F1: -12,
		F2: "test1",
		f3: 100,
		F4: []int{1, 2, 4},
		F: &t2{
			F1: "test_str1",
			F2: "test_str2",
		},
	}
	inMap  = map[int]string{1: "f1", 2: "f2", 4: "f4"}
	inMap1 = map[int]*t2{1: &t2{"f1", "f2"}, 2: &t2{"f2", "f2"}}
	inMap2 = map[int]t2{1: t2{"f1", "f2"}, 2: t2{"f2", "f2"}}
	inMapN = map[t2]int{t2{"f1", "f2"}: 1, t2{"f2", "f2"}: 2}

	outInt    int
	outStr    string
	outBool   bool
	outSlice  []string
	outStruct = &t1{F: &t2{}, f3: 100}
	outMap    map[int]string
	outMap1   map[int]*t2
	outMap2   map[int]t2
	outMapN   map[t2]int
)

func TestED(t *testing.T) {
	t1 := test.New(t)

	Encode(buf, inInt)
	Decode(buf, &outInt)
	t1.Eq(inInt, outInt)

	Encode(buf, inStr)
	Decode(buf, &outStr)
	t1.Eq(inStr, outStr)

	Encode(buf, inStr1)
	Decode(buf, &outStr)
	t1.Eq(inStr1, outStr)

	Encode(buf, inBoolf)
	Decode(buf, &outBool)
	t1.Eq(inBoolf, outBool)

	Encode(buf, inBoolt)
	Decode(buf, &outBool)
	t1.Eq(inBoolt, outBool)

	Encode(buf, inSlice)
	Decode(buf, &outSlice)
	t1.Eq(inSlice, outSlice)

	Encode(buf, inStruct)
	Decode(buf, outStruct)
	t1.Eq(inStruct, outStruct)

	Encode(buf, inMap)
	Decode(buf, &outMap)
	t1.Eq(inMap, outMap)

	Encode(buf, inMap1)
	Decode(buf, &outMap1)
	t1.Eq(inMap1, outMap1)

	Encode(buf, inMap2)
	Decode(buf, &outMap2)
	t1.Eq(inMap2, outMap2)

	Encode(buf, inMapN)
	Decode(buf, &outMapN)
	t1.Eq(inMapN, outMapN)
}
