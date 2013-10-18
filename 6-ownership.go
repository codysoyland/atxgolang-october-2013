package main

import (
	"fmt"
	"runtime"
)

var grant chan int
var reclaim chan int
var done chan int

func main() {
	runtime.GOMAXPROCS(4)
	grant = make(chan int)
	reclaim = make(chan int)
	done = make(chan int)

	go incrementor()

	grant <- 0

	for {
		select {
		case counter := <-reclaim:
			grant <- counter
		case counter := <-done:
			fmt.Println(counter)
			return
		}
	}
}

func incrementor() {
	for i:=0; i<1000; i++ {
		counter := <-grant
		counter = counter + 1
		reclaim <- counter
	}
	done <- <-grant
}
