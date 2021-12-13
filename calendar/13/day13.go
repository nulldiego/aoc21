package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"aoc21/utils/files"
)

func mustAtoi(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func buildCoordinates(input []string) [][2]int {
	res := [][2]int{}
	for _, line := range input {
		xy := strings.Split(line, ",")
		res = append(res, [2]int{mustAtoi(xy[0]), mustAtoi(xy[1])})
	}
	return res
}

func main() {
	start := time.Now()
	input := files.ReadFile(13, "\r\n\r\n")
	coordinates := buildCoordinates(strings.Split(input[0], "\r\n"))
	instructions := strings.Split(input[1], "\r\n")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countDotsAfterFirstFold(coordinates, instructions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	coordinates = buildCoordinates(strings.Split(input[0], "\r\n"))
	start = time.Now()
	resolve(coordinates, instructions)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func countDotsAfterFirstFold(coordinates [][2]int, instructions []string) (int, error) {
	fields := strings.Fields(instructions[0])
	fold := strings.Split(fields[len(fields)-1], "=")
	foldPos := mustAtoi(fold[1])
	if fold[0] == "y" {
		for i, coordinate := range coordinates {
			if coordinate[1] > foldPos {
				coordinates[i] = [2]int{coordinate[0], foldPos - (coordinate[1] - foldPos)}
			}
		}
	}
	if fold[0] == "x" {
		for i, coordinate := range coordinates {
			if coordinate[0] > foldPos {
				coordinates[i] = [2]int{foldPos - (coordinate[0] - foldPos), coordinate[1]}
			}
		}
	}
	coordinateMap := make(map[string]bool)
	for _, coordinate := range coordinates {
		coordinateMap[fmt.Sprintf("%d,%d", coordinate[0], coordinate[1])] = true
	}

	return len(coordinateMap), nil
}

func resolve(coordinates [][2]int, instructions []string) {
	var sizeX, sizeY int
	for _, instruction := range instructions {
		fields := strings.Fields(instruction)
		fold := strings.Split(fields[len(fields)-1], "=")
		foldPos := mustAtoi(fold[1])
		if fold[0] == "y" {
			sizeY = foldPos
			for i, coordinate := range coordinates {
				if coordinate[1] > foldPos {
					coordinates[i] = [2]int{coordinate[0], foldPos - (coordinate[1] - foldPos)}
				}
			}
			continue
		}
		if fold[0] == "x" {
			sizeX = foldPos
			for i, coordinate := range coordinates {
				if coordinate[0] > foldPos {
					coordinates[i] = [2]int{foldPos - (coordinate[0] - foldPos), coordinate[1]}
				}
			}
			continue
		}
	}
	coordinateMap := make(map[string]bool)
	for _, coordinate := range coordinates {
		coordinateMap[fmt.Sprintf("%d,%d", coordinate[0], coordinate[1])] = true
	}

	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if _, found := coordinateMap[fmt.Sprintf("%d,%d", x, y)]; found {
				print("#")
			} else {
				print(".")
			}
		}
		print("\n")
	}
}
