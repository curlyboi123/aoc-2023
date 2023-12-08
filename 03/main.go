package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func isNonFullStopSymbol(char rune) bool {
	return string(char) != "." && (unicode.IsSymbol(char) || unicode.IsPunct(char))
}

func main() {
	fileLines := getFileContentByLine()
	for _, line := range fileLines {
		// previous_line := line[y+1]
		// next_line := line[y-1]
		var number strings.Builder

		for _, char := range line {
			if unicode.IsDigit(char) {
				number.WriteRune(char)
			}
			if isNonFullStopSymbol(char) {
				fmt.Println(string(char))
			}

		}

	}
}
