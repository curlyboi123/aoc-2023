package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func splitMirrors() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n\n")
}

func transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func getPerfectLinesOfSymmetry(lines [][]string) [][2]int {
	// Get duplicate lines
	duplicateLines := map[string][]int{}
	visitedLines := []int{}
	for i := 0; i < len(lines); i++ {
		// Skip check if already in map
		if slices.Contains(visitedLines, i) {
			continue
		}
		for j := i + 1; j < len(lines); j++ {
			// Skip check if already in map
			if slices.Contains(visitedLines, j) {
				continue
			}

			if slices.Equal(lines[i], lines[j]) {
				duplicateLines[strings.Join(lines[i], "")] = append(duplicateLines[strings.Join(lines[i], "")], i, j)
			}
		}
	}

	// Find lines where there is line of symmetry
	linesOfSymmetry := [][2]int{}
	for _, v := range duplicateLines {
		for i := 0; i < len(v)-1; i++ {
			if v[i] == v[i+1]-1 {
				linesOfSymmetry = append(linesOfSymmetry, [2]int{v[i], v[i+1]})

			}
		}
	}

	// For each line of symmetry see how many lines it extends to
	linesOfPerfectSym := [][2]int{}
	for _, los := range linesOfSymmetry {
		curSymPair := los
		for {
			if curSymPair[0] == 0 || curSymPair[1] >= len(lines)-1 {
				linesOfPerfectSym = append(linesOfPerfectSym, los)
				break
			}
			prevLine := lines[curSymPair[0]-1]
			nextLine := lines[curSymPair[1]+1]
			// Part 2 needs to allow 1 element in slice to be different and still be considered equal
			if !slices.Equal(prevLine, nextLine) {
				break
			}
			curSymPair = [2]int{curSymPair[0] - 1, curSymPair[1] + 1}
		}
	}
	return linesOfPerfectSym
}

func main() {
	mirrors := splitMirrors()

	total := 0
	for idx, mirror := range mirrors {
		fmt.Println("Mirror: ", idx+1)
		lines := strings.Split(mirror, "\n")

		linesAsRows := [][]string{}
		for _, line := range lines {
			bar := strings.Split(line, "")
			linesAsRows = append(linesAsRows, bar)
		}
		linesAsCols := transpose[string](linesAsRows)

		rowLinesOfSym := getPerfectLinesOfSymmetry(linesAsRows)
		colLinesOfSym := getPerfectLinesOfSymmetry(linesAsCols)

		if len(rowLinesOfSym) > 0 {
			fmt.Println("Horizontal line of symmetry between lines: ", rowLinesOfSym)
		}
		if len(colLinesOfSym) > 0 {
			fmt.Println("Column line of symmetry between lines: ", colLinesOfSym)
		}

		for _, rowLines := range rowLinesOfSym {
			total += (rowLines[0] + 1) * 100
		}
		for _, colLines := range colLinesOfSym {
			total += colLines[0] + 1
		}
	}
	fmt.Println(total)
}
