package main

import (
	"fmt"
	"github.com/Cergoo/gol/cache"
	"github.com/Cergoo/gol/hash"
	"strconv"
	"time"
)

func main() {
	var (
		t int
		l interface{}
		//i int
	)
	n := cache.New(hash.HashFAQ6, 0, true, 0*time.Minute, nil)
	for t = 0; t < 100000; t++ {
		n.Set(&cache.TCortege{"ind uhgyug e x try" + strconv.Itoa(t), t}, cache.ModeSet_UpdateOrInsert)
	}
	fmt.Println(n.Get("ind uhgyug e x try1"))
	time.Sleep(10 * time.Second)
	for t = 0; t < 100000; t++ {
		l = n.Get("ind uhgyug e x try" + strconv.Itoa(t))
		if l == nil {
			fmt.Println(l, t, "nn")
		}
		fmt.Println(l)
		//i = l.(int)
		//fmt.Println(i)

	}

	fmt.Print(n.GetBucketsStat())
	//ind uhgyug e x try
	//fmt.Print(n.Get("ind uhgyug e x try5814"))

}
