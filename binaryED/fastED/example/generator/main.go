package main

import (
	"github.com/Cergoo/gol/binaryED/fastED"
	"github.com/Cergoo/gol/binaryED/fastED/example/exportedtypes"
)

func main() {
	g := fastED.New("../encoderdecoder/encoderdecoder.go", "github.com/Cergoo/gol/binaryED/fastED/example/exportedtypes")
	g.Encode(&exportedtypes.T1{})
	g.Decode(&exportedtypes.T1{})
	g.Close()
}
