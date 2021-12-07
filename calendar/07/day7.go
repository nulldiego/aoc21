package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

func main() {
	start := time.Now()
	input := mapper.ToIntSlice(files.ReadFile(7, ","))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := minFuelToAlign(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = minFuelToAlign2(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minFuelToAlign(input []int) (int, error) {
	var minFuel int
	sort.Ints(input)
	for position := input[0]; position <= input[len(input)-1]; position++ {
		fuel := 0
		for _, crab := range input {
			fuel += abs(position - crab)
		}
		if position == input[0] || fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel, nil
}

func minFuelToAlign2(input []int) (int, error) {
	var minFuel int
	sort.Ints(input)
	for position := input[0]; position <= input[len(input)-1]; position++ {
		fuel := 0
		for _, crab := range input {
			fuel += abs(position-crab) * (abs(position-crab) + 1) / 2
		}
		if position == input[0] || fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel, nil
}
