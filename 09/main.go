package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func checkContainsNonZero(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	lines := getFileContentByLine()

	allHistoryTotal := 0
	for _, line := range lines {
		initialVals := strings.Fields(line)
		vals := []int{}
		for _, val := range initialVals {
			v, _ := strconv.Atoi(val)
			vals = append(vals, v)
		}

		allDiffs := [][]int{} // Holds all lines after passthrough
		allDiffs = append(allDiffs, vals)

		// Get all histories until one containing all zeroes is reached
		for {
			currentDiffs := []int{}
			for i := 0; i < len(vals)-1; i++ {
				curVal := vals[i]
				nextVal := vals[i+1]
				diff := nextVal - curVal
				currentDiffs = append(currentDiffs, diff)
			}
			allDiffs = append(allDiffs, currentDiffs)
			vals = currentDiffs
			if checkContainsNonZero(currentDiffs) {
				break
			}
		}

		historyTotal := 0
		slices.Reverse(allDiffs) // Reverse to make loop easier
		for i := 0; i < len(allDiffs); i++ {
			lastVal := allDiffs[i][len(allDiffs[i])-1]
			historyTotal = historyTotal + lastVal
		}
		allHistoryTotal = allHistoryTotal + historyTotal
	}
	fmt.Println(allHistoryTotal)
}
