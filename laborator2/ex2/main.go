package main

func receive(ch chan string) {
	var i = <-ch
	println(i)
}

func main() {
	var ch = make(chan string)
	go receive(ch)

	ch <- "salut"
}
