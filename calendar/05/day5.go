package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

func main() {
	start := time.Now()
	input := files.ReadFile(5, "\r\n")
	lines := make([][][]int, len(input)) // n lines
	for iline, line := range input {
		lines[iline] = make([][]int, 2) // 2 positions
		for iposition, position := range strings.Split(line, " -> ") {
			lines[iline][iposition] = make([]int, 2) // 2 coordinates
			lines[iline][iposition] = mapper.ToIntSlice(strings.Split(position, ","))
		}
	}
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countOverlaps(lines, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countOverlaps(lines, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func countOverlaps(lines [][][]int, diagonals bool) (int, error) {
	diagram := make(map[string]int)
	var count int
	for _, line := range lines {
		fromX, toX, fromY, toY := line[0][0], line[1][0], line[0][1], line[1][1] // fromY and toY always correspond to its x coordinate
		if !diagonals && fromX != toX && fromY != toY {
			continue
		}
		if fromX > toX {
			fromX, toX = toX, fromX
			fromY, toY = toY, fromY
		}
		distance := maxInt(abs(toX-fromX), abs(toY-fromY))
		for sum := 0; sum <= distance; sum++ {
			x, y := fromX, fromY
			if fromX != toX {
				x = fromX + sum
			}
			if fromY > toY {
				y = fromY - sum
			} else if fromY < toY {
				y = fromY + sum
			}
			coord := fmt.Sprintf("%d,%d", x, y)
			if _, found := diagram[coord]; found {
				diagram[coord]++
				if diagram[coord] == 2 {
					count++
				}
			} else {
				diagram[coord] = 1
			}
		}
	}
	return count, nil
}
