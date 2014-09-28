package main

import (
	"github.com/Cergoo/gol/encode/binary/encodebinaryFast"
	"github.com/Cergoo/gol/encode/binary/encodebinaryFast/example/exportedtypes"
)

func main() {
	g := encodebinaryFast.New("../encoderdecoder/encoderdecoder.go", "github.com/Cergoo/gol/encode/binary/encodebinaryFast/example/exportedtypes")
	g.Encode(&exportedtypes.T1{})
	g.Decode(&exportedtypes.T1{})
	g.Close()
}
