package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
	"aoc21/utils/maths"
)

func main() {
	start := time.Now()
	input := mapper.ToIntSlice(files.ReadFile(7, ","))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := minFuelConstant(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = minFuelIncremental(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func minFuelConstant(input []int) (int, error) {
	return sumFuel(input, maths.Median(input), func(distance int) int {
		return distance
	}), nil
}

func minFuelIncremental(input []int) (int, error) {
	target := maths.Mean(input)
	fuelCost := func(distance int) int {
		return distance * (distance + 1) / 2
	}
	return maths.MinInt(sumFuel(input, int(math.Ceil(target)), fuelCost), sumFuel(input, int(math.Floor(target)), fuelCost)), nil
}

func sumFuel(input []int, target int, fuelCost func(distance int) int) int {
	var fuel int
	for _, crab := range input {
		fuel += fuelCost(maths.Abs(target - crab))
	}
	return fuel
}
