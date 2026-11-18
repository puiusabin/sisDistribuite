package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
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
	tokens := []int{}
	for i := 0; i < n; i++ {
		num, err := strconv.Atoi(line[i])
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, num)
	}

	adj := [][]int{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = strings.Fields(scanner.Text())
		neighbours := []int{}
		for j := 0; j < len(line); j++ {
			num, err := strconv.Atoi(line[j])
			if err != nil {
				panic(err)
			}
			neighbours = append(neighbours, num)
		}
		adj = append(adj, neighbours)
	}

	diam := diametru(adj)

	nodes := []Node{}
	for i := 0; i < n; i++ {
		m := Memory{}
		m[Info{ID: i, Value: tokens[i]}] = struct{}{}
		nodes = append(nodes, Node{ID: i, Neighbors: adj[i], M: m})
	}

	crashed := make([]bool, n)
	crashes := 0

	rounds := (s + 1) * diam

	for r := 1; r <= rounds; r++ {
		if crashes < s && rand.Intn(2) == 1 {
			x := rand.Intn(n)
			if !crashed[x] {
				crashed[x] = true
				crashes++
			}
		}

		newM := []Memory{}
		for i := 0; i < n; i++ {
			m := Memory{}
			for info := range nodes[i].M {
				m[info] = struct{}{}
			}
			newM = append(newM, m)
		}

		for i := 0; i < n; i++ {
			if crashed[i] {
				continue
			}
			for _, v := range nodes[i].Neighbors {
				for info := range nodes[i].M {
					newM[v][info] = struct{}{}
				}
			}
		}

		for i := 0; i < n; i++ {
			nodes[i].M = newM[i]
		}
	}

	correct := 0
	for i := 0; i < n; i++ {
		if !crashed[i] {
			correct = i
			break
		}
	}

	fmt.Println(diam)
	fmt.Println(rounds)

	for id := 0; id < n; id++ {
		for info := range nodes[correct].M {
			if info.ID == id {
				fmt.Print("(", info.ID, ", ", info.Value, ") ")
			}
		}
	}
	fmt.Println()
}
