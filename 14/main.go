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

func partOne() {
	rockMap := map[string]int{
		"O": 1,
		".": 2,
	}
	lines := utils.GetFileContentByLine()

	// Convert lines into cols
	linesAsRows := [][]string{}
	for _, line := range lines {
		bar := strings.Split(line, "")
		linesAsRows = append(linesAsRows, bar)
	}
	linesAsCols := utils.Transpose[string](linesAsRows)

	total := 0
	re := regexp.MustCompile(`[\.O]+`)
	for _, line := range linesAsCols {
		foo := strings.Join(line, "")
		sections := re.FindAllStringIndex(foo, -1)
		for _, section := range sections {
			slices.SortStableFunc(line[section[0]:section[1]], func(a, b string) int {
				return rockMap[a] - rockMap[b]
			})

		}
		// Get value of round stones in positions
		lineLength := len(line)
		for idx, item := range line {
			if item == "O" {
				total += lineLength - idx
			}
		}
	}
	fmt.Println(total)
}

func runCycle(lines [][]string) [][]string {
	rockMap := map[string]int{
		"O": 1,
		".": 2,
	}
	re := regexp.MustCompile(`[\.O]+`)
	currentLines := utils.Transpose[string](lines)
	// fmt.Println("Initial lines")
	// printRows(currentLines)
	// fmt.Println()

	// North tilt
	for _, line := range currentLines {
		foo := strings.Join(line, "")
		sections := re.FindAllStringIndex(foo, -1)
		for _, section := range sections {
			slices.SortStableFunc(line[section[0]:section[1]], func(a, b string) int {
				return rockMap[a] - rockMap[b]
			})

		}
	}
	currentLines = utils.Transpose[string](currentLines)
	// fmt.Println("Lines after north tilt")
	// printRows(currentLines)
	// fmt.Println()

	// West tilt
	for _, line := range currentLines {
		foo := strings.Join(line, "")
		sections := re.FindAllStringIndex(foo, -1)
		for _, section := range sections {
			slices.SortStableFunc(line[section[0]:section[1]], func(a, b string) int {
				return rockMap[a] - rockMap[b]
			})

		}
	}
	// fmt.Println("Lines after west tilt")
	// printRows(currentLines)
	// fmt.Println()

	// South tilt
	currentLines = utils.Transpose[string](currentLines)

	for _, line := range currentLines {
		foo := strings.Join(line, "")
		sections := re.FindAllStringIndex(foo, -1)
		for _, section := range sections {
			slices.SortStableFunc(line[section[0]:section[1]], func(a, b string) int {
				return rockMap[b] - rockMap[a]
			})

		}
	}
	currentLines = utils.Transpose[string](currentLines)

	// fmt.Println("Lines after south tilt")
	// printRows(currentLines)
	// fmt.Println()

	// East tilt
	for _, line := range currentLines {
		foo := strings.Join(line, "")
		sections := re.FindAllStringIndex(foo, -1)
		for _, section := range sections {
			slices.SortStableFunc(line[section[0]:section[1]], func(a, b string) int {
				return rockMap[b] - rockMap[a]
			})

		}
	}
	// fmt.Println("Lines after east tilt")
	// printRows(currentLines)
	// fmt.Println()
	return currentLines
}

func partTwo() {

	lines := utils.GetFileContentByLine()

	// Convert lines into cols
	linesAsRows := [][]string{}
	for _, line := range lines {
		bar := strings.Split(line, "")
		linesAsRows = append(linesAsRows, bar)
	}

	linesFoo := linesAsRows

	numOfCycles := 1000000000
	for i := 0; i < numOfCycles; i++ {
		linesFoo = runCycle(linesFoo)
		fmt.Println("Cycle ", i+1)
		// printRows(linesFoo)

	}

	total := 0
	for _, line := range linesFoo {
		lineLength := len(line)
		for idx, item := range line {
			if item == "O" {
				total += lineLength - idx
			}
		}
	}

}

func main() {
	// partOne()
	partTwo()
}
