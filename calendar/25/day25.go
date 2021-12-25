package main

import (
	"fmt"
	"time"

	"aoc21/utils/files"
)

func parseInput(input []string) (east map[[2]int]bool, south map[[2]int]bool, mapSize [2]int) {
	east = map[[2]int]bool{}
	south = map[[2]int]bool{}
	for i, row := range input {
		for j, char := range row {
			if char == '>' {
				east[[2]int{i, j}] = true
			} else if char == 'v' {
				south[[2]int{i, j}] = true
			}
		}
	}
	mapSize = [2]int{len(input), len(input[0])}
	return east, south, mapSize
}

func countSteps(east map[[2]int]bool, south map[[2]int]bool, mapSize [2]int) int {
	count := 0
	done := false
	for !done {
		done = true
		oldEast := map[[2]int]bool{}
		for k, v := range east {
			oldEast[k] = v
		}
		for eastCucumber := range oldEast {
			newPos := [2]int{eastCucumber[0], (eastCucumber[1] + 1) % mapSize[1]}
			if !oldEast[newPos] && !south[newPos] {
				delete(east, eastCucumber)
				east[newPos] = true
				done = false
			}
		}

		oldSouth := map[[2]int]bool{}
		for k, v := range south {
			oldSouth[k] = v
		}
		for southCucumber := range oldSouth {
			newPos := [2]int{(southCucumber[0] + 1) % mapSize[0], southCucumber[1]}
			if !east[newPos] && !oldSouth[newPos] {
				delete(south, southCucumber)
				south[newPos] = true
				done = false
			}
		}
		count++
	}

	return count
}

func main() {
	start := time.Now()
	east, south, mapSize := parseInput(files.ReadFile(25, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution := countSteps(east, south, mapSize)
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// // Part 2
	// start = time.Now()
	// solution = reboot(input)
	// fmt.Println(solution)
	// fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}
