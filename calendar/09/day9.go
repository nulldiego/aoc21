package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

func toIntMatrix(input []string) [][]int {
	result := make([][]int, len(input))
	for i, line := range input {
		result[i] = mapper.ToIntSlice(strings.Split(line, ""))
	}
	return result
}

func main() {
	start := time.Now()
	input := toIntMatrix(files.ReadFile(9, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := sumLowPoints(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = multiplyBasins(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func sumLowPoints(input [][]int) (int, error) {
	var result int
	iDirection := [4]int{-1, 1, 0, 0} // up, down, left, right
	jDirection := [4]int{0, 0, -1, 1} // up, down, left, right
	for i, row := range input {
		for j, num := range row {
			isLowPoint := true
			for moveIdx, iMove := range iDirection {
				iPos, jPos := i+iMove, j+jDirection[moveIdx]
				if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(row) && input[iPos][jPos] <= num {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				result += num + 1
			}
		}
	}
	return result, nil
}

func findBasin(i, j int, input [][]int) []string {
	var basinPositions []string

	iDirection := [4]int{-1, 1, 0, 0} // up, down, left, right
	jDirection := [4]int{0, 0, -1, 1} // up, down, left, right
	for moveIdx, iMove := range iDirection {
		iPos, jPos := i+iMove, j+jDirection[moveIdx]
		if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(input[0]) && input[iPos][jPos] > input[i][j] && input[iPos][jPos] < 9 {
			basinPositions = append(basinPositions, fmt.Sprintf("%d,%d", i-1, j))
		}
	}

	for _, basinPosition := range basinPositions {
		position := mapper.ToIntSlice(strings.Split(basinPosition, ","))
		basinPositions = append(basinPositions, findBasin(position[0], position[1], input)...)
	}

	// remove duplicates
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range basinPositions {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func multiplyBasins(input [][]int) (int, error) {
	var basinSizes []int
	iDirection := [4]int{-1, 1, 0, 0} // up, down, left, right
	jDirection := [4]int{0, 0, -1, 1} // up, down, left, right
	for i, row := range input {
		for j, num := range row {
			isLowPoint := true
			for moveIdx, iMove := range iDirection {
				iPos, jPos := i+iMove, j+jDirection[moveIdx]
				if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(row) && input[iPos][jPos] <= num {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				basinSize := len(findBasin(i, j, input)) + 1
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3], nil
}
