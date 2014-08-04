package main

import (
	"fmt"
	"github.com/Cergoo/gol/genid"
)

func main() {
	gi := genid.NewHTTPGen(184)
	for i := 0; i < 10; i++ {
		id := gi.NewID()
		fmt.Println(id, len(id))
	}
}
