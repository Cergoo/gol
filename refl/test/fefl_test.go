package refl

import (
	"gol/refl"
	"testing"
)

func Test_MapKeysEq_ok1(t *testing.T) {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	map1["n1"] = 10
	map1["n2"] = 12
	map2["n1"] = 20
	map2["n2"] = 22

	v := refl.MapKeysEq(map1, map2)
	if !v {
		t.Error("err")
	}
}

func Test_MapKeysEq_ok2(t *testing.T) {
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	v := refl.MapKeysEq(map1, map2)
	if !v {
		t.Error("err")
	}
}

func Test_MapKeysEq_ok3(t *testing.T) {
	map1 := make(map[string]int)
	map2 := make(map[string]string)
	map1["n1"] = 10
	map1["n2"] = 12
	map2["n2"] = ""
	map2["n1"] = ""

	v := refl.MapKeysEq(map1, map2)
	if !v {
		t.Error("err")
	}
}

func Test_MapKeysEq_notok1(t *testing.T) {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	map1["n1"] = 10
	map2["n1"] = 12
	map2["n2"] = 20

	v := refl.MapKeysEq(map1, map2)
	if v {
		t.Error("err")
	}
}

func Test_MapKeysEq_notok2(t *testing.T) {
	map1 := make(map[string]int)
	map2 := make(map[string]string)
	map1["n1"] = 10
	map1["n2"] = 12
	map2["n1"] = ""
	map2["n4"] = ""

	v := refl.MapKeysEq(map1, map2)
	if v {
		t.Error("err")
	}
}
