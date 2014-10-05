// Benchmark
package main

import (
	"fmt"
	"github.com/Cergoo/gol/encode/binary/primitive"
	"github.com/Cergoo/gol/fastbuf"
	"testing"
)

var (
	buf1 = fastbuf.New(nil, 0, nil)
	buf2 = make([]byte, 0, 64)
)

func main() {
	fmt.Println(primitive.EndianBig())
	operation_name := "Write"
	r1 := testing.Benchmark(benchmark_buf1)
	r2 := testing.Benchmark(benchmark_buf2)
	fmt.Print(operation_name, "\n", "buf1:", r1, r1.MemString(), "\n", "buf2:", r2, r2.MemString(), "\n")

}

func benchmark_buf1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 1; i < 2000; i++ {
			primitive.PutUint16(buf1, uint16(257))
		}
		buf1.ReadWriteReset()
	}
}

func benchmark_buf2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 1; i < 2000; i++ {
			buf2 = primitive.PutUint16v1(buf2, uint16(257))
		}
		buf2 = buf2[:0]
	}
}
