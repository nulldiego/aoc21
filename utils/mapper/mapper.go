package mapper

import (
	"strconv"
	"strings"
)

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic(err)
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}

func ToIntMatrix(slice []string, delimiter string) [][]int {
	sliceToReturn := make([][]int, len(slice))
	for i, line := range slice {
		sliceToReturn[i] = ToIntSlice(strings.Split(line, delimiter))
	}
	return sliceToReturn
}
