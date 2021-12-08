package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"aoc21/utils/files"
)

func main() {
	start := time.Now()
	input := files.ReadFile(8, "\r\n")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countUniques(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = sumOutputs(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func countUniques(input []string) (int, error) {
	var count int
	for _, line := range input {
		parts := strings.Split(line, " | ")
		for _, digit := range strings.Fields(parts[1]) {
			if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
				count++
			}
		}
	}
	return count, nil
}

func fillUniqueMappings(digits []string) map[int]string {
	uniques := make(map[int]string, 4)
	for _, digit := range digits {
		switch len(digit) {
		case 2:
			uniques[1] = digit
		case 4:
			uniques[4] = digit
		case 3:
			uniques[7] = digit
		case 7:
			uniques[8] = digit
		}
	}
	return uniques
}

func countIntersections(a, b string) int {
	var count int
	for _, rune := range b {
		if strings.ContainsRune(a, rune) {
			count++
		}
	}
	return count
}

func sumOutputs(input []string) (int, error) {
	var sum int
	for _, line := range input {
		parts := strings.Split(line, " | ")
		uniques := fillUniqueMappings(strings.Fields(parts[0]))
		var output int
		for _, digit := range strings.Fields(parts[1]) {
			output = output * 10
			switch len(digit) {
			case 2:
				output += 1
			case 4:
				output += 4
			case 3:
				output += 7
			case 7:
				output += 8
			case 5:
				if countIntersections(digit, uniques[1]) == 2 {
					output += 3
				} else if countIntersections(digit, uniques[4]) == 3 {
					output += 5
				} else {
					output += 2
				}
			case 6:
				if countIntersections(digit, uniques[1]) == 1 {
					output += 6
				} else if countIntersections(digit, uniques[4]) == 4 {
					output += 9
				}
			}
		}
		sum += output
	}
	return sum, nil
}
