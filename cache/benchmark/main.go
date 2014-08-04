package main

import (
	"fmt"
	gocache "github.com/pmylund/go-cache"
	"github.com/Cergoo/gol/cache"
	"github.com/Cergoo/gol/hash"
	"strconv"
	"sync"
	"testing"
	"time"
)

var (
	_cache   cache.Cache
	go_cache *gocache.Cache
	count    int
	m        sync.RWMutex
)

func init() {
	count = 1000
	_cache = cache.New(hash.HashFAQ6, count, true, 10*time.Minute, nil)
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
			_cache.Set("item"+strconv.Itoa(i), i)
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
