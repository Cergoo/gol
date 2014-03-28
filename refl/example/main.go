package main

import (
	"fmt"
	"gol/refl"
)

func main() {

	map1 := make(map[string]int)
	map2 := make(map[string]int)

	map1["n1"] = 10
	map1["n2"] = 12
	map2["n1"] = 20
	map2["n2"] = 22

	fmt.Println(refl.MapKeysEq(map1, map2))

}
