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

func getLineOfSymmetryIndexPartOne(lines [][]string) int {
	for i := 0; i < len(lines)-1; i++ {
		symmetryPossible := true
		for a, b := i, i+1; a >= 0 && b < len(lines); a, b = a-1, b+1 {
			if !slices.Equal(lines[a], lines[b]) {
				symmetryPossible = false
				break
			}
		}
		if symmetryPossible {
			return i + 1
		}
	}
	return 0
}

func getLineOfSymmetryIndexPartTwo(lines [][]string) int {
	for i := 0; i < len(lines)-1; i++ {
		// Track how many tiles changed
		// We are looking for solution that changed 1 tile
		tilesChanged := 0
		symmetryPossible := true
		for a, b := i, i+1; a >= 0 && b < len(lines); a, b = a-1, b+1 {
			if slices.CompareFunc(lines[a], lines[b], func(e1, e2 string) int {
				if e1 != e2 {
					if tilesChanged == 0 {
						tilesChanged++
						return 0
					} else {
						return -1
					}
				}
				return 0
			}) != 0 {
				symmetryPossible = false
				break
			}
		}
		if symmetryPossible && tilesChanged == 1 {
			return i + 1
		}
	}
	return 0
}

func partOne() {
	mirrors := splitMirrors()
	total := 0
	for _, mirror := range mirrors {
		lines := strings.Split(mirror, "\n")

		linesAsRows := [][]string{}
		for _, line := range lines {
			bar := strings.Split(line, "")
			linesAsRows = append(linesAsRows, bar)
		}
		linesAsCols := transpose[string](linesAsRows)

		rowNumLineOfSym := getLineOfSymmetryIndexPartOne(linesAsRows)
		colNumLineOfSym := getLineOfSymmetryIndexPartOne(linesAsCols)

		total += rowNumLineOfSym * 100
		total += colNumLineOfSym

	}
	fmt.Println(total)
}

func partTwo() {
	mirrors := splitMirrors()
	total := 0
	for _, mirror := range mirrors {
		lines := strings.Split(mirror, "\n")

		linesAsRows := [][]string{}
		for _, line := range lines {
			bar := strings.Split(line, "")
			linesAsRows = append(linesAsRows, bar)
		}
		linesAsCols := transpose[string](linesAsRows)

		rowNumLineOfSym := getLineOfSymmetryIndexPartTwo(linesAsRows)
		colNumLineOfSym := getLineOfSymmetryIndexPartTwo(linesAsCols)

		total += rowNumLineOfSym * 100
		total += colNumLineOfSym

	}
	fmt.Println(total)
}

func main() {
	partOne()
	partTwo()
}
