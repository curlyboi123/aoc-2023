package main

import (
	"aoc-2023/utils"
	"fmt"
	"regexp"
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
	nonStrongBoulderRe := regexp.MustCompile(`[\.O]+`)
	strongBoulderRe := regexp.MustCompile(`#+`)
	for _, line := range linesAsCols {
		fmt.Println("Unsorted line:", line)
		// Simulate rolling rocks to the north
		sortedLine := make([]string, 0, len(line))
		foo := strings.Join(line, "")
		boulderGroupings := nonStrongBoulderRe.FindAllStringIndex(foo, -1)
		heavyBoulderGroupings := strongBoulderRe.FindAllStringIndex(foo, -1)

		for _, grouping := range boulderGroupings {
			section := line[grouping[0]:grouping[1]]
			slices.SortStableFunc(section, func(a, b string) int {
				return rockMap[a] - rockMap[b]
			})
			sortedLine = append(sortedLine, section...)
			// slices.Insert(sortedLine, grouping[0], section...)
		}

		for _, bar := range heavyBoulderGroupings {
			fmt.Println("Heavy boulder at position: ", bar[0])
			heavyBoulderLength := bar[1] - bar[0]
			s := make([]string, heavyBoulderLength)
			for i := range s {
				s[i] = "#"
			}
			slices.Insert(sortedLine, bar[0], s...)
			fmt.Println(s)
		}

		// for _, bar := range heavyBoulderGroupings {
		// 	s := make([]string, bar[1]-bar[0])
		// 	for i := range s {
		// 		s[i] = "#"
		// 	}
		// 	slices.Insert(sortedLine, bar[0], s...)
		// }
		fmt.Println("Sorted line:  ", sortedLine)

		// Get value of round stones in positions
		lineLength := len(line)
		for idx, item := range line {
			if item == "O" {
				total += lineLength - idx
			}
		}
		fmt.Println()
	}
	fmt.Println(total)
}
