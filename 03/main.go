package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func partOne() {
	lines := getFileContentByLine()
	numberRegex := regexp.MustCompile("[0-9]+")
	symbolRegex := regexp.MustCompile(`[^0-9.]+`)

	total := 0

	for idx, line := range lines {
		lineLength := len(line)
		numberMatches := numberRegex.FindAllStringSubmatchIndex(line, -1)

		for _, numberMatch := range numberMatches {
			number := line[numberMatch[0]:numberMatch[1]]

			// Check for symbol before number
			if numberMatch[0] != 0 {
				previousCharacter := string(line[numberMatch[0]-1])
				if symbolRegex.MatchString(previousCharacter) {
					i, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					total = total + i
					continue
				}
			}

			// Check for symbol after number
			if numberMatch[1] != lineLength {
				nextCharacter := string(line[numberMatch[1]])
				if symbolRegex.MatchString(nextCharacter) {
					i, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					total = total + i
					continue
				}
			}

			// Calculate indexes so that it doesn't go out of range
			var startIndex int
			var endIndex int
			if numberMatch[0] == 0 {
				startIndex = 0
			} else {
				startIndex = numberMatch[0] - 1
			}
			if numberMatch[1] == lineLength {
				endIndex = lineLength
			} else {
				endIndex = numberMatch[1] + 1
			}

			// Check for symbol on line above number
			if idx != 0 {
				prevLineSlice := lines[idx-1][startIndex:endIndex]
				if symbolRegex.MatchString(prevLineSlice) {
					i, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					total = total + i
					continue
				}
			}

			// Check for symbol on line below number
			if idx != len(lines)-1 {
				nextLineSlice := lines[idx+1][startIndex:endIndex]
				if symbolRegex.MatchString(nextLineSlice) {
					i, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					total = total + i
					continue
				}
			}

		}
	}
	fmt.Println(total)
}

func partTwo() {
	lines := getFileContentByLine()
	numberRegex := regexp.MustCompile("[0-9]+")
	asterixRegex := regexp.MustCompile("\\*")

	total := 0

	for idx, line := range lines {
		lineLength := len(line)
		asterixMatches := asterixRegex.FindAllStringIndex(line, -1)
		for _, asterixMatch := range asterixMatches {
			var adjacentNumbers []string
			// Check for adjacent number to the left of asterix
			if asterixMatch[0] != 0 {
				prevChar := string(line[asterixMatch[0]-1])
				if numberRegex.MatchString(prevChar) {
					numbersToLeft := numberRegex.FindAllString(line[:asterixMatch[0]], -1)
					fullNumber := numbersToLeft[len(numbersToLeft)-1]
					adjacentNumbers = append(adjacentNumbers, fullNumber)
				}
			}
			// Check for adjacent number to the right of asterix
			if asterixMatch[1] != lineLength {
				nextChar := string(line[asterixMatch[1]])
				if numberRegex.MatchString(nextChar) {
					fullNumber := numberRegex.FindAllString(line[asterixMatch[1]:], 1)[0]
					adjacentNumbers = append(adjacentNumbers, fullNumber)
				}
			}
			// Calculate indexes so that it doesn't go out of range
			var startIndex int
			var endIndex int
			if asterixMatch[0] == 0 {
				startIndex = 0
			} else {
				startIndex = asterixMatch[0] - 1
			}
			if asterixMatch[1] == lineLength {
				endIndex = lineLength
			} else {
				endIndex = asterixMatch[1] + 1
			}
			// Check for adjacent number on the line below the asterix
			if idx != len(lines)-1 {
				belowLineChars := lines[idx+1][startIndex:endIndex]
				if numberRegex.MatchString(belowLineChars) {
					numbersBelowIndexes := numberRegex.FindAllStringIndex(lines[idx+1], -1)
					for _, belowNumberIdx := range numbersBelowIndexes {
						if asterixMatch[0] <= belowNumberIdx[1] && asterixMatch[0] >= belowNumberIdx[0]-1 {
							fullNumber := lines[idx+1][belowNumberIdx[0]:belowNumberIdx[1]]
							adjacentNumbers = append(adjacentNumbers, fullNumber)
						}
					}
				}
			}
			// Check for adjacent number on the line above the asterix
			if idx != 0 {
				aboveLineChars := lines[idx-1][startIndex:endIndex]
				if numberRegex.MatchString(aboveLineChars) {
					numbersAboveIndexes := numberRegex.FindAllStringIndex(lines[idx-1], -1)
					for _, aboveNumberIdx := range numbersAboveIndexes {
						if asterixMatch[0] <= aboveNumberIdx[1] && asterixMatch[0] >= aboveNumberIdx[0]-1 {
							fullNumber := lines[idx-1][aboveNumberIdx[0]:aboveNumberIdx[1]]
							adjacentNumbers = append(adjacentNumbers, fullNumber)
						}
					}
				}
			}

			if len(adjacentNumbers) == 2 {
				i, _ := strconv.Atoi(adjacentNumbers[0])
				j, _ := strconv.Atoi(adjacentNumbers[1])
				product := i * j
				total = total + product
			}
		}
	}
	fmt.Println(total)
}

func main() {
	// partOne()
	partTwo()
}
