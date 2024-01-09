package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
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

type Container struct {
	mu           sync.Mutex
	finishedNums map[int][]string // map of the number of moves it took a starting node to reach a finish position
}

// Return greatest common divisor using Euclid's Algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Return lowest common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Return lcm of args
func LCMM(integers ...int) int {
	result := integers[0]
	for i := 1; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func partTwo() {
	lines := getFileContentByLine()
	network := createNetworkMap(lines[2:])
	moves := lines[0]

	finishedNums := make(map[int][]string)
	c := Container{
		finishedNums: finishedNums,
	}

	// Get the starting nodes
	startingNodes := []string{}
	for currentNode := range network {
		lastCharacter := currentNode[len(currentNode)-1]
		if string(lastCharacter) == "A" {
			startingNodes = append(startingNodes, currentNode)
		}
	}

	var wg sync.WaitGroup
	allNodesFinished := false
	for !allNodesFinished {
		for _, node := range startingNodes {
			wg.Add(1)
			n := node
			go func() {
				// Run check
				defer wg.Done()
				numberOfMoves := getNumMovesToFinishNode(n, network, moves, 2)

				c.mu.Lock()
				defer c.mu.Unlock()

				// If amount to finish hasn't already been counted then append
				if !slices.Contains(c.finishedNums[numberOfMoves], n) {
					c.finishedNums[numberOfMoves] = append(c.finishedNums[numberOfMoves], n)
				}

				if len(c.finishedNums) == len(startingNodes) {
					allNodesFinished = true
					return
				}
			}()
		}
		wg.Wait()
	}

	var movesTaken []int
	for k := range c.finishedNums {
		movesTaken = append(movesTaken, k)
	}
	totalMovesTaken := LCMM(movesTaken...)
	fmt.Println(totalMovesTaken)
}

func main() {
	partOne()
	partTwo()
}
