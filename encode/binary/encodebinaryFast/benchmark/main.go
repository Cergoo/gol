package main

import (
	"encoding/gob"
	"fmt"
	"github.com/Cergoo/gol/encode/binary/encodebinary"
	"github.com/Cergoo/gol/encode/binary/encodebinaryFast/example/encoderdecoder"
	"github.com/Cergoo/gol/encode/binary/encodebinaryFast/example/exportedtypes"
	"github.com/Cergoo/gol/fastbuf"
	"testing"
)

var (
	buf1, buf2    *fastbuf.Buf
	inVar, outVar *exportedtypes.T1
	decoder       *encodebinary.TDecoder
	gobEncoder    *gob.Encoder
	gobDecoder    *gob.Decoder
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

	buf1 = fastbuf.New(nil, 0, nil)
	buf2 = fastbuf.New(nil, 0, nil)
	decoder = encodebinary.NewDecoder(buf2)
	encodebinary.Encode(buf2, inVar)
	gobEncoder = gob.NewEncoder(buf1)
	gobDecoder = gob.NewDecoder(buf2)
}

func main() {
	var (
		binaryEDt, fastEDt, gobt testing.BenchmarkResult
		operation_name           string
	)

	operation_name = "Encode"
	binaryEDt = testing.Benchmark(Benchmark_BinaryED_Encode)
	fastEDt = testing.Benchmark(Benchmark_FastED_Encode)
	gobt = testing.Benchmark(Benchmark_gob_Encode)
	fmt.Print(operation_name, "\n", "binaryED:", binaryEDt, binaryEDt.MemString(), "\n", "fastED:", fastEDt, fastEDt.MemString(), "\n", "gob:", gobt, gobt.MemString(), "\n")
	operation_name = "Decode"
	binaryEDt = testing.Benchmark(Benchmark_BinaryED_Decode)
	fastEDt = testing.Benchmark(Benchmark_FastED_Decode)
	gobt = testing.Benchmark(Benchmark_gob_Decode)
	fmt.Print(operation_name, "\n", "binaryED:", binaryEDt, binaryEDt.MemString(), "\n", "fastED:", fastEDt, fastEDt.MemString(), "\n", "gob:", gobt, gobt.MemString(), "\n")

}

func Benchmark_BinaryED_Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf1.ReadWriteReset()
		encodebinary.Encode(buf1, inVar)
	}
}

func Benchmark_BinaryED_Decode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		decoder.Decode(&outVar)
		buf2.ReadReset()
	}
}

func Benchmark_FastED_Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf1.ReadWriteReset()
		encoderdecoder.Encode(buf1, inVar)
	}
}

func Benchmark_FastED_Decode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		outVar, _ = encoderdecoder.Decode(buf2)
		buf2.ReadReset()
	}
}

func Benchmark_gob_Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf1.ReadWriteReset()
		gobEncoder.Encode(inVar)
	}
}

func Benchmark_gob_Decode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// gobDecoder.Decode(&outVar1)
		buf2.ReadReset()
	}
}
