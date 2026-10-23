package main

import "fmt"

func main() {
	heights := map[string]int{
		"Ana":    165,
		"Mihai":  182,
		"Ioana":  170,
		"Andrei": 182,
		"Elena":  158,
	}

	var names []string
	for name := range heights {
		names = append(names, name)
	}

	n := len(names)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if heights[names[j]] < heights[names[j+1]] {
				aux := names[j]
				names[j] = names[j+1]
				names[j+1] = aux
			}
		}
	}

	for _, name := range names {
		fmt.Println(name, heights[name])
	}
}
