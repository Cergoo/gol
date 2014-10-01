package main

import (
	"fmt"
	"github.com/Cergoo/gol/encode/json/encodejson"
	"github.com/Cergoo/gol/encode/json/encodejsonFast/example/encoderdecoder"
	"github.com/Cergoo/gol/encode/json/encodejsonFast/example/exportedtypes"
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

	buf := make([]byte, 0, 64)
	buf1 := make([]byte, 0, 64)
	buf = encoderdecoder.Encode(buf, inVar)
	buf1 = encodejson.Encode(buf1, inVar)

	fmt.Println("ok", "\n", string(buf), "\n", string(buf1))
}
