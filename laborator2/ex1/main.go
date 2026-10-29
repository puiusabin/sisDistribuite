package main

func send(ch chan string) {
	ch <- "salut"
}

func main() {
	var ch = make(chan string)
	go send(ch)
	var i = <-ch
	println(i)
}
