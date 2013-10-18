package main

import (
	"fmt"
	"runtime"
)

var counter int
var inc chan int
var done chan int

func main() {
	runtime.GOMAXPROCS(4)
	counter = 0
	inc = make(chan int)
	done = make(chan int)

	go incrementor()
	for i:=0; i<1000; i++ {
		go increment()
	}

	<-done
	fmt.Println(counter)
}

func increment() {
	inc <- 1
}

func incrementor() {
	for i:=0; i<1000; i++ {
		<-inc
		counter = counter + 1
	}
	done<-1
}
