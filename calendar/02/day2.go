package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"aoc21/utils/files"
)

func main() {
	input := files.ReadFile(2, "\n")

	// Part 1
	start := time.Now()
	solution, err := findPosition(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = findPositionWithAim(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n", time.Since(start))
}

func findPosition(input []string) (int, error) {
	var depth, horizontal int
	for _, move := range input {
		movement := strings.Fields(move)
		amount, _ := strconv.Atoi(movement[1])
		switch movement[0] {
		case "forward":
			horizontal += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}

	return depth * horizontal, nil
}

func findPositionWithAim(input []string) (int, error) {
	var depth, horizontal, aim int
	for _, move := range input {
		movement := strings.Fields(move)
		amount, _ := strconv.Atoi(movement[1])
		switch movement[0] {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}

	return depth * horizontal, nil
}
