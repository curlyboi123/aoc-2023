package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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
		"2": 0,
		"3": 1,
		"4": 2,
		"5": 3,
		"6": 4,
		"7": 5,
		"8": 6,
		"9": 7,
		"T": 8,
		"J": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}

	handCount := make([][]string, 7)

	handBidMap := map[string]int{}
	for _, line := range lines {
		hand, b, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(b)
		handBidMap[hand] = bid
	}

	for hand := range handBidMap {
		// Count number of each card in hand
		cardCount := make(map[string]int)
		for _, char := range hand {
			if _, ok := cardCount[string(char)]; !ok {
				cardCount[string(char)] = strings.Count(hand, string(char))
			}
		}
		// Potential different approach could be taken by looping through cards
		// count and increase hand ranking as hands are discovered

		// Frequency of each number of occurrences
		cardCountCount := make(map[int]int)
		for _, v := range cardCount {
			cardCountCount[v]++
		}

		// Assign hand rank
		if _, ok := cardCountCount[5]; ok {
			handCount[6] = append(handCount[6], hand)
		} else if _, ok := cardCountCount[4]; ok {
			handCount[5] = append(handCount[5], hand)
		} else if _, ok := cardCountCount[3]; ok {
			if _, ok := cardCountCount[2]; ok {
				handCount[4] = append(handCount[4], hand)
			} else {
				handCount[3] = append(handCount[3], hand)
			}
		} else if val, ok := cardCountCount[2]; ok {
			if val == 2 {
				handCount[2] = append(handCount[2], hand)
			} else if val == 1 {
				handCount[1] = append(handCount[1], hand)
			}
		} else {
			handCount[0] = append(handCount[0], hand)
		}
	}

	// Rank hands
	for _, hands := range handCount {
		sort.Slice(hands, func(i, j int) bool {
			var cond bool
			for n := 0; n < len(hands[i]); n++ {
				if hands[i][n] == hands[j][n] {
					continue
				} else {
					cond = itemOrder[string(hands[i][n])] < itemOrder[string(hands[j][n])]
					break
				}
			}
			return cond
		})
	}

	total := 0
	rank := 1
	for _, handType := range handCount {
		for _, hand := range handType {
			total = total + (handBidMap[hand] * rank)
			rank++
		}
	}
	fmt.Println(total)
}
