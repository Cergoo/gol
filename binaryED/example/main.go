package main

import (
	"fmt"
	"github.com/Cergoo/gol/binaryED"
	"github.com/Cergoo/gol/fastbuf"
)

func main() {
	buf := fastbuf.New(nil)

	binaryED.PutInt32(buf, 2)
	binaryED.PutInt32(buf, 4)
	fmt.Print(buf.FlushP())
}
