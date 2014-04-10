package main

import (
	"fmt"
	"gol/jsonConfig"
)

func main() {
	//m := make(map[string]string)
	type (
		t1
	)
	m := make([][2]string, 10)
	jsonConfig.Load("conf.json", &m)
	fmt.Println(m)
}
