package main

import (
	"fmt"
	"regexp"
	"strings"

	"example.com/utils"
)

func main() {
	lines := utils.GetFileContentByLine()
	for _, line := range lines {
		s := strings.Fields(line)
		springs, damagedSpringGroups := s[0], s[1]
		fmt.Println("Springs:", springs)
		fmt.Println("Damaged spring groups:", damagedSpringGroups)

		brokenUnknownRe := regexp.MustCompile(`(\?|#)+`)

		brokenUnknownIndexes := brokenUnknownRe.FindAllStringIndex(springs, -1)
		fmt.Println(brokenUnknownIndexes)

		brokenRe := regexp.MustCompile(`#+`)
		brokenIndexes := brokenRe.FindAllStringIndex(springs, -1)
		fmt.Println(brokenIndexes)

		fmt.Println()
	}
}
