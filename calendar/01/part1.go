package main

import (
	"fmt"
	"log"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := mapper.ToIntSlice(inputSliceAsString)

	solution, err := countIncreases(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
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
