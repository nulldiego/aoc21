package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"aoc21/utils/files"
)

func buildGraph(input []string, delimiter string) map[string][]string {
	result := map[string][]string{}
	for _, line := range input {
		path := strings.Split(line, delimiter)
		result[path[0]] = append(result[path[0]], path[1])
		result[path[1]] = append(result[path[1]], path[0])
	}
	return result
}

func main() {
	start := time.Now()
	input := buildGraph(files.ReadFile(12, "\r\n"), "-")
	fmt.Printf("Data readed in %v \n\n", time.Since(start))

	// Part 1
	start = time.Now()
	solution, err := countPaths(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// Part 2
	start = time.Now()
	solution, err = countPathsDoubleVisit(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}

func countPaths(input map[string][]string) (int, error) {
	var count int
	visited := map[string]bool{}
	var search func(node string)

	search = func(node string) {
		if node == "end" {
			count++
			return
		}
		if visited[node] && strings.ToLower(node) == node {
			return
		}

		visited[node] = true
		for _, next := range input[node] {
			search(next)
		}
		visited[node] = false
	}
	search("start")

	return count, nil
}

func countPathsDoubleVisit(input map[string][]string) (int, error) {
	var count int
	visited := map[string]bool{}
	var search func(node string, hasDouble bool)

	search = func(node string, hasDouble bool) {
		if node == "end" {
			count++
			return
		}
		if visited[node] && strings.ToLower(node) == node {
			if hasDouble || node == "start" {
				return
			}
			hasDouble = true
		}

		prevState := visited[node]
		visited[node] = true
		for _, next := range input[node] {
			search(next, hasDouble)
		}
		visited[node] = prevState
	}
	search("start", false)

	return count, nil
}
