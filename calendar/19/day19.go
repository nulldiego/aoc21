package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/maths"
)

func mustAtoi(input string) int {
	res, _ := strconv.ParseInt(input, 10, 32)
	return int(res)
}

func parseInput(input []string) [][][3]int {
	scanners := [][][3]int{}
	for _, scannerInput := range input {
		scanner := [][3]int{}
		for _, pos := range strings.Split(scannerInput, "\r\n")[1:] {
			xyz := strings.Split(pos, ",")
			scanner = append(scanner, [3]int{mustAtoi(xyz[0]), mustAtoi(xyz[1]), mustAtoi(xyz[2])})
		}
		scanners = append(scanners, scanner)
	}
	return scanners
}

func main() {
	start := time.Now()
	input := parseInput(files.ReadFile(19, "\r\n\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1 & 2
	start = time.Now()
	solution1, solution2 := sumScanners(input)
	fmt.Println(solution1)
	fmt.Println(solution2)
	fmt.Printf("Part 1 & 2 solved in %v \n\n", time.Since(start))
}

func roll(point [3]int) [3]int {
	return [3]int{point[0], point[2], -point[1]}
}

func rollScanner(scanner [][3]int) [][3]int {
	var result [][3]int
	for _, point := range scanner {
		result = append(result, roll(point))
	}
	return result
}

func turn(point [3]int, direction string) [3]int {
	switch direction {
	case "CLOCKWISE":
		return [3]int{point[1], -point[0], point[2]}
	case "COUNTER_CLOCKWISE":
		return [3]int{-point[1], point[0], point[2]}
	}
	panic("wrong direction")
}

func turnScanner(scanner [][3]int, direction string) [][3]int {
	var result [][3]int
	for _, point := range scanner {
		result = append(result, turn(point, direction))
	}
	return result
}

func sumWithOffset(a, b [][3]int, offset [3]int) ([][3]int, bool) {
	keys := make(map[[3]int]bool)
	sum := [][3]int{}
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			sum = append(sum, entry)
		}
	}
	for _, entry := range b {
		entryOffset := [3]int{entry[0] + offset[0], entry[1] + offset[1], entry[2] + offset[2]}
		if _, value := keys[entryOffset]; !value {
			keys[entryOffset] = true
			sum = append(sum, entryOffset)
		}
	}

	if len(a)+len(b) >= len(sum)+12 {
		return sum, true
	}
	return nil, false
}

func match(a, b [][3]int) ([][3]int, [3]int, bool) {
	for _, bPoint := range b {
		for _, aPoint := range a {
			offset := [3]int{aPoint[0] - bPoint[0], aPoint[1] - bPoint[1], aPoint[2] - bPoint[2]}
			if sum, ok := sumWithOffset(a, b, offset); ok {
				return sum, offset, ok
			}
		}
	}
	return nil, [3]int{}, false
}

func sum(a, b [][3]int) ([][3]int, [3]int, bool) {
	for rollIndex := 0; rollIndex < 6; rollIndex++ {
		b = rollScanner(b)
		if result, offset, ok := match(a, b); ok {
			return result, offset, ok
		}
		for turnIndex := 0; turnIndex < 3; turnIndex++ {
			if rollIndex%2 == 0 {
				b = turnScanner(b, "CLOCKWISE")
				if result, offset, ok := match(a, b); ok {
					return result, offset, ok
				}
				continue
			}
			b = turnScanner(b, "COUNTER_CLOCKWISE")
			if result, offset, ok := match(a, b); ok {
				return result, offset, ok
			}
		}
	}
	return a, [3]int{}, false
}

func manhattanDistance(a, b [3]int) int {
	return maths.Abs(a[0]-b[0]) + maths.Abs(a[1]-b[1]) + maths.Abs(a[2]-b[2])
}

func sumScanners(scanners [][][3]int) (int, int) {
	result := scanners[0]
	offsets := [][3]int{}

	done := map[int]bool{}
	done[0] = true

	for len(done) < len(scanners) {
		for i, scanner := range scanners[1:] {
			if done[i+1] {
				continue
			}
			resAux, offset, ok := sum(result, scanner)
			result = resAux
			if !ok {
				continue
			}
			done[i+1] = true
			offsets = append(offsets, offset)
		}
	}

	maxOffset := 0
	for i, offset := range offsets[:len(offsets)-1] {
		for _, offset2 := range offsets[i+1:] {
			maxOffset = maths.MaxInt(manhattanDistance(offset, offset2), maxOffset)
		}
	}

	return maxOffset, len(result)
}
