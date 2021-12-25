package main

import (
	"fmt"
	"regexp"
	"sort"
	"time"

	"aoc21/utils/files"
	"aoc21/utils/maths"
)

type state struct {
	hall  [11]string
	rooms [4][]string
}

func parseInput(input []string) state {
	re := regexp.MustCompile("[A-D]")

	inputState := state{
		hall:  [11]string{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		rooms: [4][]string{{}, {}, {}, {}},
	}

	for _, row := range input[2:4] {
		rowAmphipods := re.FindAllString(row, 4)
		for x, letter := range rowAmphipods {
			inputState.rooms[x] = append(inputState.rooms[x], letter)
		}
	}
	return inputState
}

// Are we done?
func everythingInItsRightPlace(amphipods state) bool {
	rightPlace := map[int]string{0: "A", 1: "B", 2: "C", 3: "D"}

	for i, room := range amphipods.rooms {
		for _, amphipod := range room {
			if amphipod != rightPlace[i] {
				return false
			}
		}
	}

	return true
}

// Can enter its room?
func roomAvailable(letter string, room []string) bool {
	for _, occupying := range room {
		if occupying != letter && occupying != "." {
			return false
		}
	}
	return true
}

// Position of room door in hall
func roomDoor(letter string) int {
	switch letter {
	case "A":
		return 2
	case "B":
		return 4
	case "C":
		return 6
	case "D":
		return 8
	}
	panic("wrong room, can't get room door")
}

// Distance to its room, position in its room, can get there?
func path(letter string, from int, hall [11]string, room []string) (dist, roomPos int, ok bool) {
	fromTo := []int{from, roomDoor(letter)}
	sort.Ints(fromTo)
	for _, occupying := range hall[fromTo[0]+1 : fromTo[1]] {
		if occupying != "." {
			return 0, 0, false
		}
	}
	dist = fromTo[1] - fromTo[0]
	roomPos = -1
	for _, occupying := range room {
		if occupying == "." {
			roomPos++
			dist++
		}
	}
	return dist, roomPos, true
}

// Distance to hall position from room, can get there?
func pathToHall(letter string, to int, hall [11]string, room []string) (dist, roomPos int, ok bool) {
	fromTo := []int{roomDoor(letter), to}
	sort.Ints(fromTo)
	for _, occupying := range hall[fromTo[0]+1 : fromTo[1]+1] {
		if occupying != "." {
			return 0, 0, false
		}
	}
	dist = fromTo[1] - fromTo[0]
	dist++
	for _, occupying := range room {
		if occupying == "." {
			dist++
			roomPos++
		}
	}
	return dist, roomPos, true
}

// Index of room in rooms
func roomIndex(letter string) int {
	switch letter {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	case "D":
		return 3
	}
	panic("wrong room, can't get room index")
}

// Letter of room in rooms
func roomLetter(index int) string {
	switch index {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "D"
	}
	panic("wrong room, can't get room index")
}

// Should amphipods move from this room?
func shouldMove(letter string, room []string) bool {
	for _, occupying := range room {
		if occupying != letter && occupying != "." {
			return true
		}
	}
	return false
}

func findMinEnergy(init state) int64 {
	energyMap := map[string]int64{"A": 1, "B": 10, "C": 100, "D": 1000}
	memory := map[string]int64{}

	var recursiveFind func(current state) int64
	recursiveFind = func(current state) int64 {

		if everythingInItsRightPlace(current) {
			return 0
		}

		key := fmt.Sprintf("%v - %v", current.hall, current.rooms)

		if energy, found := memory[key]; found {
			return energy
		}

		for i, letter := range current.hall {
			if letter == "." {
				continue
			}
			roomIdx := roomIndex(letter)
			if roomAvailable(letter, current.rooms[roomIdx]) {
				if dist, roomPos, ok := path(letter, i, current.hall, current.rooms[roomIdx]); ok {
					energy := int64(dist) * energyMap[letter]
					newState := state{
						hall:  current.hall,
						rooms: [4][]string{{}, {}, {}, {}},
					}
					newState.hall[i] = "."
					for j, room := range current.rooms {
						newState.rooms[j] = append(newState.rooms[j], room...)
					}
					newState.rooms[roomIdx][roomPos] = letter
					// fmt.Println("moved to room using " + fmt.Sprint(energy) + " energy")
					// fmt.Println(current)
					// fmt.Println(newState)
					// fmt.Println(fmt.Printf("Moved from hall pos %d to room %s", i, letter))
					return energy + recursiveFind(newState)
				}

			}
		}

		minEnergy := int64(9999999999999)

		for i, room := range current.rooms {
			roomLet := roomLetter(i)
			if !shouldMove(roomLet, room) {
				continue
			}
			for j, occupying := range current.hall {
				if (j%2 == 0 && j != 0 && j != 10) || occupying != "." { // room entrances || occupied position
					continue
				}
				if dist, roomPos, ok := pathToHall(roomLet, j, current.hall, room); ok {
					newState := state{
						hall:  current.hall,
						rooms: [4][]string{{}, {}, {}, {}},
					}
					newState.hall[j] = room[roomPos]
					for k, room := range current.rooms {
						newState.rooms[k] = append(newState.rooms[k], room...)
					}
					newState.rooms[i][roomPos] = "."
					energy := energyMap[room[roomPos]] * int64(dist)
					// fmt.Println("moved to hall using " + fmt.Sprint(energy) + " energy")
					// fmt.Println(current)
					// fmt.Println(newState)
					minEnergy = maths.MinInt64(minEnergy, energy+recursiveFind(newState))
				}
			}
		}

		memory[key] = minEnergy

		return minEnergy

	}

	return recursiveFind(init)
}

func main() {
	start := time.Now()
	input := parseInput(files.ReadFile(23, "\r\n"))

	fmt.Printf("Data readed in %v \n\n", time.Since(start))
	fmt.Println(input)

	// Part 1
	start = time.Now()
	solution := findMinEnergy(input)
	fmt.Println(solution)
	fmt.Printf("Part 1 solved in %v \n\n", time.Since(start))

	// // Part 2
	start = time.Now()
	input.rooms[0] = []string{input.rooms[0][0], "D", "D", input.rooms[0][1]}
	input.rooms[1] = []string{input.rooms[1][0], "C", "B", input.rooms[1][1]}
	input.rooms[2] = []string{input.rooms[2][0], "B", "A", input.rooms[2][1]}
	input.rooms[3] = []string{input.rooms[3][0], "A", "C", input.rooms[3][1]}
	fmt.Println(input)
	solution = findMinEnergy(input)
	fmt.Println(solution)
	fmt.Printf("Part 2 solved in %v \n\n", time.Since(start))
}
