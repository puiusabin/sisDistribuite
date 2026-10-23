package main

import "fmt"

func main() {
	var intArr [3]int = [3]int{3, 2, 1}

	for i := 0; i < 3-1; i++ {
		for j := 0; j < 3-i-1; j++ {
			if intArr[j] > intArr[j+1] {
				aux := intArr[j]
				intArr[j] = intArr[j+1]
				intArr[j+1] = aux
			}
		}
	}
	fmt.Println(intArr)
}
