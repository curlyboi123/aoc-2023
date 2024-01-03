package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type mapping struct {
	sourceMin, sourceMax, change int
}

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func partOne(mappings [][]mapping) {
	lines := getFileContentByLine()
	_, seedsStr, _ := strings.Cut(lines[0], ":")
	seeds := strings.Fields(seedsStr)
	var locationValues []int
	for _, seed := range seeds {
		value, _ := strconv.Atoi(seed)
		for _, mapping := range mappings {
			for _, mappingRange := range mapping {
				if value >= mappingRange.sourceMin && value <= mappingRange.sourceMax {
					value = value + mappingRange.change
					break
				}
			}
		}
		locationValues = append(locationValues, value)
	}
	fmt.Println(slices.Min(locationValues))
}

func partTwo(mappings [][]mapping) {
	lines := getFileContentByLine()
	_, s, _ := strings.Cut(lines[0], ":")
	seedRanges := strings.Fields(s)

	result := math.MaxInt
	var resultMutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < len(seedRanges)-1; i += 2 {
		seedMin, _ := strconv.Atoi(seedRanges[i])
		seedRange, _ := strconv.Atoi(seedRanges[i+1])
		seedMax := seedMin + seedRange
		wg.Add(1)
		go func() {
			defer wg.Done()
			subResult := math.MaxInt
			for j := seedMin; j < seedMax; j++ {
				value := j
				for _, mapping := range mappings {
					for _, mappingRange := range mapping {
						if value >= mappingRange.sourceMin && value <= mappingRange.sourceMax {
							value = value + mappingRange.change
							break
						}
					}
				}
				subResult = min(subResult, value)
			}
			resultMutex.Lock()
			defer resultMutex.Unlock()
			result = min(result, subResult)
		}()
	}
	wg.Wait()
	fmt.Println(result)
}

func main() {
	lines := getFileContentByLine()

	// Create mappings
	var mappings [][]mapping
	var innerMapping []mapping
	var mapNames []string
	for idx, line := range lines[2:] {
		if strings.Contains(line, "map") {
			mapName := strings.Fields(line)[0]
			mapNames = append(mapNames, mapName)
			continue
		} else if line != "" {
			var destinationMin, sourceMin, mappingRange int
			fmt.Sscanf(line, "%d %d %d", &destinationMin, &sourceMin, &mappingRange)
			mapping := mapping{sourceMin, sourceMin + mappingRange - 1, destinationMin - sourceMin}
			innerMapping = append(innerMapping, mapping)
		}
		if line == "" || idx == len(lines[2:])-1 {
			if len(innerMapping) != 0 {
				mappings = append(mappings, innerMapping)
			}
			innerMapping = nil
		}
	}

	partOne(mappings)
	partTwo(mappings)
}
