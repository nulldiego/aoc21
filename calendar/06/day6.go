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
	input := mapper.ToIntSlice(files.ReadFile(6, ","))
	fishStates, totalFish := readFish(input)
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countFish(fishStates, totalFish, 80)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countFish(fishStates, totalFish, 256)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func readFish(input []int) ([]int, int) {
	var total int
	fishStates := make([]int, 9) // each position holds the number of fish in that day
	for _, fishState := range input {
		fishStates[fishState]++
		total++
	}
	return fishStates, total
}

func countFish(fishStates []int, fishCount, days int) (int, error) {
	for day := 0; day < days; day++ {
		todayFish := fishStates[0]                     // every fish in state 0 produces a new fish
		fishStates = append(fishStates[1:], todayFish) // every fish reduces 1 day its state, newFish start at the end at state 8
		fishStates[6] += todayFish                     // fish that produced a fish today, produce a new fish in 7 days
		fishCount += todayFish
	}

	return fishCount, nil
}
