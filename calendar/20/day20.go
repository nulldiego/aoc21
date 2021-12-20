package main

import (
	"fmt"
	"time"

	"aoc21/utils/files"
)

func parseInput(input []string) ([]rune, [][]rune) {
	algorithm := []rune(input[0])
	inputImage := [][]rune{}
	for _, line := range input[2:] {
		inputImage = append(inputImage, []rune(line))
	}
	return algorithm, inputImage
}

func main() {
	start := time.Now()
	algorithm, input := parseInput(files.ReadFile(20, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution := countLit(enhance(algorithm, input, 2))
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution = countLit(enhance(algorithm, input, 50))
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func countLit(image [][]rune) int {
	var count int
	for _, line := range image {
		for _, pixel := range line {
			if pixel == '#' {
				count++
			}
		}
	}
	return count
}

func singleEnhance(algorithm []rune, input [][]rune, step int) [][]rune {
	directions := [9][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} // up-left, up, up-right, left, 0, right, down-left, down, down-right

	output := make([][]rune, len(input)+2)
	for i := -1; i <= len(input); i++ {
		output[i+1] = make([]rune, len(input[0])+2)
		for j := -1; j <= len(input[0]); j++ {
			var algorithmPos int
			for _, move := range directions {
				algorithmPos <<= 1
				iPos, jPos := i+move[0], j+move[1]
				// # = 1; . = 0;
				if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(input[0]) {
					if input[iPos][jPos] == '#' {
						algorithmPos |= 1
					}
				} else if algorithm[0] == '#' && step%2 == 0 {
					algorithmPos |= 1
				}
			}
			output[i+1][j+1] = algorithm[algorithmPos]
		}
	}
	return output
}

func enhance(algorithm []rune, input [][]rune, enhancements int) [][]rune {

	output := make([][]rune, len(input))
	copy(output, input)

	for step := 0; step < enhancements; step++ {
		output = singleEnhance(algorithm, output, step+1)
	}

	return output
}
