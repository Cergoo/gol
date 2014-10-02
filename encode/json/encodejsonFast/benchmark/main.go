package main

import (
	"encoding/json"
	"fmt"
	"github.com/Cergoo/gol/encode/json/encodejson"
	"github.com/Cergoo/gol/encode/json/encodejsonFast/example/encoderdecoder"
	"github.com/Cergoo/gol/encode/json/encodejsonFast/example/exportedtypes"
	"testing"
)

var (
	e             error
	buf           []byte
	inVar, outVar *exportedtypes.T1
)

func init() {
	inVar = new(exportedtypes.T1)
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

	buf = make([]byte, 0, 64)
}

func main() {
	var (
		std_json, gol_encodejson, gol_encodejsonFast testing.BenchmarkResult
		operation_name                               string
	)

	operation_name = "Encode"
	std_json = testing.Benchmark(Benchmark_std_json)
	gol_encodejson = testing.Benchmark(Benchmark_gol_encodejson)
	gol_encodejsonFast = testing.Benchmark(Benchmark_gol_encodejsonFast)
	fmt.Print(operation_name, "\n", "std_json:", std_json, std_json.MemString(),
		"\n", "gol_encodejson:", gol_encodejson, gol_encodejson.MemString(),
		"\n", "gol_encodejsonFast:", gol_encodejsonFast, gol_encodejsonFast.MemString(), "\n")

}

func Benchmark_std_json(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf, e = json.Marshal(inVar)
		buf = buf[:0]
	}
	fmt.Println(e)
}

func Benchmark_gol_encodejson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf = encodejson.Encode(buf, inVar)
		buf = buf[:0]
	}
}

func Benchmark_gol_encodejsonFast(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf = encoderdecoder.Encode(buf, inVar)
		buf = buf[:0]
	}
}
