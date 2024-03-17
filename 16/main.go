package main

import (
	"fmt"
	"slices"
	"sync"

	"example.com/utils"
)

func getIndexOfNextMirror(line []string) int {
	return slices.IndexFunc(line, func(e string) bool {
		return e != "."
	})
}

type coords struct {
	row, col int
}

func getNewDirections(curDir, nextMirror string) []string {
	mirrorMap := map[string]map[string][]string{
		"|": {
			"N": []string{"N"},
			"S": []string{"S"},
			"E": []string{"S", "N"},
			"W": []string{"S", "N"},
		},
		"-": {
			"N": []string{"E", "W"},
			"S": []string{"E", "W"},
			"E": []string{"E"},
			"W": []string{"W"},
		},
		"/": {
			"N": []string{"E"},
			"S": []string{"W"},
			"E": []string{"N"},
			"W": []string{"S"},
		},
		"\\": {
			"N": []string{"W"},
			"S": []string{"E"},
			"E": []string{"S"},
			"W": []string{"N"},
		},
	}
	return mirrorMap[nextMirror][curDir]
}

type Container struct {
	mu         sync.Mutex
	coordPairs map[[2]coords]bool
}

func (c *Container) addPair(pair [2]coords) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.coordPairs[pair] = true
}

func getNextMirrorCoords(gridAsRows [][]string, gridAsCols [][]string, currentDir string, currentCoords coords) coords {
	switch currentDir {
	case "N":
		fmt.Println("N")
		sectionToCheck := gridAsCols[currentCoords.col][:currentCoords.row]
		nextMirrorIdx := getIndexOfNextMirror(sectionToCheck)
		newCoords := coords{nextMirrorIdx, currentCoords.col}
		return newCoords
	case "E":
		sectionToCheck := gridAsRows[currentCoords.row][currentCoords.col+1:]
		nextMirrorIdx := getIndexOfNextMirror(sectionToCheck)
		newCoords := coords{currentCoords.row, currentCoords.col + nextMirrorIdx + 1}
		return newCoords
	case "S":
		fmt.Println("S")
		sectionToCheck := gridAsCols[currentCoords.col][currentCoords.row+1:]
		nextMirrorIdx := getIndexOfNextMirror(sectionToCheck)
		newCoords := coords{currentCoords.row + nextMirrorIdx + 1, currentCoords.col}
		return newCoords
	case "W":
		fmt.Println("W")
		sectionToCheck := gridAsRows[currentCoords.row][:currentCoords.col]
		nextMirrorIdx := getIndexOfNextMirror(sectionToCheck)
		newCoords := coords{currentCoords.row, nextMirrorIdx}
		return newCoords
	}
	panic("Invalid direction supplied")
}

func main() {
	lines := utils.GetFileContentByLine()
	gridAsRows := utils.Get2dGridFromLines(lines)
	gridAsCols := utils.Transpose(gridAsRows)

	c := Container{coordPairs: map[[2]coords]bool{}}
	fmt.Println(c.coordPairs)

	currentCoords := coords{0, 0}
	currentDir := "E"

	fmt.Println("Starting at position: ", currentCoords)
	fmt.Println("Traveling: ", currentDir)

	// Current coords
	// Current direction
	// Return new coords where next mirror encountered

	// Take current direction
	// Take new mirror encountered
	// Return new direction(s)

	// Take current coords
	// Take end coords
	// Add to slice tracking start and end coord pairings

	var wg sync.WaitGroup

	var runCycle func(dir string, curCoords coords)

	runCycle = func(dir string, curCoords coords) {
		wg.Add(1)
		nextMirrorCoords := getNextMirrorCoords(gridAsRows, gridAsCols, dir, curCoords)
		nextMirror := string(gridAsRows[nextMirrorCoords.col][nextMirrorCoords.row])
		nextDirections := getNewDirections(dir, nextMirror)
		newPair := [2]coords{curCoords, nextMirrorCoords}
		if _, ok := c.coordPairs[newPair]; ok {
			fmt.Println("Pair already in map")
			fmt.Println(newPair)
			wg.Done()
			return
		}
		c.addPair(newPair)
		for _, dir := range nextDirections {
			runCycle(dir, nextMirrorCoords)
		}
	}
	runCycle(currentDir, currentCoords)

	// Get next direction(s) for light to travel in
	// Start pos, dir to travel -> end pos
	// If splitter then start multiple concurrent new beams

	// Var to hold all beam stretches
	beamSections := [][]int{}
	fmt.Println(beamSections)
}
