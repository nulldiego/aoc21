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

func flash(input [][]int, flashed map[[2]int]bool) map[[2]int]bool {
	var flashedRightNow [][2]int

	for i, row := range input {
		for j, num := range row {
			if num > 9 {
				ij := [2]int{i, j}
				if !flashed[ij] {
					input[i][j] = 0
					flashed[ij] = true
					flashedRightNow = append(flashedRightNow, ij)
				}
			}
		}
	}

	directions := [8][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} // up, down, left, right, up-left, up-right, down-left, down-right
	for _, point := range flashedRightNow {
		for _, move := range directions {
			iPos, jPos := point[0]+move[0], point[1]+move[1]
			if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(input[0]) {
				if !flashed[[2]int{iPos, jPos}] {
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
		flashed := make(map[[2]int]bool, 100)
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
		flashed := make(map[[2]int]bool, 100)
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
