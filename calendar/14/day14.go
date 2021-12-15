package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"aoc21/utils/files"
)

func mustAtoi(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func buildRules(input []string) map[string]string {
	res := map[string]string{}
	for _, line := range input {
		rule := strings.Split(line, " -> ")
		res[rule[0]] = rule[1]
	}
	return res
}

func main() {
	start := time.Now()
	input := files.ReadFile(14, "\r\n\r\n")
	starting := input[0]
	rules := buildRules(strings.Split(input[1], "\r\n"))
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countAfter10Steps(starting, rules)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countAfter40Steps(starting, rules)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func countAfter40Steps(starting string, rules map[string]string) (int, error) {
	qty := map[string]int{}
	for _, letter := range strings.Split(starting, "") {
		qty[letter]++
	}

	qtyPairs := map[string]int{}
	for i, letter := range strings.Split(starting, "")[1:] {
		lookUp := string(starting[i]) + string(letter)
		qtyPairs[lookUp]++
	}

	qtyPairsCopy := map[string]int{}
	for k, v := range qtyPairs {
		qtyPairsCopy[k] = v
	}
	for i := 0; i < 40; i++ {
		qtyPairs = map[string]int{}
		for k, v := range qtyPairsCopy {
			qtyPairs[k] = v
		}
		for key, qtyPair := range qtyPairs {
			if rule, found := rules[key]; found {
				qtyPairsCopy[key] -= qtyPair
				qtyPairsCopy[string(key[0])+rule] += qtyPair
				qtyPairsCopy[rule+string(key[1])] += qtyPair
				qty[rule] += qtyPair
			}
		}
	}

	var min, max int
	for _, qt := range qty {
		if min == 0 && max == 0 {
			min = qt
			max = qt
			continue
		}
		if min > qt {
			min = qt
		}
		if max < qt {
			max = qt
		}
	}
	return max - min, nil
}

func countAfter10Steps(starting string, rules map[string]string) (int, error) {
	qty := map[string]int{}
	for _, letter := range strings.Split(starting, "") {
		if _, found := qty[letter]; found {
			qty[letter] = qty[letter] + 1
		} else {
			qty[letter] = 1
		}
	}

	result := starting
	resultCopy := result
	for i := 0; i < 10; i++ {
		result := resultCopy
		for i, letter := range strings.Split(result, "")[1:] {
			lookUp := string(result[i]) + string(letter)
			if rule, found := rules[lookUp]; found {
				idx := strings.Index(resultCopy, lookUp)
				resultCopy = resultCopy[:idx+1] + rule + resultCopy[idx+1:]

				if _, found := qty[rule]; found {
					qty[rule] = qty[rule] + 1
				} else {
					qty[rule] = 1
				}
			}
		}
	}

	var min, max int
	for _, qt := range qty {
		if min == 0 && max == 0 {
			min = qt
			max = qt
			continue
		}
		if min > qt {
			min = qt
		}
		if max < qt {
			max = qt
		}
	}
	return max - min, nil
}
