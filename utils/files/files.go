package files

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Reads content of the input file and returns it in an array, split by the specified delimiter
// If the input file does not exist, it will be created
func ReadFile(day int, delimiter string) []string {
	currentDay := strconv.Itoa(day)

	if len(currentDay) == 1 {
		currentDay = "0" + currentDay
	}

	filePath := fmt.Sprintf("calendar/%v/input.txt", currentDay)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

	if delimiter == "\n" {
		// there is a new line at the end of the input file, the last row is removed
		return slicedContent[:len(slicedContent) - 1]
	} else {
		// the last char is removed (the extra newline)
		lastElement := slicedContent[len(slicedContent) - 1]
		slicedContent[len(slicedContent) - 1] = lastElement[:len(lastElement) - 1]
		return slicedContent
	}
}