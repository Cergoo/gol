package refl

import (
	"gol/refl"
	"gol/test"
	"testing"
)

var t1 *test.TT

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

func Test_IsNil(t *testing.T) {
	t1 = test.New(t)

	type (
		T struct {
			f1 string
		}
	)

	var (
		m   map[string]string
		i   int
		obj interface{}
		tt  *T
	)
	// true
	t1.Eq(refl.IsNil(m), true)
	t1.Eq(refl.IsNil(obj), true)
	t1.Eq(refl.IsNil(tt), true)

	// false
	m = make(map[string]string)
	tt = new(T)
	obj = tt
	t1.Eq(refl.IsNil(i), false)
	t1.Eq(refl.IsNil(m), false)
	t1.Eq(refl.IsNil(tt), false)
	t1.Eq(refl.IsNil(obj), false)

	// from interface{}
	tt = nil
	obj = tt
	t1.Eq(obj == nil, false)
	t1.Eq(refl.IsNil(obj), true)
}

func Test_IsEmpty(t *testing.T) {
	t1 = test.New(t)

	var (
		v1 int
		v2 uint
		v3 float32
		v4 []string
	)

	t1.Eq(refl.IsEmpty(v1), true)
	t1.Eq(refl.IsEmpty(v2), true)
	t1.Eq(refl.IsEmpty(v3), true)
	t1.Eq(refl.IsEmpty(v4), true)

	v1 = 1
	v2 = 1
	v3 = -1
	v4 = make([]string, 1)

	t1.Eq(refl.IsEmpty(v1), false)
	t1.Eq(refl.IsEmpty(v2), false)
	t1.Eq(refl.IsEmpty(v3), false)
	t1.Eq(refl.IsEmpty(v4), false)

}
