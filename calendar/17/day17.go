package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
	"aoc21/utils/maths"
)

func main() {
	start := time.Now()
	input := files.ReadFile(17, ", ")
	x_range := mapper.ToIntSlice(strings.Split(strings.Split(input[0], "=")[1], ".."))
	y_range := mapper.ToIntSlice(strings.Split(strings.Split(input[1], "=")[1], ".."))
	sort.Ints(x_range)
	sort.Ints(y_range)
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := yMaximize(x_range, y_range)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countOptions(x_range, y_range)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func summation(n int) int {
	return (n*n + n) / 2
}

func yMaximize(xRange, yRange []int) (int, error) {
	yMax := maths.MaxInt(maths.Abs(yRange[0]), maths.Abs(yRange[1])) - 1 // any higher would miss the target due to increasing velocity downward
	return summation(yMax), nil
}

func minSummation(result int) int {
	for n := 0; true; n++ {
		if summation(n) >= result {
			return n
		}
	}
	return 0
}

func shootsTarget(x, y int, xRange, yRange []int) bool {
	yMin := maths.MinInt(yRange[0], yRange[1])
	xVel := x
	yVel := y
	xPos, yPos := 0, 0
	for true {
		xPos += xVel
		yPos += yVel
		yVel--

		if xVel == 0 && (xPos < xRange[0] || xPos > xRange[1]) {
			return false
		} else if xVel > 0 {
			xVel--
		} else if xVel < 0 {
			xVel++
		}

		if xRange[0] <= xPos && xPos <= xRange[1] && yRange[0] <= yPos && yPos <= yRange[1] {
			return true
		}
		if xPos > xRange[1] || yPos < yMin {
			return false
		}
	}
	return false
}

func countOptions(xRange, yRange []int) (int, error) {
	yMax := maths.MaxInt(maths.Abs(yRange[0]), maths.Abs(yRange[1])) - 1 // any higher would miss the target due to increasing velocity downward
	yMin := maths.MinInt(yRange[0], yRange[1])                           // any less would overshoot the target
	xMin := minSummation(xRange[0])                                      // any less wouldn't reach the target due to decreasing velocity
	xMax := maths.MaxInt(xRange[0], xRange[1])                           // any higher would overshoot the target

	count := 0

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if shootsTarget(x, y, xRange, yRange) {
				count++
			}
		}
	}

	return count, nil
}
