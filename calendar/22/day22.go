package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/mapper"
	"aoc21/utils/maths"
)

type step struct {
	on bool
	x  [2]int
	y  [2]int
	z  [2]int
}

func parseInput(input []string) []step {
	re := regexp.MustCompile("-*[0-9]+")

	steps := []step{}
	for _, line := range input {
		instruction := strings.Fields(line)
		coords := mapper.ToIntSlice(re.FindAllString(instruction[1], 6))
		on := false
		if instruction[0] == "on" {
			on = true
		}
		steps = append(steps, step{
			on: on,
			x:  [2]int{coords[0], coords[1] + 1},
			y:  [2]int{coords[2], coords[3] + 1},
			z:  [2]int{coords[4], coords[5] + 1},
		})
	}
	return steps
}

func main() {
	start := time.Now()
	input := parseInput(files.ReadFile(22, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution := reboot(input[:20])
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution = reboot(input)
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func getIntersectCube(a, b step) (step, bool) {
	if a.x[1] <= b.x[0] || a.x[0] >= b.x[1] || a.y[1] <= b.y[0] || a.y[0] >= b.y[1] || a.z[1] <= b.z[0] || a.z[0] >= b.z[1] {
		return step{}, false
	}
	return step{
		on: b.on,
		x:  [2]int{maths.MaxInt(a.x[0], b.x[0]), maths.MinInt(a.x[1], b.x[1])},
		y:  [2]int{maths.MaxInt(a.y[0], b.y[0]), maths.MinInt(a.y[1], b.y[1])},
		z:  [2]int{maths.MaxInt(a.z[0], b.z[0]), maths.MinInt(a.z[1], b.z[1])},
	}, true
}

func countNonIntersected(in step, future []step) int64 {
	var pending []step

	for _, step := range future {
		if innerCube, ok := getIntersectCube(in, step); ok {
			pending = append(pending, innerCube)
		}
	}

	total := int64(maths.Abs(in.x[1]-in.x[0])) * int64(maths.Abs(in.y[1]-in.y[0])) * int64(maths.Abs(in.z[1]-in.z[0]))
	for i, step := range pending {
		total -= countNonIntersected(step, pending[i+1:])
	}

	return total
}

func reboot(input []step) int64 {
	lit := int64(0)
	for i, step := range input {
		if step.on {
			lit += countNonIntersected(step, input[i+1:])
		}
	}
	return lit
}
