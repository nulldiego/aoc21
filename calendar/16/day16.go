package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"aoc21/utils/files"
)

func main() {
	start := time.Now()
	input := files.ReadFile(16, "\r\n")[0]
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1 & 2
	start = time.Now()
	verSum, result, err := solvePackets(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(verSum)
	fmt.Println(result)
	fmt.Printf("Part 1 and 2 solved in %v \n\n", time.Since(start))
}

func solvePackets(input string) (int, int64, error) {
	hexMap := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	operator := map[int]func(values []int64) int64{
		0: func(values []int64) int64 { // sum
			var result int64
			for _, val := range values {
				result += val
			}
			return result
		},
		1: func(values []int64) int64 { // product
			result := int64(1)
			for _, val := range values {
				result *= val
			}
			return result
		},
		2: func(values []int64) int64 { // min
			result := int64(math.MaxInt)
			for _, val := range values {
				if val < result {
					result = val
				}
			}
			return result
		},
		3: func(values []int64) int64 { // max
			var result int64
			for _, val := range values {
				if val > result {
					result = val
				}
			}
			return result
		},
		5: func(values []int64) int64 { // more
			var result int64
			if values[0] > values[1] {
				result = 1
			}
			return result
		},
		6: func(values []int64) int64 { // less
			var result int64
			if values[0] < values[1] {
				result = 1
			}
			return result
		},
		7: func(values []int64) int64 { // equal
			var result int64
			if values[0] == values[1] {
				result = 1
			}
			return result
		},
	}

	var binary string
	for _, hex := range input {
		binary += hexMap[hex]
	}

	var parsePackets func(string, int) (newPos int, result int64)

	verSum := 0
	parsePackets = func(binaryStr string, pos int) (int, int64) {
		version, _ := strconv.ParseInt(binaryStr[pos:pos+3], 2, 32)
		packetType, _ := strconv.ParseInt(binaryStr[pos+3:pos+6], 2, 32)
		verSum += int(version)
		pos += 6

		if packetType == 4 { // literal
			var binaryNumberStr string
			for true {
				binaryNumberStr += binaryStr[pos+1 : pos+5]
				if binaryStr[pos] == '0' { // last group
					literal, _ := strconv.ParseInt(binaryNumberStr, 2, 64)
					return pos + 5, literal
				}
				pos += 5
			}
		}

		// operator
		var values []int64
		length := 15
		bySubPackets := false
		if binaryStr[pos] == '1' { // 11 length
			length = 11
			bySubPackets = true
		}
		pos += 1
		subPacketsLength, _ := strconv.ParseInt(binaryStr[pos:pos+length], 2, 32)
		pos += length
		if !bySubPackets {
			endPos := pos + int(subPacketsLength)
			for true {
				newPos, val := parsePackets(binaryStr, pos)
				pos = newPos
				values = append(values, val)
				if pos >= endPos {
					break
				}
			}
		} else {
			for i := 0; i < int(subPacketsLength); i++ {
				newPos, val := parsePackets(binaryStr, pos)
				pos = newPos
				values = append(values, val)
			}
		}

		return pos, operator[int(packetType)](values)
	}

	_, res := parsePackets(binary, 0)
	return verSum, res, nil

}
