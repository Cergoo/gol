// Benchmark
package main

import (
	"fmt"
	"github.com/Cergoo/gol/cache/cacheStr"
	"github.com/Cergoo/gol/hash"
	gocache "github.com/pmylund/go-cache"
	"strconv"
	"sync"
	"testing"
	"time"
)

const count = 1000

var (
	_cache   cacheStr.Cache
	go_cache *gocache.Cache
	m        sync.RWMutex
)

func init() {
	_cache = cacheStr.New(hash.HashFAQ6, true, 10*time.Minute, nil)
	go_cache = gocache.New(5*time.Minute, 10*time.Minute)
}

func main() {
	var (
		rcache, rgo_cache testing.BenchmarkResult
		operation_name    string
	)

	operation_name = "Set"
	rcache = testing.Benchmark(Benchmark_cacheSet)
	rgo_cache = testing.Benchmark(Benchmark_gocacheSet)
	fmt.Print(operation_name, "\n", "Cergoo.cache:", rcache, rcache.MemString(), "\n", "go-cache:", rgo_cache, rgo_cache.MemString(), "\n")
	operation_name = "Get"
	rcache = testing.Benchmark(Benchmark_cacheGet)
	rgo_cache = testing.Benchmark(Benchmark_gocacheGet)
	fmt.Print(operation_name, "\n", "Cergoo.cache:", rcache, rcache.MemString(), "\n", "go-cache:", rgo_cache, rgo_cache.MemString(), "\n")
	operation_name = "Inc"
	rcache = testing.Benchmark(Benchmark_cacheInc)
	rgo_cache = testing.Benchmark(Benchmark_gocacheInc)
	fmt.Print(operation_name, "\n", "Cergoo.cache:", rcache, rcache.MemString(), "\n", "go-cache:", rgo_cache, rgo_cache.MemString(), "\n")
}

func Benchmark_cacheSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			_cache.Set("item"+strconv.Itoa(i), i, 1, cacheStr.UpdateOrInsert)
		}
	}
}

func Benchmark_gocacheSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go_cache.Set("item"+strconv.Itoa(i), i, 0)
		}
	}
}

func Benchmark_cacheGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			_cache.Get("item" + strconv.Itoa(i))
		}
	}
}

func Benchmark_gocacheGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go_cache.Get("item" + strconv.Itoa(i))
		}
	}
}

func Benchmark_cacheInc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			_cache.Inc("item"+strconv.Itoa(i), 25)
		}
	}
}

func Benchmark_gocacheInc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			go_cache.Increment("item"+strconv.Itoa(i), 25)
		}
	}
}
