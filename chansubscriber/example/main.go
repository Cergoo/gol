package main

import (
	"fmt"
	"gol/chansubscriber"
)

var (
	chwriter chan interface{}
)

func init() {
	chwriter = make(chan interface{}, 2)
}

func main() {
	hub := chansubscriber.New(chwriter, true)
	chreader1 := make(chan interface{}, 2)
	chreader2 := make(chan interface{}, 2)
	chreader3 := make(chan interface{}, 2)
	hub.Subscribe("1", chreader1)
	hub.Subscribe("2", chreader2)
	hub.Subscribe("3", chreader3)
	
	chwriter <- 10
	chwriter <- 11

	fmt.Println(<-chreader1, <-chreader1)
	fmt.Println(<-chreader2, <-chreader2)
	fmt.Println(<-chreader3, <-chreader3)
}
