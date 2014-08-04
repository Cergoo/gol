package main

import (
	"fmt"
	"github.com/Cergoo/gol/jsonConfig"
)

func main() {
	m := make(map[string]string)
	jsonConfig.Load("conf.json", &m)
	fmt.Println(m)
}
