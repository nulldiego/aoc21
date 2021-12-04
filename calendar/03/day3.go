package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"aoc21/utils/files"
)

func main() {
	start := time.Now()
	input := files.ReadFile(3, "\r\n")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := calculatePowerConsumption(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = calculateLifeSupport(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func calculatePowerConsumption(input []string) (int, error) {
	half := len(input) / 2
	lineLength := len(input[0])
	zeros := make([]int, lineLength)
	for _, line := range input {
		for i, char := range line {
			if char == '0' {
				zeros[i]++
			}
		}
	}

	var gamma, epsilon int
	for i := 0; i < lineLength; i++ {
		gamma <<= 1
		epsilon <<= 1
		if half > zeros[i] {
			gamma |= 1
			continue
		}
		epsilon |= 1
	}

	return gamma * epsilon, nil
}

func calculateParameter(input []string, condition func(rune, rune) bool) (int64, error) {
	filtered := input
	index := 0
	for len(filtered) > 1 {
		var remaining []string
		mostCommonBit := '1' // if there is a tie, we consider '1' as the most common
		var ones, zeros int
		for _, line := range filtered {
			if rune(line[index]) == '0' {
				zeros++
				continue
			}
			ones++
		}

		if zeros > ones {
			mostCommonBit = '0'
		}

		for _, line := range filtered {
			if condition(rune(line[index]), mostCommonBit) {
				remaining = append(remaining, line)
			}
		}
		filtered = remaining
		index++
	}
	result, _ := strconv.ParseInt(filtered[0], 2, 32)
	return result, nil
}

func sameRune(a, b rune) bool {
	return a == b
}

func difRune(a, b rune) bool {
	return a != b
}

func calculateLifeSupport(input []string) (int, error) {
	oxy, _ := calculateParameter(input, sameRune)
	co2, _ := calculateParameter(input, difRune)
	return int(oxy * co2), nil
}
