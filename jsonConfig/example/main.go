package main

import (
	"github.com/Cergoo/gol/jsonConfig"
	"github.com/davecgh/go-spew/spew"
)

type (
	v1 struct {
		Portudp string
		In      struct {
			F1 int
			F2 string
		}
		Readtimeout      string
		Maxheader        int
		Initbucketscount int
		Duration         string
		Udp              string
		Porttcp          string
		Workerscount     int
	}
)

func main() {
	m := new(v1)
	jsonConfig.Load("conf.json", &m)
	spew.Dump(m)
}
