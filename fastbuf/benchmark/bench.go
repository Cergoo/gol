// Benchmark
package main

import (
	"bytes"
	"fmt"
	"github.com/Cergoo/gol/fastbuf"
	"testing"
)

var (
	fastBuf  = fastbuf.New(nil)
	bytesBuf bytes.Buffer
	r1, r2   testing.BenchmarkResult
	p        = []byte("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")
)

func main() {
	operation_name := "Write"
	r1 = testing.Benchmark(benchmark_fastbuf)
	r2 = testing.Benchmark(benchmark_bytesbuf)
	fmt.Print(operation_name, "\n", "fastbuf:", r1, r1.MemString(), "\n", "bytes.Buffer:", r2, r2.MemString(), "\n")

}

func benchmark_fastbuf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10; i++ {
			fastBuf.Write(p)
		}
		fastBuf.FlushP()
	}
}

func benchmark_bytesbuf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10; i++ {
			bytesBuf.Write(p)
		}
		bytesBuf.Reset()
	}
}
