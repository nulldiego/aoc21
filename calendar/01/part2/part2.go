package main

import (
	"aoc21/utils/files"
	"aoc21/utils/mapper"
	"fmt"
	"log"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := mapper.ToIntSlice(inputSliceAsString)

	solution, err := countIncreasesByWindow(input, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
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
