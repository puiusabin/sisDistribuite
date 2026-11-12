package main

import (
	"fmt"
	"math"
	"os"
)

const radius = 16.0
const maxIterations = 1000
const eps = 0.000001

type Point struct {
	X, Y float64
}

func distance(a, b Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func compress(robots []Point) []Point {
	var result []Point
	for _, r := range robots {
		exists := false
		for _, p := range result {
			if distance(r, p) < eps {
				exists = true
			}
		}
		if !exists {
			result = append(result, r)
		}
	}
	return result
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var n int
	fmt.Fscan(f, &n)
	robots := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(f, &robots[i].X, &robots[i].Y)
	}

	for iter := 1; iter <= maxIterations && len(robots) > 1; iter++ {
		newPositions := make([]Point, len(robots))

		for i := range robots {
			sum := robots[i]
			count := 1
			for j := range robots {
				if i != j && distance(robots[i], robots[j]) < radius {
					sum.X += robots[j].X
					sum.Y += robots[j].Y
					count++
				}
			}

			newPositions[i] = Point{sum.X / float64(count), sum.Y / float64(count)}
		}

		robots = compress(newPositions)
		fmt.Println(iter, ":", robots)
	}

	fmt.Println("final:", robots)
}
