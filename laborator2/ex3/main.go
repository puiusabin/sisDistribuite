package main

import "fmt"

func send(ch chan string, msg string) {
	ch <- msg
}

func main() {
	var ch = make(chan string)
	go send(ch, "salut 1")
	fmt.Println(<-ch)
	go send(ch, "salut 2")
	fmt.Println(<-ch)
}
