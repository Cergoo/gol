package lookup

import (
	"github.com/Cergoo/gol/test"
	"testing"
)

type (
	t1 struct {
		Size     uint64
		Result   string
		SliceVar []string
	}
	t2 struct {
		Val1 string
		Val2 []*t1
	}
)

var V map[string]*t2

func init() {
	vt2 := &t2{
		Val1: "стр",
		Val2: make([]*t1, 10),
	}
	vt1 := &t1{
		Result:   "rezult",
		Size:     78,
		SliceVar: []string{"ctr1", "ctr2"},
	}

	vt2.Val2[2] = vt1
	V = make(map[string]*t2)
	V["point"] = vt2
}

func Test1(t *testing.T) {
	var fullpath bool
	t1 := test.New(t)

	v, fullpath := LookupI(V, []byte("point/Val1"))
	t1.Eq(v.(string), "стр")
	t1.Eq(fullpath, true)

	v, fullpath = LookupI(V, []byte("point/Val2/2/SliceVar/1"))
	t1.Eq(v.(string), "ctr2")
	t1.Eq(fullpath, true)

	v, fullpath = LookupI(V, []byte("point/Val1/ogoPath"))
	t1.Eq(v.(string), "стр")
	t1.Eq(fullpath, false)
}
