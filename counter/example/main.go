package main

import (
	"fmt"
	"github.com/Cergoo/gol/counter"
)

func main() {
	n := uint8(12)
	n = n - 20
	fmt.Println(n)

	c := new(counter.T_counter)
	c.Set(20)
	fmt.Println(c.Add(-22))
}
