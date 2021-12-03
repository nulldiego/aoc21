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
	return strings.Split(fileContent, delimiter)
}
