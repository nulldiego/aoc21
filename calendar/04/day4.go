package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
)

type board struct {
	solutions [][]int
	score     int
}

func buildBoards(input []string) []board {
	var boards []board
	for _, boardStr := range input {
		var board board
		board.solutions = make([][]int, 10) // 5 lines + 5 columns
		for i := range board.solutions {
			board.solutions[i] = make([]int, 5) // of 5 numbers
		}
		for i, boardLine := range strings.Split(boardStr, "\r\n") {
			for j, boardNumber := range mapper.ToIntSlice(strings.Fields(boardLine)) {
				board.solutions[i][j] = boardNumber   // line
				board.solutions[j+5][i] = boardNumber // column
				board.score += boardNumber
			}
		}
		boards = append(boards, board)
	}
	return boards
}

func main() {
	start := time.Now()
	input := files.ReadFile(4, "\r\n\r\n")
	drawNumbers := mapper.ToIntSlice(strings.Split(input[0], ","))
	boards := buildBoards(input[1:])
	fmt.Printf("Data parsed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := findWinnerScore(drawNumbers, boards)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = findLoserScore(drawNumbers, boards)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func findWinnerScore(draw []int, boards []board) (int, error) {
	for _, number := range draw {
		for idxBoard, board := range boards {
			for i, boardSolution := range board.solutions {
				for j, boardNumber := range boardSolution {
					if number == boardNumber { // remove number from possible solution (mark)
						board.solutions[i][j] = board.solutions[i][len(board.solutions[i])-1]
						board.solutions[i] = board.solutions[i][:len(board.solutions[i])-1]
						if i < 5 {
							boards[idxBoard].score -= number // only substract score on lines (avoid double substraction)
						}
						if len(board.solutions[i]) == 0 { // WINNER SOLUTION
							return boards[idxBoard].score * number, nil
						}
					}
				}
			}
		}
	}
	return 0, errors.New("couldn't find a winner")
}

func findLoserScore(draw []int, boards []board) (int, error) {
	for idxNumber, number := range draw {
		for idxBoard, board := range boards {
			for i, boardSolution := range board.solutions {
				for j, boardNumber := range boardSolution {
					if number == boardNumber { // remove number from possible solution (mark)
						board.solutions[i][j] = board.solutions[i][len(board.solutions[i])-1]
						board.solutions[i] = board.solutions[i][:len(board.solutions[i])-1]
						if i < 5 {
							boards[idxBoard].score -= number // only substract score on lines (avoid double substraction)
						}
						if len(board.solutions[i]) == 0 { // WINNER
							if len(boards) == 1 { // last board to win
								return boards[idxBoard].score * number, nil
							}
							return findLoserScore(draw[idxNumber:], append(boards[:idxBoard], boards[idxBoard+1:]...)) // keep going without winning board
						}
					}
				}
			}
		}
	}
	return 0, errors.New("couldn't find a winner")
}
