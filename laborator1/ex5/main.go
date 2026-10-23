package main

import "fmt"

func main() {
	words := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	n := len(words)
	for i := 0; i < n/2; i++ {
		aux := words[i]
		words[i] = words[n-1-i]
		words[n-1-i] = aux
	}

	fmt.Println(words)
}
