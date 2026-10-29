package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robot(ch chan string, name string) {
	var distance = 0
	for distance < 100 {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		distance = distance + 1
	}
	ch <- name
}

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	go robot(ch1, "robot 1")
	go robot(ch2, "robot 2")

	select {
	case name := <-ch1:
		fmt.Println(name)
	case name := <-ch2:
		fmt.Println(name)
	}
}
