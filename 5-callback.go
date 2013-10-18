package main

import (
	"fmt"
	"runtime"
)

var counter int
var inc chan func(int) int
var done chan int

func main() {
	runtime.GOMAXPROCS(4)
	counter = 0
	inc = make(chan func(int) int)
	done = make(chan int)

	go incrementor()
	for i:=0; i<1000; i++ {
		inc <- func(count int) int {
			return count + 1
		}
	}

	<-done
	fmt.Println(counter)
}

func incrementor() {
	for i:=0; i<1000; i++ {
		callback := <-inc
		counter = callback(counter)
	}
	done<-1
}
