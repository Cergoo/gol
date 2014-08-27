// Example use pkg
package main

import (
	"fmt"
	"github.com/Cergoo/gol/reflect/caller"
	"strconv"
)

type (
	type1 int
)

func (t type1) f1(b int) int {
	return int(t) + b
}

func main() {

	// Example caller
	calle()

}

// Caller
func calle() {
	fmt.Println("Caller example:")

	// example1
	caller := make(caller.FuncMap)
	caller.Add("itoa", strconv.Itoa)
	i := caller.Calli("itoa", 10)[0].(string)
	fmt.Println(i)

	// example2
	var v type1
	v = 10
	caller.Add("f1", v.f1)
	i1 := caller.Calli("f1", 10)[0].(int)
	fmt.Println(i1)
}
