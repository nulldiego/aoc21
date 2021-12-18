package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	"aoc21/utils/files"
)

func mustAtoi(input string) int64 {
	res, _ := strconv.ParseInt(input, 10, 64)
	return res
}

func parseInput(input []string) [][]interface{} {
	snailNumbers := [][]interface{}{}
	re := regexp.MustCompile("[0-9]+")
	for _, line := range input {
		snailNumber := []interface{}{}
		pos := 0
		for pos < len(line) {
			char := line[pos : pos+1]
			if _, err := strconv.Atoi(char); err == nil {
				number := re.FindString(line[pos:])
				snailNumber = append(snailNumber, mustAtoi(number))
				pos += len(number)
				continue
			}
			if char != "," {
				snailNumber = append(snailNumber, char)
			}
			pos++
		}
		snailNumbers = append(snailNumbers, snailNumber)
	}
	return snailNumbers
}

func main() {
	start := time.Now()
	input := parseInput(files.ReadFile(18, "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution := magnitude(sumLines(input))
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution = maxMagnitude(input)
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func magnitude(a []interface{}) int64 { // 3*First + 2*Second
	if len(a) == 1 {
		return a[0].(int64)
	}

	switch v := a[1].(type) {
	case string: // it's a "["
		pos := 1
		opening := 0
		for true {
			switch val := a[pos].(type) {
			case string:
				if val == "[" {
					opening++
				} else {
					opening--
				}
				if opening == 0 {
					return 3*magnitude(a[1:pos+1]) + 2*magnitude(a[pos+1:len(a)-1])
				}
			}
			pos++
		}

	case int64: // it's a number - [1 [1 3]]
		return 3*v + 2*magnitude(a[2:len(a)-1])
	}
	return 0
}

func explode(a []interface{}, pairPos int) []interface{} {
	num1, num2 := a[pairPos+1].(int64), a[pairPos+2].(int64)
	a = append(append(a[:pairPos], int64(0)), a[pairPos+4:]...) // replace pair (snail number) with 0
left:
	for i := pairPos - 1; i >= 0; i-- {
		switch v := a[i].(type) {
		case int64:
			a[i] = int64(v + num1)
			break left
		}
	}
right:
	for i := pairPos + 1; i < len(a); i++ {
		switch v := a[i].(type) {
		case int64:
			a[i] = int64(v + num2)
			break right
		}
	}
	return a
}

func split(a []interface{}, pos int, value int64) []interface{} {
	div := float64(value) / float64(2)
	floor := int64(math.Floor(div))
	ceil := int64(math.Ceil(div))
	aux := make([]interface{}, len(a))
	copy(aux, a)
	aux = append(append(append(append(aux[:pos], "["), floor), ceil), "]")
	return append(aux, a[pos+1:]...)
}

func reduce(a []interface{}) []interface{} {
	var openedPairs int
	for i, char := range a {
		if char == "[" {
			openedPairs++
		} else if char == "]" {
			openedPairs--
		}
		if openedPairs > 4 { // explode
			return reduce(explode(a, i))
		}
	}
	for i, char := range a {
		switch v := char.(type) {
		case int64:
			if v > 9 {
				return reduce(split(a, i, v))
			}
		}
	}
	return a
}

func sum(a, b []interface{}) []interface{} {
	result := []interface{}{"["}
	result = append(result, append(append(a, b...), "]")...)
	return reduce(result)
}

func sumLines(input [][]interface{}) []interface{} {
	result := input[0]
	for _, line := range input[1:] {
		result = sum(result, line)
	}
	return result
}

func maxMagnitude(input [][]interface{}) int64 {
	max := int64(0)
	for i, snail1 := range input {
		for j, snail2 := range input {
			if i != j {
				mag := magnitude(sum(snail1, snail2))
				if mag > max {
					max = mag
				}
			}
		}
	}
	return max
}
