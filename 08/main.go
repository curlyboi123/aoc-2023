package main

import (
	"fmt"
	"os"
	"strings"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func createNetworkMap(nodes []string) map[string][]string {
	network := make(map[string][]string)
	for _, node := range nodes {
		current, moves, _ := strings.Cut(node, "=")
		current = strings.TrimSpace(current)

		left, right, _ := strings.Cut(strings.Trim(strings.TrimSpace(moves), "()"), ",")
		left = strings.TrimSpace(left)
		right = strings.TrimSpace(right)
		network[current] = append(network[current], left, right)
	}
	return network
}

func isNodeFinished(node string, part int) bool {
	if part == 1 {
		return node == "ZZZ"
	} else if part == 2 {
		return string(node[len(node)-1]) == "Z"
	} else {
		panic("Invalid part number supplied")
	}
}

func getNumMovesToFinishNode(startingNode string, network map[string][]string, moves string, part int) int {

	numOfMoves := 0
	nodeFinished := false
	currentNode := startingNode
	for !nodeFinished {
		for _, move := range moves {
			numOfMoves++
			if string(move) == "L" {
				currentNode = network[currentNode][0]
			} else {
				currentNode = network[currentNode][1]
			}
			// Could calculate what part to use for finished logic instead of passthrough
			nodeFinished = isNodeFinished(currentNode, part)
			if nodeFinished {
				break
			}
		}
	}
	return numOfMoves
}

func partOne() {
	lines := getFileContentByLine()
	network := createNetworkMap(lines[2:])
	moves := lines[0]
	startingNode := "AAA"
	if _, ok := network[startingNode]; !ok {
		fmt.Println("Input unusable for part one")
		return
	}
	numberOfMoves := getNumMovesToFinishNode(startingNode, network, moves, 1)
	fmt.Println(numberOfMoves)
}

func partTwo() {
	lines := getFileContentByLine()
	network := createNetworkMap(lines[2:])
	movements := lines[0]
	fmt.Println(movements)

	// Get the starting nodes
	currentNodes := []string{}
	for currentNode := range network {
		lastCharacter := currentNode[len(currentNode)-1]
		if string(lastCharacter) == "A" {
			currentNodes = append(currentNodes, currentNode)
		}
	}
	fmt.Println(currentNodes)

	// Go routine for each starting node
	// When finish found check other nodes have same value finish

	numberOfMoves := 0
	nodesFinished := false
	for !nodesFinished {
		for _, move := range movements {
			numberOfMoves++
			numberFinishedNodes := 0
			for idx, node := range currentNodes {
				var newNodeVal string
				if string(move) == "L" {
					newNodeVal = network[node][0]
				} else {
					newNodeVal = network[node][1]
				}
				currentNodes[idx] = newNodeVal
				if isNodeFinished(newNodeVal, 2) {
					numberFinishedNodes++
				}
			}
			if numberFinishedNodes > 0 {
				fmt.Println("Node Finished")
				fmt.Println(currentNodes)
			}
			if numberFinishedNodes == len(currentNodes) {
				nodesFinished = true
				break
			}

		}
	}
	fmt.Println(numberOfMoves)
}

func main() {
	partOne()
	// partTwo()
}
