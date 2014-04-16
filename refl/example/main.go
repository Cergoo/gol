package main

import (
	"fmt"
	"gol/refl"
	"strconv"
)

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
	caller := make(refl.FuncMap)
	caller.Add("itoa", strconv.Itoa)
	i := caller.Calli("itoa", 10)[0].(string)
	fmt.Println(i)

}
