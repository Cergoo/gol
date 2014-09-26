package main

import (
	"fmt"
	"github.com/Cergoo/gol/binaryED/fastED/example/encoderdecoder"
	"github.com/Cergoo/gol/binaryED/fastED/example/exportedtypes"
	"github.com/Cergoo/gol/fastbuf"
	//"github.com/davecgh/go-spew/spew"
	"reflect"
)

func main() {
	inVar := new(exportedtypes.T1)
	inVar.N1 = 20
	inVar.N2 = "test string"
	inVar.N3 = 1
	inVar.N4 = new(exportedtypes.T2)
	inVar.N4.N1 = 12
	inVar.N4.N2 = "new test string"
	inVar.N4.N3 = 1
	inVar.N4.N4 = []string{0: "str0", 1: "str1", 4: "str4"}
	inVar.N4.N5 = map[int]string{1: "str1", 10: "str10"}
	inVar.N4.N6 = map[int]*exportedtypes.T3{}
	inVar.N4.N6[1] = &exportedtypes.T3{N: 12}
	inVar.N4.N6[2] = &exportedtypes.T3{N: 2}

	buf := fastbuf.New(nil, 0, nil)

	encoderdecoder.Encode(inVar, buf)
	outVar, e := encoderdecoder.Decode(buf)

	if reflect.DeepEqual(inVar, outVar) {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok", e)
	}

}
