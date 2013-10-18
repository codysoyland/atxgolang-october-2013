package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int
var waitgroup sync.WaitGroup
var mutex sync.Mutex

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
	mutex.Lock()
	counter = counter + 1
	mutex.Unlock()
	waitgroup.Done()
}
