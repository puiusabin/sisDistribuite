package main

import "fmt"

func notify(ch chan int, s int) {
	ch <- s
}

func sum(ch chan int, a int, b int) {
	var s = a + b
	defer notify(ch, s)
}

func main() {
	var ch = make(chan int)
	go sum(ch, 2, 3)
	var s = <-ch
	fmt.Println(s)
}
