package main

import (
	"aoc-2023/utils"
	"fmt"
	"slices"
	"strings"
)

type pipe struct {
	val  string
	xPos int
	yPos int
}

func getInitialCoords(lines []string) (x int, y int) {
	for y, line := range lines {
		x := strings.Index(line, "S")
		if x != -1 {
			return x, y
		}
	}
	panic("Starting position not found")
}

// Returns the values of the pipes that are connected to the given pipe
func getStartingPipeVal(pipeGrid []string, curPipe pipe, connectionMap map[string][]string) string {
	connectedPipeLocations := []string{}
	// Check if north pipe is connected
	if curPipe.yPos > 0 {
		northPipe := string(pipeGrid[curPipe.yPos-1][curPipe.xPos])
		if slices.Contains(connectionMap[northPipe], "S") {
			connectedPipeLocations = append(connectedPipeLocations, "N")

		}
	}

	// Check if east pipe is connected
	if curPipe.xPos < len(pipeGrid[curPipe.yPos])-1 {
		eastPipe := string(pipeGrid[curPipe.yPos][curPipe.xPos+1])
		if slices.Contains(connectionMap[eastPipe], "W") {
			connectedPipeLocations = append(connectedPipeLocations, "E")
		}
	}

	// Check if south pipe is connected
	if curPipe.yPos < len(pipeGrid)-1 {
		southPipe := string(pipeGrid[curPipe.yPos+1][curPipe.xPos])
		if slices.Contains(connectionMap[southPipe], "N") {
			connectedPipeLocations = append(connectedPipeLocations, "S")
		}
	}

	// Check if west pipe is connected
	if curPipe.xPos > 0 {
		westPipe := string(pipeGrid[curPipe.yPos][curPipe.xPos-1])
		if slices.Contains(connectionMap[westPipe], "E") {
			connectedPipeLocations = append(connectedPipeLocations, "W")
		}
	}

	for k, v := range connectionMap {
		if slices.Equal(v, connectedPipeLocations) {
			return k
		}
	}
	panic("Pipe supplied is not a starting pipe")
}

func main() {
	lines := utils.GetFileContentByLine()
	connectionMap := map[string][]string{
		"|": {"N", "S"},
		"-": {"E", "W"},
		"L": {"N", "E"},
		"J": {"N", "W"},
		"7": {"S", "W"},
		"F": {"E", "S"},
		".": {},
	}

	directionMap := map[string]string{
		"N": "S",
		"E": "W",
		"S": "N",
		"W": "E",
	}

	x, y := getInitialCoords(lines)
	startingPipe := pipe{"S", x, y}

	startingPipeVal := getStartingPipeVal(lines, startingPipe, connectionMap)

	travelDirection := connectionMap[startingPipeVal][0]

	currentPipe := startingPipe

	stepsTaken := 0

	pipesInNetwork := []pipe{}

	for {
		pipesInNetwork = append(pipesInNetwork, currentPipe)
		stepsTaken++
		switch travelDirection {
		case "N":
			currentPipe.yPos--
		case "E":
			currentPipe.xPos++
		case "S":
			currentPipe.yPos++
		case "W":
			currentPipe.xPos--
		}
		connectedPipe := string(lines[currentPipe.yPos][currentPipe.xPos])
		for _, dir := range connectionMap[connectedPipe] {
			if directionMap[dir] != travelDirection {
				travelDirection = dir
				break
			}
		}
		currentPipe.val = connectedPipe
		if currentPipe.val == startingPipe.val {
			break
		}
	}
	fmt.Println(stepsTaken / 2)
	fmt.Println(pipesInNetwork)

}
