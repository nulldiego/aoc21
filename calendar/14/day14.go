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
	solution, err := countAfterSteps(starting, rules, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countAfterSteps(starting, rules, 40)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func countAfterSteps(starting string, rules map[string]string, steps int) (int, error) {
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
	for i := 0; i < steps; i++ {
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
