package main

import (
	"aoc-2023/utils"
	"fmt"
	"strings"
)

func getInitialCoords(lines []string) (x int, y int) {
	for y, line := range lines {
		x := strings.Index(line, "S")
		if x != -1 {
			return x, y
		}
	}
	panic("Starting position not found")
}

func getDirToTravel(curPipeVal string, curDirection string) string {
	switch curDirection {
	case "N":
		switch curPipeVal {
		case "|":
			fmt.Println("S")
			return "S"
		case "L":
			fmt.Println("L")
			return "E"
		case "J":
			fmt.Println("J")
			return "W"
		}
	case "E":
		switch curPipeVal {
		case "-":
			fmt.Println("-")
			return "W"
		case "F":
			fmt.Println("F")
			return "S"
		case "L":
			fmt.Println("L")
			return "N"
		}
	case "S":
		switch curPipeVal {
		case "|":
			fmt.Println("|")
			return "N"
		case "F":
			fmt.Println("F")
			return "E"
		case "7":
			fmt.Println("7")
			return "W"
		}
	case "W":
		switch curPipeVal {
		case "-":
			fmt.Println("-")
			return "E"
		case "7":
			fmt.Println("7")
			return "S"
		case "J":
			fmt.Println("J")
			return "N"
		}
	}
	panic("Invalid direction supplied")
}

type pipe struct {
	val  string
	xPos int
	yPos int
}

func getInitialDirection(pipeGrid []string, startingPipe pipe) string {
	if startingPipe.yPos > 0 {
		northPipe := string(pipeGrid[startingPipe.yPos-1][startingPipe.xPos])
		switch northPipe {
		case "|":
			return "N"
		case "7":
			return "W"
		case "F":
			return "E"
		}
	}
	if startingPipe.yPos < len(pipeGrid) {
		southPipe := string(pipeGrid[startingPipe.yPos+1][startingPipe.xPos])
		switch southPipe {
		case "|":
			return "S"
		case "J":
			return "W"
		case "L":
			return "E"
		}
	}
	if startingPipe.xPos < len(pipeGrid[startingPipe.yPos]) {
		eastPipe := string(pipeGrid[startingPipe.yPos+1][startingPipe.xPos+1])
		switch eastPipe {
		case "-":
			return "E"
		case "7":
			return "S"
		case "J":
			return "N"
		}
	}
	if startingPipe.xPos > 0 {
		westPipe := string(pipeGrid[startingPipe.yPos+1][startingPipe.xPos-1])
		switch westPipe {
		case "-":
			return "W"
		case "F":
			return "S"
		case "L":
			return "N"
		}
	}
	panic("No pipes connected to starting pipe")
}

func getNewPipe(pipeGrid []string, pipe pipe, travelDir string) pipe {
	switch travelDir {
	case "N":
		pipe.yPos = pipe.yPos - 1
	case "E":
		pipe.xPos = pipe.xPos + 1
	case "S":
		pipe.yPos = pipe.yPos + 1
	case "W":
		pipe.xPos = pipe.xPos - 1
	}
	fmt.Println(pipe.xPos, pipe.yPos)
	newPipeVal := string(pipeGrid[pipe.yPos][pipe.xPos])
	pipe.val = newPipeVal
	return pipe
}

// ///
func getConnectedPipes(pipeGrid []string, pipe pipe) []string {
	connectionMap := map[string][]string{
		"|": {"N", "S"},
		"-": {"E", "W"},
		"L": {"N", "E"},
		"J": {"N", "W"},
		"7": {"S", "W"},
		"F": {"E", "S"},
		".": {},
	}
	fmt.Println(connectionMap)
	// Check if north pipe is connected
	if pipe.yPos > 0 {
		northPipe := pipeGrid[pipe.yPos-1][pipe.xPos]
	}

	// Check if east pipe is connected

	// Check if south pipe is connected

	// Check if west pipe is connected
	return []string{}
}

func checkPipesConnected(pipeA pipe, pipeB pipe) bool {

	return true
}

func main() {
	lines := utils.GetFileContentByLine()

	x, y := getInitialCoords(lines)
	fmt.Printf("Start located at coords: (%d, %d)\n", x, y)

	startingPipeVal := string(lines[y][x])
	startingPipe := pipe{startingPipeVal, x, y}
	fmt.Printf("Starting pipe val: %s\n", startingPipe.val)

	currentTravelDir := getInitialDirection(lines, startingPipe)
	fmt.Printf("Current travel dir: %s\n", currentTravelDir)
	distanceTraveled := 0

	currentPipe := getNewPipe(lines, startingPipe, currentTravelDir)

	for currentPipe.val != startingPipe.val {
		fmt.Printf("Current pipe val: %s\n", currentPipe.val)

		currentTravelDir = getDirToTravel(currentPipe.val, currentTravelDir)
		currentPipe = getNewPipe(lines, currentPipe, currentTravelDir)
		distanceTraveled++

		fmt.Printf("Next pipe val: %s\n", currentPipe.val)
	}

	// Get connected pipes to start pipe
}
