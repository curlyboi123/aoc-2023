package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func partOne() {
	lines := getFileContentByLine()

	pointsTotal := 0

	for _, line := range lines {
		_, nums, _ := strings.Cut(line, ":")
		winningNums, playerNums, _ := strings.Cut(nums, "|")
		numIntersection := HashGeneric[string](strings.Fields(winningNums), strings.Fields(playerNums))
		if len(numIntersection) > 0 {
			points := int(math.Pow(2, float64((len(numIntersection))-1)))
			pointsTotal = pointsTotal + points
		}
	}
	fmt.Println(pointsTotal)
}

func partTwo() {
	lines := getFileContentByLine()

	total := 0

	m := make(map[int]int)
	for n := 1; n < len(lines)+1; n++ {
		m[n] = 0
	}
	for idx, line := range lines {
		_, nums, _ := strings.Cut(line, ":")
		winningNums, playerNums, _ := strings.Cut(nums, "|")
		numIntersection := HashGeneric[string](strings.Fields(winningNums), strings.Fields(playerNums))

	}

	fmt.Println(total)
}

func main() {
	partOne()
	partTwo()
}
