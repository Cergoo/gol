// Example use pkg
package main

import (
	"fmt"
	"github.com/Cergoo/gol/reflect/refl"
	"reflect"
)

func main() {
	// Example mapKeysEq
	//mapKeysEq()
	// Example structToMap
	//structToMap()

	n2()
	//mapResearch()
	//structResearch()
	//sliceResearch()
}

func n1(n interface{}) {
	vt := reflect.ValueOf(n)

	fmt.Println(vt.Field(0).Kind(), vt.Field(0).Elem().Type().String())
}

func n2() {
	type (
		nn struct {
			b interface{}
		}
	)
	var n nn
	n = nn{b: 7}
	n1(n)
}

// Get slice elem type
func sliceResearch() {
	var (
		slice []byte
	)

	//slice = make([]byte, 10)
	//slice = nil
	vt := reflect.TypeOf(slice)
	vv := reflect.ValueOf(slice)
	fmt.Println(vt.Elem().Kind(), vv.Kind())
}

func structResearch() {
	type (
		t1 struct {
			v1 int
			v2 string
		}
	)
	var v1 t1
	v := new(t1)
	//v = nil
	fmt.Println(reflect.ValueOf(v).Elem().Kind())

	//v1 := t1{}
	fmt.Println(reflect.ValueOf(v1).Kind())

	var v2 *t1
	v2 = nil
	rv := reflect.ValueOf(v2)
	fmt.Println(rv.Kind(), rv.Elem().Kind())
}

// Get map type key and type val
func mapResearch() {
	map1 := make(map[string]int)
	vt := reflect.TypeOf(map1)
	n := vt.Elem().String()
	fmt.Println(vt.Key().String(), vt.Kind().String(), n)
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
