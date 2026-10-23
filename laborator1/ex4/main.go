package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := []float32{1.0, 2.0, 3.0, 4.0}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	if index < 0 || index >= len(numbers) {
		fmt.Println("index in afara intervalului")
		return
	}

	var result []float32
	for i, num := range numbers {
		if i != index {
			result = append(result, num)
		}
	}

	fmt.Println(result)
}
