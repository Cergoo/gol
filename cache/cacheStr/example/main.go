// Example use pkg
package main

import (
	"fmt"
	"github.com/Cergoo/gol/cache/cacheStr"
	"github.com/Cergoo/gol/hash"
	"strconv"
	"time"
)

func main() {
	var (
		t int
	)
	n := cacheStr.New(hash.HashJenkins, true, 0*time.Minute, nil)
	for t = 0; t < 200000; t++ {
		n.Set("ind uhgyug e x try"+strconv.Itoa(t), t, 1, cacheStr.UpdateOrInsert)
	}
	n.Set("ind", 101, 1, cacheStr.UpdateOrInsert)
	n.Inc("ind", -1)
	fmt.Println("Decrement: ", n.Get("ind"))
	//time.Sleep(10 * time.Second)
	fmt.Print(n.GetBucketsStat())
	n = nil
	//runtime.GC()
	//runtime.Gosched()
	//time.Sleep(10 * time.Second)
}
