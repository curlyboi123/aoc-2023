package main

import (
	"aoc-2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func printRows(lines [][]string) {
	for idx, i := range lines {
		fmt.Println(idx, i)
	}
}

func sumRocks(lines [][]string) int {
	total := 0
	for _, line := range lines {
		lineLength := len(line)
		for idx, item := range line {
			if item == "O" {
				total += lineLength - idx
			}
		}
	}
	return total
}

func simulateTilt(lines [][]string, movableSections [][][]int, direction string) [][]string {
	rockMap := map[string]int{
		"O": 1,
		".": 2,
	}
	for idx, line := range lines {
		for _, section := range movableSections[idx] {
			slices.SortStableFunc(line[section[0]:section[1]], func(a, b string) int {
				if direction == "north" || direction == "west" {
					return rockMap[a] - rockMap[b]
				} else if direction == "south" || direction == "east" {
					return rockMap[b] - rockMap[a]
				} else {
					panic("Invalid direction supplied")
				}
			})
		}
	}
	return lines
}

// Lines should be in column format when passed in
func runCycle(lines [][]string) [][]string {
	re := regexp.MustCompile(`[\.O]+`)

	colSections := [][][]int{}
	for _, line := range lines {
		foo := strings.Join(line, "")
		colSections = append(colSections, re.FindAllStringIndex(foo, -1))
	}

	lines = utils.Transpose[string](lines) // Cols -> Rows
	rowSections := [][][]int{}
	for _, line := range lines {
		foo := strings.Join(line, "")
		rowSections = append(rowSections, re.FindAllStringIndex(foo, -1))
	}

	// North tilt
	lines = utils.Transpose[string](lines) // Rows -> Cols
	lines = simulateTilt(lines, colSections, "north")
	lines = utils.Transpose[string](lines) // Cols -> Rows

	// West tilt
	lines = simulateTilt(lines, rowSections, "west")

	// South tilt
	lines = utils.Transpose[string](lines) // Rows -> Cols
	lines = simulateTilt(lines, colSections, "south")
	lines = utils.Transpose[string](lines) // Cols -> Rows

	// East tilt
	lines = simulateTilt(lines, rowSections, "east")

	lines = utils.Transpose[string](lines) // Rows -> Cols
	return lines
}

func partOne() {
	linesFoo := utils.GetFileContentByLine()

	// Turn lines into a slice of slices initially as rows
	lines := [][]string{}
	for _, line := range linesFoo {
		bar := strings.Split(line, "")
		lines = append(lines, bar)
	}
	lines = utils.Transpose[string](lines) // Convert lines into cols

	re := regexp.MustCompile(`[\.O]+`)
	movableSections := [][][]int{}
	for _, line := range lines {
		foo := strings.Join(line, "")
		movableSections = append(movableSections, re.FindAllStringIndex(foo, -1))
	}
	lines = simulateTilt(lines, movableSections, "north")

	// Get value of round stones in positions
	total := sumRocks(lines)
	fmt.Println(total)
}

func partTwo() {
	linesFoo := utils.GetFileContentByLine()

	// Get lines as rows
	lines := [][]string{} // Rows
	for _, line := range linesFoo {
		bar := strings.Split(line, "")
		lines = append(lines, bar)
	}
	lines = utils.Transpose[string](lines) // Rows -> Cols

	// Get a map of the cycle totals and the cycle runs they correspond to
	cycleTotals := map[int][]int{}
	numOfCycles := 1000 // Arbitrary num of cycles to run to get all possible recurring totals
	for i := 0; i < numOfCycles; i++ {
		lines = runCycle(lines)
		cycleTotal := sumRocks(lines)
		cycleTotals[cycleTotal] = append(cycleTotals[cycleTotal], i+1)
	}

	// Find the number of cycles it takes to get the same rock sum twice.
	// This can differ so sometimes is the same between all values e.g. 10 but
	// could alternate diffs e.g. 2,8,2,8 but the diff between n and n+2 will
	// equal common recurring diff
	var recurringDiff int
	for _, v := range cycleTotals {
		if len(v) > 1 {
			diffs := []int{}
			for i := 0; i < len(v)-1; i++ {
				diff := v[i+1] - v[i]
				diffs = append(diffs, diff)
			}
			for i := 0; i < len(diffs)-2; i++ {
				if diffs[i] == diffs[i+1] {
					recurringDiff = diffs[i]
				}
			}
		}
	}

	// See for each rock sum sequence if target would be in it
	targetCycles := 1000000000
	var valueAtCycleEnd int
	for k, v := range cycleTotals {
		if (targetCycles-v[0])%recurringDiff == 0 && len(v) > 1 {
			valueAtCycleEnd = k
			break
		}
	}
	fmt.Println(valueAtCycleEnd)
}

func main() {
	partOne()
	partTwo()
}
