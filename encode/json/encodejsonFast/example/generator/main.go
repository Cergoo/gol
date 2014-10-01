package main

import (
	"github.com/Cergoo/gol/encode/json/encodejsonFast"
	"github.com/Cergoo/gol/encode/json/encodejsonFast/example/exportedtypes"
)

func main() {
	g := encodejsonFast.New("../encoderdecoder/encoderdecoder.go", "github.com/Cergoo/gol/encode/json/encodejsonFast/example/exportedtypes")
	g.Encode(&exportedtypes.T1{})
	g.Close()
}
