package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int
var waitgroup sync.WaitGroup
var token chan int

func main() {
	runtime.GOMAXPROCS(4)
	counter = 0
	waitgroup = sync.WaitGroup{}
	token = make(chan int)

	waitgroup.Add(1000)
	for i:=0; i<1000; i++ {
		go increment()
	}
	token <- 1
	go func() {<-token}()

	waitgroup.Wait()


	fmt.Println(counter)
}

func increment() {
	<-token
	counter = counter + 1
	token <- 1
	waitgroup.Done()
}
