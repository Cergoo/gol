package main

import (
	"fmt"
	"github.com/Cergoo/gol/binaryED/fbinaryED"
	"github.com/Cergoo/gol/fastbuf"
)

func main() {
	demoED()
}

func demoED() {
	buf := fastbuf.New(nil)

	var (
		i int
		s string
	)

	i = 4
	fbinaryED.Encode(buf, i)
	i = 2
	fbinaryED.Decode(buf, &i)
	s = "nnnnnn"
	fbinaryED.Encode(buf, s)
	s = ""
	fbinaryED.Decode(buf, &s)

	vslice := []uint32{4: 12}
	var outslice []uint32
	fbinaryED.Encode(buf, vslice)
	fbinaryED.Decode(buf, &outslice)
	fmt.Println(i, s, outslice)

	// struct ED
	type (
		t1 struct {
			V1 int
			V2 string
		}
	)

	vobj1 := &t1{
		V1: -4,
		V2: "test1",
	}
	vobj2 := new(t1)
	fbinaryED.Encode(buf, vobj1)
	fbinaryED.Decode(buf, vobj2)
	fmt.Println(vobj2)
}
