// Example use pkg
package main

import (
	"fmt"
	"github.com/Cergoo/gol/cache/cacheUint"
	"time"
)

func f(v *interface{}) {
	n := (*v).(int)
	n *= 5
	*v = n
}

func main() {
	var (
		t int
	)
	n := cacheUint.New(true, 0*time.Minute, nil)
	for t = 0; t < 200000; t++ {
		n.Set(uint64(t), t, 1, cacheUint.UpdateOrInsert)
	}
	fmt.Println(n.GetBucketsStat())
	fmt.Println(n.Get(1))
	fmt.Println(n.Func(1, f))
	fmt.Println(n.Get(1))

	n = nil

	//runtime.GC()
	//runtime.Gosched()
	//time.Sleep(10 * time.Second)
}
