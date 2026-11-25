package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Message struct {
	SenderID int
	Round    int
	value    int
}

type Node struct {
	id          int
	x           int
	s           int
	n           int
	isByzantine bool
	inbox       chan Message
}

func Broadcast(nodes []Node, from int, round int, value int) {
	for i := range nodes {
		nodes[i].inbox <- Message{SenderID: from, Round: round, value: value}
	}
}

func ByzantineBroadcast(nodes []Node, from int, round int) {
	for i := range nodes {
		nodes[i].inbox <- Message{SenderID: from, Round: round, value: rand.Intn(2)}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := strings.Fields(scanner.Text())

	n, err := strconv.Atoi(line[0])
	if err != nil {
		panic(err)
	}

	s, err := strconv.Atoi(line[1])
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	line = strings.Fields(scanner.Text())
	stari := []int{}
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(line[i])
		if err != nil {
			panic(err)
		}
		stari = append(stari, num)
	}

	if n <= 4*s {
		fmt.Println("need n > 4s")
		return
	}

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			id:          i,
			x:           stari[i],
			s:           s,
			n:           n,
			isByzantine: false,
			inbox:       make(chan Message, n),
		}
	}

	perm := rand.Perm(n)
	for i := 0; i < s; i++ {
		nodes[perm[i]].isByzantine = true
	}

	for t := 0; t <= s+1; t++ {
		for i := range nodes {
			if nodes[i].isByzantine {
				ByzantineBroadcast(nodes, i, 1)
			} else {
				Broadcast(nodes, i, 1, nodes[i].x)
			}
		}

		majority := make([]int, n)
		mult := make([]int, n)
		for i := range nodes {
			ones := 0
			for k := 0; k < n; k++ {
				msg := <-nodes[i].inbox
				ones += msg.value
			}
			if ones >= n/2+1 {
				majority[i] = 1
				mult[i] = ones
			} else {
				majority[i] = 0
				mult[i] = n - ones
			}
		}

		if nodes[t].isByzantine {
			ByzantineBroadcast(nodes, t, 2)
		} else {
			Broadcast(nodes, t, 2, majority[t])
		}

		for i := range nodes {
			tie := (<-nodes[i].inbox).value
			if nodes[i].isByzantine {
				continue
			}
			if 2*mult[i] > n+2*s {
				nodes[i].x = majority[i]
			} else {
				nodes[i].x = tie
			}
		}
	}

	for i := range nodes {
		if !nodes[i].isByzantine {
			fmt.Println(nodes[i].x)
			return
		}
	}
}
