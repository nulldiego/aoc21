package main

import (
	"fmt"
	"log"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

func main() {
	start := time.Now()
	input := mapper.ToIntMatrix(files.ReadFile(11, "\r\n"), "")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countFlashes(input, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	input = mapper.ToIntMatrix(files.ReadFile(11, "\r\n"), "")
	start = time.Now()
	solution, err = stepsUntilSync(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

type xy struct {
	x, y int
}

func flash(input [][]int, flashed map[xy]xy) map[xy]xy {
	flashedRightNow := make(map[xy]xy, 100)
	iDirection := [8]int{-1, 1, 0, 0, -1, -1, 1, 1} // up, down, left, right, up-left, up-right, down-left, down-right
	jDirection := [8]int{0, 0, -1, 1, -1, 1, -1, 1} // up, down, left, right, up-left, up-right, down-left, down-right
	for i, row := range input {
		for j, num := range row {
			if num > 9 {
				if _, ok := flashed[xy{x: j, y: i}]; !ok {
					input[i][j] = 0
					flashed[xy{x: j, y: i}] = xy{x: j, y: i}
					flashedRightNow[xy{x: j, y: i}] = xy{x: j, y: i}
				}
			}
		}
	}

	for _, point := range flashedRightNow {
		for moveIdx, iMove := range iDirection {
			iPos, jPos := point.y+iMove, point.x+jDirection[moveIdx]
			if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(input[0]) {
				if _, found := flashed[xy{x: jPos, y: iPos}]; !found {
					input[iPos][jPos] = input[iPos][jPos] + 1
				}
			}
		}
	}

	if len(flashedRightNow) > 0 {
		return flash(input, flashed)
	}
	return flashed
}

func countFlashes(input [][]int, step int) (int, error) {
	var count int

	for i := 0; i < step; i++ {
		flashed := make(map[xy]xy, 100)
		for i, row := range input {
			for j, num := range row {
				input[i][j] = num + 1
			}
		}
		flash(input, flashed)
		count += len(flashed)
	}
	return count, nil
}

func stepsUntilSync(input [][]int) (int, error) {
	var count int

	for true {
		count++
		flashed := make(map[xy]xy, 100)
		for i, row := range input {
			for j, num := range row {
				input[i][j] = num + 1
			}
		}
		flash(input, flashed)

		if len(flashed) == 100 {
			return count, nil
		}
	}
	return count, nil
}
