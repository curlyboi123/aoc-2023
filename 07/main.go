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

func main() {
	lines := getFileContentByLine()

	itemOrder := map[string]int{
		"2":  0,
		"3":  1,
		"4":  2,
		"5":  3,
		"6":  4,
		"7":  5,
		"8":  6,
		"9":  7,
		"10": 8,
		"J":  9,
		"Q":  10,
		"K":  11,
		"A":  12,
	}

	for _, line := range lines {
		hand, bid, _ := strings.Cut(line, " ")
		fmt.Printf("Hand: %s\n", hand)
		fmt.Printf("Bet: %s\n", bid)
		foo := strings.Split(hand, "")
		fmt.Println()
		for _, card := range foo {
			fmt.Println(itemOrder[string(card)])
		}
	}
}
