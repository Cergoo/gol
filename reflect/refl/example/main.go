// Example use pkg
package main

import (
	"fmt"
	"github.com/Cergoo/gol/reflect/refl"
)

func main() {
	// Example mapKeysEq
	mapKeysEq()
	// Example structToMap
	structToMap()
}

// MapKeysEq
func mapKeysEq() {
	fmt.Println("MapKeysEq example:")
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	map1["n1"] = 10
	map1["n2"] = 12
	map2["n1"] = 20
	map2["n2"] = 22

	fmt.Println(refl.MapKeysEq(map1, map2))
}

// StructToMap
func structToMap() {
	fmt.Println("StructToMap example:")
	type (
		tobj2 struct {
			FA string
		}
		tobj struct {
			F1 int
			F2 string
			FA *tobj2
		}
	)

	obj := new(tobj)
	obj.F1 = 2
	obj.F2 = "text1"
	obj.FA = new(tobj2)
	obj.FA.FA = "nn"

	m := make(map[string]interface{})
	m["n"] = 5
	refl.StructToMap(obj, m, "obj.")
	fmt.Println(m)
}
