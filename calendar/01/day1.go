package main

import (
	"fmt"
	"log"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := mapper.ToIntSlice(inputSliceAsString)

	// Part 1
	start := time.Now()
	solution, err := countIncreasesByWindow(input, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countIncreasesByWindow(input, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n", time.Since(start))
}

func countIncreasesByWindow(input []int, windowSize int) (int, error) {
	var prevWindow, nextWindow, count int
	for _, depth := range input[:windowSize] {
		prevWindow += depth
	}
	for i, depth := range input[windowSize:] {
		nextWindow = prevWindow - input[i] + depth
		if nextWindow > prevWindow {
			count++
		}
		prevWindow = nextWindow
	}

	return count, nil
}
