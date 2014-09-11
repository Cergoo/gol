package binaryED

import (
	//"fmt"
	"github.com/Cergoo/gol/fastbuf"
	"github.com/Cergoo/gol/test"
	//"github.com/davecgh/go-spew/spew"
	"testing"
)

type (
	t1 struct {
		F1 int
		F2 string
		f3 int // unexported field no encode no decode
		F4 []int
		F  *t2
		F5 int
	}
	t2 struct {
		F1 string
		F  []int
		F2 string
	}
	tt struct {
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
		F5: 12,
	}
	inMap          = map[int]string{1: "f1", 2: "f2", 4: "f4"}
	inMap1         = map[int]*t2{1: &t2{"f1", nil, "f2"}, 2: &t2{"f2", nil, "f2"}}
	inMap2         = map[int]t2{1: t2{"f1", []int{}, "f2"}, 2: t2{"f2", []int{}, "f2"}}
	inMapN         = map[tt]int{tt{"f1", "f2"}: 1, tt{"f2", "f2"}: 2}
	inMapInterface = map[int]interface{}{1: 1, 2: tt{"f2", "f2"}, 3: &t2{"f2", nil, "f2"}}

	outInt          int
	outStr          string
	outBool         bool
	outSlice        []string
	outStruct       = &t1{f3: 100}
	outMap          map[int]string
	outMap1         map[int]*t2
	outMap2         map[int]t2
	outMapN         map[tt]int
	outMapInterface map[int]interface{}
)

func TestED(t *testing.T) {
	t1 := test.New(t)

	Decoder := NewDecoder(buf)
	Decoder.Register(&t2{}, tt{})

	Encode(buf, inInt)
	Decoder.Decode(&outInt)
	t1.Eq(inInt, outInt)

	Encode(buf, inStr)
	Decoder.Decode(&outStr)
	t1.Eq(inStr, outStr)

	Encode(buf, inStr1)
	Decoder.Decode(&outStr)
	t1.Eq(inStr1, outStr)

	Encode(buf, inBoolf)
	Decoder.Decode(&outBool)
	t1.Eq(inBoolf, outBool)

	Encode(buf, inBoolt)
	Decoder.Decode(&outBool)
	t1.Eq(inBoolt, outBool)

	Encode(buf, inSlice)
	Decoder.Decode(&outSlice)
	t1.Eq(inSlice, outSlice)

	Encode(buf, inStruct)
	Decoder.Decode(&outStruct)
	t1.Eq(inStruct, outStruct)

	Encode(buf, inMap)
	Decoder.Decode(&outMap)
	t1.Eq(inMap, outMap)

	Encode(buf, inMap1)
	Decoder.Decode(&outMap1)
	t1.Eq(inMap1, outMap1)

	Encode(buf, inMap2)
	Decoder.Decode(&outMap2)
	t1.Eq(inMap2, outMap2)

	Encode(buf, inMapN)
	Decoder.Decode(&outMapN)
	t1.Eq(inMapN, outMapN)

	Encode(buf, inMapInterface)
	Decoder.Decode(&outMapInterface)
	t1.Eq(inMapInterface, outMapInterface)
}
