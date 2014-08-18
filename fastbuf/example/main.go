package main

import (
	"fmt"
	"github.com/Cergoo/gol/fastbuf"
)

func main() {
	b := fastbuf.New(nil)
	b.Write([]byte("12"))
	fmt.Print(b.FlushP())

}
