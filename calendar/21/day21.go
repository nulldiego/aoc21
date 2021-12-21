package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc21/utils/files"
)

func mustAtoi(input string) int64 {
	res, _ := strconv.ParseInt(input, 10, 64)
	return res
}

func parseInput(input []string) (int64, int64) {
	player1 := strings.Fields(input[0])
	player2 := strings.Fields(input[1])
	return mustAtoi(player1[len(player1)-1]), mustAtoi(player2[len(player2)-1])
}

func main() {
	start := time.Now()
	player1, player2 := parseInput(files.ReadFile(21, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution := playGame(player1, player2)
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution = playQuantumGame(player1, player2)
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func playQuantumGame(player1, player2 int64) int64 {
	wins := [2]int64{0, 0}
	diceSums := map[int64]int64{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

	// 3 - 1 1 1
	// 4 - 1 1 2, 1 2 1, 2 1 1
	// 5 - 1 1 3, 1 2 2, 1 3 1, 2 1 2, 2 2 1, 3 1 1
	// 6 - 1 2 3, 1 3 2, 2 1 3, 2 2 2, 2 3 1, 3 1 2, 3 2 1
	// 7 - 1 3 3, 2 2 3, 2 3 2, 3 1 3, 3 2 2, 3 3 1
	// 8 - 2 3 3, 3 2 3, 3 3 3
	// 9 - 3 3 3

	var quantum func(scores, positions []int64, turn int, multiplier int64)
	quantum = func(scores, positions []int64, turn int, multiplier int64) {
		for sum, universes := range diceSums {
			scoresUniverse := make([]int64, 2)
			copy(scoresUniverse, scores)
			positionsUniverse := make([]int64, 2)
			copy(positionsUniverse, positions)
			positionsUniverse[turn] = (positionsUniverse[turn] + sum) % 10
			if positionsUniverse[turn] == 0 {
				scoresUniverse[turn] += 10
			} else {
				scoresUniverse[turn] += positionsUniverse[turn]
			}
			if scoresUniverse[turn] >= 21 {
				wins[turn] += multiplier * universes
			} else {
				quantum([]int64{scoresUniverse[0], scoresUniverse[1]}, []int64{positionsUniverse[0], positionsUniverse[1]}, (turn+1)%2, multiplier*universes)
			}
		}
	}
	quantum([]int64{0, 0}, []int64{player1, player2}, 0, 1)

	if wins[0] > wins[1] {
		return wins[0]
	}
	return wins[1]
}

func playGame(player1, player2 int64) int64 {
	var count int64
	scores := [2]int64{0, 0}
	positions := [2]int64{player1, player2}

	for true {
		dice := int64(0)
		for i := 0; i < 3; i++ {
			count++
			dice += count
		}
		turn := (count + 1) % 2
		positions[turn] = (positions[turn] + dice) % 10
		if positions[turn] == 0 {
			scores[turn] += 10
		} else {
			scores[turn] += positions[turn]
		}
		if scores[turn] >= 1000 {
			return scores[(turn+1)%2] * count
		}
	}
	panic("something went wrong")
}
