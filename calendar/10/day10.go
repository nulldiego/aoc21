package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"aoc21/utils/files"
)

func main() {
	start := time.Now()
	input := files.ReadFile(10, "\r\n")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := sumSyntaxErrors(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = multiplyAutocompletes(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func sumSyntaxErrors(input []string) (int, error) {
	openers := "([{<"
	var opened []rune
	var sum int
	scores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	matching := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}

	for _, line := range input {
		opened = []rune{}
		for _, run := range line {
			if strings.ContainsRune(openers, run) {
				opened = append(opened, run)
				continue
			}
			if opened[len(opened)-1] != matching[run] {
				sum += scores[run]
				break
			} else {
				opened = opened[:len(opened)-1]
			}
		}
	}
	return sum, nil
}

func multiplyAutocompletes(input []string) (int, error) {
	openers := "([{<"
	var opened []rune
	scores := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	matching := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	var scoredLines []int
	var error bool

	for _, line := range input {
		opened = []rune{}
		error = false
		for _, run := range line {
			if strings.ContainsRune(openers, run) {
				opened = append(opened, run)
				continue
			}
			if opened[len(opened)-1] != matching[run] {
				error = true
				break
			} else {
				opened = opened[:len(opened)-1]
			}
		}
		if !error {
			var sum int
			for i, _ := range opened {
				sum = sum * 5
				sum += scores[opened[len(opened)-i-1]]
			}
			scoredLines = append(scoredLines, sum)
		}
	}
	sort.Ints(scoredLines)
	return scoredLines[len(scoredLines)/2], nil
}
