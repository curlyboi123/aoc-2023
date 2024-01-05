package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func partOne() {
	lines := getFileContentByLine()

	_, timesStr, _ := strings.Cut(lines[0], ":")
	times := strings.Fields(strings.Trim(timesStr, ""))

	_, distStr, _ := strings.Cut(lines[1], ":")
	distances := strings.Fields(strings.Trim(distStr, ""))

	total := 1

	for raceNum := 0; raceNum < len(times); raceNum++ {
		time, _ := strconv.Atoi(times[raceNum])
		recordDistance, _ := strconv.Atoi(distances[raceNum])

		var recordBreakingHoldTimes []int
		for holdTime := 1; holdTime < time; holdTime++ {
			remainingTime := time - holdTime
			travelDistance := holdTime * remainingTime

			if travelDistance > recordDistance {
				recordBreakingHoldTimes = append(recordBreakingHoldTimes, holdTime)
			}
		}
		total = total * len(recordBreakingHoldTimes)
	}
	fmt.Println(total)
}

func partTwo() {
	lines := getFileContentByLine()

	_, t, _ := strings.Cut(lines[0], ":")
	t = strings.ReplaceAll(t, " ", "")

	_, d, _ := strings.Cut(lines[1], ":")
	d = strings.ReplaceAll(d, " ", "")

	time, _ := strconv.Atoi(t)
	recordDistance, _ := strconv.Atoi(d)

	numHoldTimesBeatingRecord := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		remainingTime := time - holdTime
		travelDistance := holdTime * remainingTime

		if travelDistance > recordDistance {
			numHoldTimesBeatingRecord++
		}
	}
	fmt.Println(numHoldTimesBeatingRecord)
}

func main() {
	partOne()
	partTwo()
}
