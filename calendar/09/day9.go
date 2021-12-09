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
		result[i] = make([]int, len(line))
		for j, num := range line {
			result[i][j] = int(num - '0')
		}
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
	for i, row := range input {
		for j, num := range row {
			if i == 0 {
				if j == 0 {
					if num < input[i+1][j] && num < input[i][j+1] {
						result += num + 1
					}
				} else if j == len(row)-1 {
					if num < input[i+1][j] && num < input[i][j-1] {
						result += num + 1
					}
				} else {
					if num < input[i+1][j] && num < input[i][j-1] && num < input[i][j+1] {
						result += num + 1
					}
				}
				continue
			}
			if i == len(input)-1 {
				if j == 0 {
					if num < input[i-1][j] && num < input[i][j+1] {
						result += num + 1
					}
				} else if j == len(row)-1 {
					if num < input[i-1][j] && num < input[i][j-1] {
						result += num + 1
					}
				} else {
					if num < input[i-1][j] && num < input[i][j-1] && num < input[i][j+1] {
						result += num + 1
					}
				}
				continue
			}
			if j == 0 {
				if num < input[i-1][j] && num < input[i+1][j] && num < input[i][j+1] {
					result += num + 1
				}
				continue
			}
			if j == len(row)-1 {
				if num < input[i-1][j] && num < input[i+1][j] && num < input[i][j-1] {
					result += num + 1
				}
				continue
			}
			if num < input[i-1][j] && num < input[i+1][j] && num < input[i][j-1] && num < input[i][j+1] {
				result += num + 1
			}
		}
	}
	return result, nil
}

func findBasin(i, j int, input [][]int) []string {
	var basinPositions []string
	// up
	if i > 0 {
		if input[i-1][j] > input[i][j] && input[i-1][j] < 9 {
			basinPositions = append(basinPositions, fmt.Sprintf("%d,%d", i-1, j))
		}
	}
	// down
	if i < len(input)-1 {
		if input[i+1][j] > input[i][j] && input[i+1][j] < 9 {
			basinPositions = append(basinPositions, fmt.Sprintf("%d,%d", i+1, j))
		}
	}
	// left
	if j > 0 {
		if input[i][j-1] > input[i][j] && input[i][j-1] < 9 {
			basinPositions = append(basinPositions, fmt.Sprintf("%d,%d", i, j-1))
		}
	}
	// right
	if j < len(input[0])-1 {
		if input[i][j+1] > input[i][j] && input[i][j+1] < 9 {
			basinPositions = append(basinPositions, fmt.Sprintf("%d,%d", i, j+1))
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
	basinSizes := make([]int, 3)
	for i, row := range input {
		for j, num := range row {
			if i == 0 {
				if j == 0 {
					if num < input[i+1][j] && num < input[i][j+1] {
						basinSize := len(findBasin(i, j, input)) + 1
						if basinSize > basinSizes[0] {
							basinSizes[0] = basinSize
						}
						sort.Ints(basinSizes)
					}
				} else if j == len(row)-1 {
					if num < input[i+1][j] && num < input[i][j-1] {
						basinSize := len(findBasin(i, j, input)) + 1
						if basinSize > basinSizes[0] {
							basinSizes[0] = basinSize
						}
						sort.Ints(basinSizes)
					}
				} else {
					if num < input[i+1][j] && num < input[i][j-1] && num < input[i][j+1] {
						basinSize := len(findBasin(i, j, input)) + 1
						if basinSize > basinSizes[0] {
							basinSizes[0] = basinSize
						}
						sort.Ints(basinSizes)
					}
				}
				continue
			}
			if i == len(input)-1 {
				if j == 0 {
					if num < input[i-1][j] && num < input[i][j+1] {
						basinSize := len(findBasin(i, j, input)) + 1
						if basinSize > basinSizes[0] {
							basinSizes[0] = basinSize
						}
						sort.Ints(basinSizes)
					}
				} else if j == len(row)-1 {
					if num < input[i-1][j] && num < input[i][j-1] {
						basinSize := len(findBasin(i, j, input)) + 1
						if basinSize > basinSizes[0] {
							basinSizes[0] = basinSize
						}
						sort.Ints(basinSizes)
					}
				} else {
					if num < input[i-1][j] && num < input[i][j-1] && num < input[i][j+1] {
						basinSize := len(findBasin(i, j, input)) + 1
						if basinSize > basinSizes[0] {
							basinSizes[0] = basinSize
						}
						sort.Ints(basinSizes)
					}
				}
				continue
			}
			if j == 0 {
				if num < input[i-1][j] && num < input[i+1][j] && num < input[i][j+1] {
					basinSize := len(findBasin(i, j, input)) + 1
					if basinSize > basinSizes[0] {
						basinSizes[0] = basinSize
					}
					sort.Ints(basinSizes)
				}
				continue
			}
			if j == len(row)-1 {
				if num < input[i-1][j] && num < input[i+1][j] && num < input[i][j-1] {
					basinSize := len(findBasin(i, j, input)) + 1
					if basinSize > basinSizes[0] {
						basinSizes[0] = basinSize
					}
					sort.Ints(basinSizes)
				}
				continue
			}
			if num < input[i-1][j] && num < input[i+1][j] && num < input[i][j-1] && num < input[i][j+1] {
				basinSize := len(findBasin(i, j, input)) + 1
				if basinSize > basinSizes[0] {
					basinSizes[0] = basinSize
				}
				sort.Ints(basinSizes)
			}
		}
	}
	return basinSizes[0] * basinSizes[1] * basinSizes[2], nil
}
