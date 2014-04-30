package main

import (
	"fmt"
	"gol/refl"
	"strconv"
)

type type1 int

func (t type1) f1(b int) int {
	return int(t) + b
}

func main() {

	// MapKeysEq
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	map1["n1"] = 10
	map1["n2"] = 12
	map2["n1"] = 20
	map2["n2"] = 22

	fmt.Println(refl.MapKeysEq(map1, map2))

	// Caller
	// example1
	caller := make(refl.FuncMap)
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
