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
	solution, err := countIncreases(input)
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

func countIncreases(input []int) (int, error) {
	var count int
	for i, depth := range input[1:] {
		if depth > input[i] {
			count++
		}
	}

	return count, nil
}

func countIncreasesByWindow(input []int, windowSize int) (int, error) {
	var prevWindow, nextWindow, count int
	for i, depth := range input {
		if i < windowSize {
			prevWindow += depth
			continue
		}
		nextWindow = depth + input[i-1] + input[i-2]
		if prevWindow < nextWindow {
			count++
		}
		prevWindow = nextWindow
	}

	return count, nil
}
