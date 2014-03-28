package main

import (
	"fmt"
	"gol/genid"
)

func main() {
	gi := genid.NewHTTPGen(32)
	for i := 0; i < 10; i++ {
		id := gi.NewID()
		fmt.Println(id, len(id))
	}
}
