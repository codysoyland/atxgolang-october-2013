package main

import (
	"fmt"
	"runtime"
	"time"
)

type Command interface {
	Call(map[string]string)
}

type GetCommand struct {
	name string
	retchan chan string
}
func (c GetCommand) Call(m map[string]string) {
	c.retchan <- m[c.name]
}

type SetCommand struct {
	name string
	value string
}
func (c SetCommand) Call(m map[string]string) {
	m[c.name] = c.value
}

type SafeMap struct {
	commandchan chan Command
}
func (m *SafeMap) Get(name string) string {
	retchan := make(chan string)
	m.commandchan <- GetCommand{name, retchan}
	return <- retchan
}
func (m *SafeMap) Set(name, value string) {
	m.commandchan <- SetCommand{name, value}
}
func (m *SafeMap) Run() {
	kv := make(map[string]string)
	for {
		command := <-m.commandchan
		command.Call(kv)
	}
}
func NewSafeMap() *SafeMap {
	m := SafeMap{}
	m.commandchan = make(chan Command)
	return &m
}
func main() {
	runtime.GOMAXPROCS(4)
	m := NewSafeMap()
	go m.Run()
	go m.Set("test1", "a")
	go m.Set("test2", "b")

	time.Sleep(time.Second)
	fmt.Println(m.Get("test2"))
}
