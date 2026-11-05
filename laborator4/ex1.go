package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool
}

func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	out <- Msg{Value: id}

	for {
		msg := <-in

		if msg.Done {
			if !*isLeader {
				out <- msg
			}
			return
		}

		if msg.Value > id {
			out <- msg
		} else if msg.Value == id {
			*isLeader = true
			out <- Msg{Value: id, Done: true}
		}
	}
}

func RingFlooding(ids []int) int {
	n := len(ids)

	chans := make([]chan Msg, n)
	for i := range chans {
		chans[i] = make(chan Msg, n+1)
	}

	leaders := make([]bool, n)
	var wg sync.WaitGroup

	for i := range ids {
		wg.Add(1)
		go node(ids[i], chans[i], chans[(i+1)%n], &leaders[i], &wg)
	}
	wg.Wait()

	for i := range ids {
		if leaders[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(RingFlooding([]int{3, 7, 2, 9, 5}))
}
