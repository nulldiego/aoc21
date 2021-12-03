package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"aoc21/utils/files"
)

func main() {
	input := files.ReadFile(3, "\n")

	// Part 1
	start := time.Now()
	solution, err := calculatePowerConsumption(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = calculateLifeSupport(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n", time.Since(start))
}

func calculatePowerConsumption(input []string) (int, error) {
	lineLength := len(input[0]) - 1
	zeros := make([]int, lineLength)
	ones := make([]int, lineLength)
	for _, line := range input {
		for i, char := range line[:lineLength] {
			if char == '0' {
				zeros[i]++
				continue
			}
			ones[i]++
		}
	}

	var gamma, epsilon int
	for i := 0; i < lineLength; i++ {
		gamma <<= 1
		epsilon <<= 1
		if ones[i] > zeros[i] {
			gamma |= 1
			continue
		}
		epsilon |= 1
	}

	return gamma * epsilon, nil
}

func calculateLifeSupport(input []string) (int, error) {
	lineLength := len(input[0]) - 1

	filteredOxygen := make([]string, len(input))
	filteredCo2 := make([]string, len(input))
	copy(filteredOxygen, input)
	copy(filteredCo2, input)

	index := 0
	for len(filteredOxygen) > 1 {
		var remaining []string
		mostCommonBit := '1'
		var ones, zeros int
		for _, line := range filteredOxygen {
			if rune(line[index]) == '0' {
				zeros++
				continue
			}
			ones++
		}

		if zeros > ones {
			mostCommonBit = '0'
		}

		for _, line := range filteredOxygen {
			if rune(line[index]) == mostCommonBit {
				remaining = append(remaining, line)
			}
		}
		filteredOxygen = remaining
		index++
	}

	index = 0
	for len(filteredCo2) > 1 {
		var remaining []string
		mostCommonBit := '1'
		var ones, zeros int
		for _, line := range filteredCo2 {
			if rune(line[index]) == '0' {
				zeros++
				continue
			}
			ones++
		}

		if zeros > ones {
			mostCommonBit = '0'
		}

		for _, line := range filteredCo2 {
			if rune(line[index]) != mostCommonBit {
				remaining = append(remaining, line)
			}
		}
		filteredCo2 = remaining
		index++
	}

	oxy, _ := strconv.ParseInt(filteredOxygen[0][:lineLength], 2, 64)
	co2, _ := strconv.ParseInt(filteredCo2[0][:lineLength], 2, 64)
	return int(oxy * co2), nil
}
