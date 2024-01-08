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

type Rank int64

const (
	HighCard Rank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func getHandRanking(hand string) int {
	// Count number of each different value card in hand
	cardCount := make(map[string]int)
	for _, char := range hand {
		if _, ok := cardCount[string(char)]; !ok {
			cardCount[string(char)] = strings.Count(hand, string(char))
		}
	}
	// Potential different approach could be taken by looping through cards
	// count and increase hand ranking as hands are discovered

	// Frequency of each number of card occurrences
	cardCountCount := make(map[int]int)
	for _, v := range cardCount {
		cardCountCount[v]++
	}

	// Assign hand rank
	if _, ok := cardCountCount[5]; ok {
		return int(FiveOfAKind)
	} else if _, ok := cardCountCount[4]; ok {
		return int(FourOfAKind)
	} else if _, ok := cardCountCount[3]; ok {
		if _, ok := cardCountCount[2]; ok {
			return int(FullHouse)
		} else {
			return int(ThreeOfAKind)
		}
	} else if val, ok := cardCountCount[2]; ok {
		if val == 2 {
			return int(TwoPair)
		} else if val == 1 {
			return int(OnePair)
		}
	} else {
		return int(HighCard)
	}
	return -1
}

func partOne() {
	lines := getFileContentByLine()

	cardRankings := map[string]int{
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

	handBidMap := map[string]int{}
	for _, line := range lines {
		hand, b, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(b)
		handBidMap[hand] = bid
	}

	// Ordered list to hold the hands grouped by rank e.g. FullHouse
	// Position in list is relative rank of hand type
	handCount := make([][]string, 7)

	for hand := range handBidMap {
		handRank := getHandRanking(hand)
		handCount[handRank] = append(handCount[handRank], hand)
	}

	// Order hands by rank
	for _, hands := range handCount {
		sort.Slice(hands, func(i, j int) bool {
			var cond bool
			for n := 0; n < len(hands[i]); n++ {
				if hands[i][n] == hands[j][n] {
					continue
				} else {
					cond = cardRankings[string(hands[i][n])] < cardRankings[string(hands[j][n])]
					break
				}
			}
			return cond
		})
	}

	// Calculate total
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

func main() {
	partOne()
}
