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

func main() {
	mirrors := splitMirrors()

	for idx, mirror := range mirrors {

		fmt.Println("Mirror: ", idx)
		horizontalLines := strings.Split(mirror, "\n")

		// Get duplicate lines
		duplicateLines := map[string][]int{}
		visitedLines := []int{}
		for i := 0; i < len(horizontalLines); i++ {
			// Skip check if already in map
			if slices.Contains(visitedLines, i) {
				continue
			}
			for j := i + 1; j < len(horizontalLines); j++ {
				// Skip check if already in map
				if slices.Contains(visitedLines, j) {
					continue
				}
				if horizontalLines[i] == horizontalLines[j] {
					duplicateLines[horizontalLines[i]] = append(duplicateLines[horizontalLines[i]], i, j)
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
		fmt.Println("Lines of symmetry: ", linesOfSymmetry)

		// For each line of symmetry see how many lines it extends to
		for _, los := range linesOfSymmetry {
			curSymPair := los
			numLinesInLineOfSym := map[[2]int]int{los: 1}
			for {
				if curSymPair[0] == 0 || curSymPair[1] >= len(horizontalLines)-1 {
					break
				}
				prevLine := horizontalLines[curSymPair[0]-1]
				nextLine := horizontalLines[curSymPair[1]+1]
				if prevLine != nextLine {
					break
				}
				curSymPair = [2]int{curSymPair[0] - 1, curSymPair[1] + 1}
				numLinesInLineOfSym[los]++
			}
			fmt.Println(numLinesInLineOfSym)
		}

		// Get largest amount of lines in line of symmetry as that is the perfect one
	}
}
