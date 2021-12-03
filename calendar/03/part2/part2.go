package main

import (
	"aoc21/utils/files"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")

	solution, err := findPosition(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}

func findPosition(input []string) (int, error) {
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
