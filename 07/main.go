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
	cardCountCount := map[int]int{
		1: 0, 2: 0, 3: 0, 4: 0, 5: 0,
	}
	for _, v := range cardCount {
		cardCountCount[v]++
	}

	// Assign hand rank
	if cardCountCount[5] > 0 {
		return int(FiveOfAKind)
	} else if cardCountCount[4] > 0 {
		return int(FourOfAKind)
	} else if cardCountCount[3] > 0 && cardCountCount[2] > 0 {
		return int(FullHouse)
	} else if cardCountCount[3] > 0 {
		return int(ThreeOfAKind)
	} else if cardCountCount[2] == 2 {
		return int(TwoPair)
	} else if cardCountCount[2] == 1 {
		return int(OnePair)
	} else {
		return int(HighCard)
	}
}

func getAdjustedHandRanking(handRank int, amountToIncrease int) int {
	if handRank == int(FourOfAKind) {
		return handRank + 1
	} else if handRank == int(ThreeOfAKind) {
		return handRank + amountToIncrease + 1
	} else if handRank == int(TwoPair) {
		return handRank + 2
	} else if handRank == int(OnePair) {
		if amountToIncrease == 2 || amountToIncrease == 3 {
			return handRank + amountToIncrease + 2
		} else if amountToIncrease == 1 {
			return handRank + amountToIncrease + 1
		} else {
			return handRank
		}
	} else {
		if amountToIncrease == 3 || amountToIncrease == 4 {
			return handRank + amountToIncrease + 2
		} else if amountToIncrease == 2 {
			return handRank + amountToIncrease + 1
		} else {
			return handRank + amountToIncrease
		}
	}
}

type player struct {
	hand string
	bid  int
}

func partOne() {
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
	lines := getFileContentByLine()

	// Slice to hold slices of hands grouped by rank where order is relative
	// ranking of hand types
	handCount := make([][]player, 7)
	for _, line := range lines {
		hand, b, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(b)
		p := player{hand: hand, bid: bid}
		handRank := getHandRanking(hand)
		handCount[handRank] = append(handCount[handRank], p)
	}

	// Order all hands by rank
	for _, hands := range handCount {
		sort.Slice(hands, func(i, j int) bool {
			var cond bool
			for n := 0; n < len(hands[i].hand); n++ {
				if hands[i].hand[n] == hands[j].hand[n] {
					continue
				} else {
					cond = cardRankings[string(hands[i].hand[n])] < cardRankings[string(hands[j].hand[n])]
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
			total = total + (hand.bid * rank)
			rank++
		}
	}
	fmt.Println(total)
}

func partTwo() {
	cardRankings := map[string]int{
		"J": 0,
		"2": 1,
		"3": 2,
		"4": 3,
		"5": 4,
		"6": 5,
		"7": 6,
		"8": 7,
		"9": 8,
		"T": 9,
		"Q": 10,
		"K": 11,
		"A": 12,
	}
	lines := getFileContentByLine()

	// Slice to hold slices of hands grouped by rank where order is relative
	// ranking of hand types
	handCount := make([][]player, 7)
	for _, line := range lines {
		hand, b, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(b)
		p := player{hand: hand, bid: bid}

		handRank := getHandRanking(strings.ReplaceAll(hand, "J", ""))
		numOfJacks := strings.Count(hand, "J")
		if numOfJacks == 5 {
			handRank = 6
		} else if numOfJacks > 0 {
			handRank = getAdjustedHandRanking(handRank, numOfJacks)
		}
		handCount[handRank] = append(handCount[handRank], p)
	}

	// Order all hands by rank
	for _, hands := range handCount {
		sort.Slice(hands, func(i, j int) bool {
			var cond bool
			for n := 0; n < len(hands[i].hand); n++ {
				if hands[i].hand[n] == hands[j].hand[n] {
					continue
				} else {
					cond = cardRankings[string(hands[i].hand[n])] < cardRankings[string(hands[j].hand[n])]
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
			total = total + (hand.bid * rank)
			rank++
		}
	}
	fmt.Println(total)
}

func main() {
	partOne()
	partTwo()
}
