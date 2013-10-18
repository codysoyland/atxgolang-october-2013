package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int
var waitgroup sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(4)
	counter = 0
	waitgroup = sync.WaitGroup{}

	waitgroup.Add(1000)

	for i:=0; i<1000; i++ {
		go increment()
	}

	waitgroup.Wait()

	fmt.Println(counter)
}

func increment() {
	counter = counter + 1
	waitgroup.Done()
}
