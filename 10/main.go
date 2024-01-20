package main

import (
	"aoc-2023/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

type pipe struct {
	val  string
	xPos int
	yPos int
}

func getInitialCoords(lines []string) [2]int {
	for y, line := range lines {
		x := strings.Index(line, "S")
		if x != -1 {
			return [2]int{x, y}
		}
	}
	panic("Starting position not found")
}

func shoestringAlg(vertices [][2]int) float64 {
	xySum := 0
	yxSum := 0
	for i := 0; i < len(vertices)-1; i++ {
		// Handle start pipe at end of slice
		xySum += vertices[i][0] * vertices[i+1][1]
		yxSum += vertices[i][1] * vertices[i+1][0]
	}

	xySum += vertices[len(vertices)-1][0] * vertices[0][1]
	yxSum += vertices[len(vertices)-1][1] * vertices[0][0]

	total := math.Abs(float64(xySum-yxSum)) / 2
	return total
}

// Returns the values of the pipes that are connected to the given pipe
func getConnectedPipeLocations(pipeGrid []string, connectionMap map[string]map[string]string, directionMap map[string][2]int, pipeCoords [2]int) [][2]int {
	var connectedPipeLocations [][2]int

	// Check if north pipe is connected
	if pipeCoords[1] < len(pipeGrid)-1 {
		northPipeCoords := [2]int{pipeCoords[0] + directionMap["N"][0], pipeCoords[1] + directionMap["N"][1]}
		northPipeVal := string(pipeGrid[northPipeCoords[1]][northPipeCoords[0]])
		if _, ok := connectionMap[northPipeVal]["S"]; ok {
			connectedPipeLocations = append(connectedPipeLocations, northPipeCoords)
		}
	}

	// Check if east pipe is connected
	if pipeCoords[0] < len(pipeGrid[0])-1 {
		eastPipeCoords := [2]int{pipeCoords[0] + directionMap["E"][0], pipeCoords[1] + directionMap["E"][1]}
		eastPipeVal := string(pipeGrid[eastPipeCoords[1]][eastPipeCoords[0]])
		if _, ok := connectionMap[eastPipeVal]["W"]; ok {
			connectedPipeLocations = append(connectedPipeLocations, eastPipeCoords)
		}
	}

	// Check if south pipe is connected
	if pipeCoords[1] > 0 {
		southPipeCoords := [2]int{pipeCoords[0] + directionMap["S"][0], pipeCoords[1] + directionMap["S"][1]}
		southPipeVal := string(pipeGrid[southPipeCoords[1]][southPipeCoords[0]])
		if _, ok := connectionMap[southPipeVal]["N"]; ok {
			connectedPipeLocations = append(connectedPipeLocations, southPipeCoords)
		}
	}

	// Check if west pipe is connected
	if pipeCoords[0] > 0 {
		westPipeCoords := [2]int{pipeCoords[0] + directionMap["W"][0], pipeCoords[1] + directionMap["W"][1]}
		westPipeVal := string(pipeGrid[westPipeCoords[1]][westPipeCoords[0]])
		if _, ok := connectionMap[westPipeVal]["E"]; ok {
			connectedPipeLocations = append(connectedPipeLocations, westPipeCoords)
		}
	}

	return connectedPipeLocations
}

func getNextConnectedPipe(pipeGrid []string, connectionMap map[string]map[string]string, directionMap map[string][2]int, curPipeCoords [2]int, entryDir string) [2]int {
	curVal := string(pipeGrid[curPipeCoords[1]][curPipeCoords[0]])
	dirToTravel := connectionMap[curVal][entryDir]
	nextPipeCoords := [2]int{curPipeCoords[0] + directionMap[dirToTravel][0], curPipeCoords[1] + directionMap[dirToTravel][1]}
	return nextPipeCoords
}

func getEntryDirection(pipeACoords [2]int, pipeBCoords [2]int, directionMap map[string][2]int) string {
	diff := [2]int{pipeACoords[0] - pipeBCoords[0], pipeACoords[1] - pipeBCoords[1]}
	for k, v := range directionMap {
		if v == diff {
			return k
		}
	}
	panic("Pipes supplied are not connected")
}

func main() {
	pipeGrid := utils.GetFileContentByLine()
	slices.Reverse(pipeGrid)
	connectionMap := map[string]map[string]string{
		"|": {"N": "S", "S": "N"},
		"-": {"E": "W", "W": "E"},
		"L": {"N": "E", "E": "N"},
		"J": {"N": "W", "W": "N"},
		"7": {"S": "W", "W": "S"},
		"F": {"E": "S", "S": "E"},
		".": {},
	}
	directionMap := map[string][2]int{
		"N": {0, 1},
		"E": {1, 0},
		"S": {0, -1},
		"W": {-1, 0},
	}

	// Coords {0, 0} is bottom left point of grid
	var pipesInNetwork [][2]int

	startingPipeCoords := getInitialCoords(pipeGrid)

	startingPipeConnectedPipes := getConnectedPipeLocations(pipeGrid, connectionMap, directionMap, startingPipeCoords)

	curDir := getEntryDirection(startingPipeCoords, startingPipeConnectedPipes[0], directionMap)

	curPipeCoords := startingPipeConnectedPipes[0]

	pipesInNetwork = append(pipesInNetwork, startingPipeCoords, curPipeCoords)

	for curPipeCoords != startingPipeCoords {
		newPipeCoords := getNextConnectedPipe(pipeGrid, connectionMap, directionMap, curPipeCoords, curDir)
		newDir := getEntryDirection(curPipeCoords, newPipeCoords, directionMap)

		pipesInNetwork = append(pipesInNetwork, newPipeCoords)

		curPipeCoords = newPipeCoords
		curDir = newDir
	}
	fmt.Println("Pipes in network: ", pipesInNetwork)
	fmt.Println(len(pipesInNetwork) / 2)

	var vertices [][2]int
	vertexTypes := []string{"L", "J", "7", "F"}
	for _, pipe := range pipesInNetwork {
		pipeVal := string(pipeGrid[pipe[1]][pipe[0]])
		if slices.Contains(vertexTypes, pipeVal) {
			vertices = append(vertices, pipe)
		}
	}
	fmt.Println(vertices)
	foo := shoestringAlg(vertices)
	fmt.Println(foo)
}
