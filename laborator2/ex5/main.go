package main

import "fmt"

func send(ch chan string) {
	ch <- "salut"
}

func receive(ch chan string, result chan string) {
	var msg = <-ch
	result <- msg
}

func main() {
	var ch = make(chan string)
	var result = make(chan string)
	go send(ch)
	go receive(ch, result)
	var msg = <-result
	fmt.Println(msg)
}
