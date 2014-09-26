package main

import (
	"fmt"
	"github.com/Cergoo/gol/binaryED/primitive"
	"github.com/Cergoo/gol/fastbuf"
)

func main() {
	buf := fastbuf.New(nil, 0, nil)
	primitive.PutInt32(buf, 2)
	primitive.PutInt32(buf, 4)
	fmt.Println(buf.FlushP())
}
