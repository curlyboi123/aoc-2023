package main

import (
	"aoc-2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func hashAlgorithm(step string) int {
	value := 0
	for _, char := range step {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

func partOne() {
	lines := utils.GetFileContentByLine()
	steps := strings.Split(lines[0], ",")
	sum := 0
	for _, step := range steps {
		sum += hashAlgorithm(step)
	}
	fmt.Println(sum)
}

type lensDetails struct {
	label       string
	focalLength int
}

func partTwo() {
	lines := utils.GetFileContentByLine()
	steps := strings.Split(lines[0], ",")

	boxes := map[int][]lensDetails{}
	for _, step := range steps {
		if strings.Contains(step, "=") {
			details := strings.Split(step, "=")
			label := details[0]
			focalLength, _ := strconv.Atoi(details[1])
			box := hashAlgorithm(label)
			// Replace lens with label already in box
			lensToAdd := lensDetails{label, focalLength}
			if lensInBox, ok := boxes[box]; ok {
				existingLensIdx := slices.IndexFunc(boxes[box], func(lens lensDetails) bool {
					return lens.label == label
				})
				// Lens does not exist in box
				if existingLensIdx == -1 {
					boxes[box] = append(lensInBox, lensToAdd)
				} else {
					boxes[box][existingLensIdx] = lensToAdd
				}
			} else {
				boxes[box] = []lensDetails{lensToAdd}
			}

		} else if strings.Contains(step, "-") {
			details := strings.Split(step, "-")
			label := details[0]
			box := hashAlgorithm(label)
			// Delete lens with label from box
			boxes[box] = slices.DeleteFunc(boxes[box], func(lens lensDetails) bool {
				return lens.label == label
			})
		}
		fmt.Println("After Step ", step)
		fmt.Println(boxes)
		fmt.Println()
	}
}

func main() {
	partOne()
	partTwo()
}
