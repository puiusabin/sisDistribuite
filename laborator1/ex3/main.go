package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	var evenSlice []int

	for _, num := range nums {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}
	fmt.Println(evenSlice)
}
