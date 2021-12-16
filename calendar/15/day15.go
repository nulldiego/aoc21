package main

import (
	"fmt"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"

	"github.com/yourbasic/graph"
)

func buildGraph(input [][]int) (*graph.Mutable, int) {
	g := graph.New(len(input) * len(input[0]))

	pos := -1
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right

	// 0 1 2 3 4 5 6 7 8 9
	// 10 11 12 13 14 15 16 17 18 19
	// 20 21 22 ...
	lenRow := len(input[0])
	for i, row := range input {
		for j := range row {
			pos++
			for _, move := range directions {
				iPos, jPos := i+move[0], j+move[1]
				if 0 <= iPos && iPos < len(input) && 0 <= jPos && jPos < len(input[0]) {
					otherPos := pos + move[0]*lenRow + move[1]
					g.AddCost(pos, otherPos, int64(input[iPos][jPos]))
				}
			}
		}
	}
	return g, pos
}

func expandIntMatrix5x(matrix [][]int) [][]int {
	sliceToReturn := make([][]int, len(matrix)*5) // 5 times rows
	lenRow := len(matrix[0])
	for i := 0; i < len(matrix)*5; i++ {
		sliceToReturn[i] = []int{}
		for j := 0; j < lenRow*5; j++ {
			toAppend := 0
			if j < lenRow && i < lenRow {
				toAppend = matrix[i][j]
			} else if j >= lenRow {
				toAppend = sliceToReturn[i][j-lenRow] + 1 // 1 2 3 2 3 4
				if toAppend == 10 {
					toAppend = 1
				}
			} else if i >= lenRow {
				toAppend = sliceToReturn[i-lenRow][j] + 1
				if toAppend == 10 {
					toAppend = 1
				}
			}
			sliceToReturn[i] = append(sliceToReturn[i], toAppend)
		}
	}
	return sliceToReturn
}

func main() {
	start := time.Now()
	inputMatrix := mapper.ToIntMatrix(files.ReadFile(15, "\r\n"), "")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	input, lastPos := buildGraph(inputMatrix)
	_, dist := graph.ShortestPath(input, 0, lastPos)
	fmt.Println(dist)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	inputMatrix = expandIntMatrix5x(inputMatrix)
	input, lastPos = buildGraph(inputMatrix)
	_, dist = graph.ShortestPath(input, 0, lastPos)
	fmt.Println(dist)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}
