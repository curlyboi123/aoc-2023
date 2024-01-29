package main

import (
	"aoc-2023/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
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
	for _, line := range linesAsCols {
		// Count continuous unbroken part
		sortedLine := []string{}
		currentSection := []string{}

		for idx, item := range line {
			if item == "#" {
				// Sort section
				slices.SortStableFunc(currentSection, func(a, b string) int {
					return rockMap[a] - rockMap[b]
				})
				sortedLine = append(sortedLine, currentSection...)
				sortedLine = append(sortedLine, "#")
				currentSection = nil
				continue
			}
			if idx == len(line)-1 {
				// Sort section
				currentSection = append(currentSection, item)
				slices.SortStableFunc(currentSection, func(a, b string) int {
					return rockMap[a] - rockMap[b]
				})
				sortedLine = append(sortedLine, currentSection...)
				currentSection = nil
				continue
			}
			currentSection = append(currentSection, item)
		}

		// Get value of round stones in positions
		lineLength := len(line)
		for idx, item := range sortedLine {
			if item == "O" {
				total += lineLength - idx
			}
		}
	}
	fmt.Println(total)
}
