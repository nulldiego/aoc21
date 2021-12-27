package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"aoc21/utils/files"
)

func mustAtoi(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func parseInput(input []string) (xAdds []int, yAdds []int, zDivs []int) {
	for i, row := range input {
		if (i-5)%18 == 0 {
			xAdds = append(xAdds, mustAtoi(strings.Fields(row)[2]))
		} else if (i-15)%18 == 0 {
			yAdds = append(yAdds, mustAtoi(strings.Fields(row)[2]))
		} else if (i-4)%18 == 0 {
			zDivs = append(zDivs, mustAtoi(strings.Fields(row)[2]))
		}
	}
	return xAdds, yAdds, zDivs
}

func backward(xAdd, yAdd, zDiv, z2, w int) (zs []int) {
	x := z2 - w - yAdd
	if x%26 == 0 {
		zs = append(zs, int(math.Floor(float64(x)/float64(26)))*zDiv)
	}

	if 0 <= w-xAdd && w-xAdd < 26 {
		z0 := z2 * zDiv
		zs = append(zs, w-xAdd+z0)
	}

	return zs
}

func find(xAdds, yAdds, zDivs, ws []int) []int {
	zs := map[int]int{0: 0}
	result := map[int][]int{}

	for i := len(xAdds) - 1; i >= 0; i-- {
		newZs := map[int]int{}
		for _, w := range ws {
			for _, z := range zs {
				z0s := backward(xAdds[i], yAdds[i], zDivs[i], z, w)
				for _, z0 := range z0s {
					newZs[z0] = z0
					result[z0] = append([]int{w}, result[z]...)
				}
			}
		}
		zs = newZs
	}
	return result[0]
}

func main() {
	start := time.Now()
	x, y, z := parseInput(files.ReadFile(24, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	ws := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	solution := find(x, y, z, ws)
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	ws = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	solution = find(x, y, z, ws)
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}
