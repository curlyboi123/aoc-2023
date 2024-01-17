package main

import (
	"aoc-2023/utils"
	"fmt"
	"math"
	"regexp"
	"slices"
)

func getTotalGalaxyDist(expansionMultiplier int) int {
	lines := utils.GetFileContentByLine()

	emptyGalaxies := map[string][]int{
		"rows": {},
		"cols": {},
	}

	type coord struct {
		row int
		col int
	}

	// Get coords of galaxies
	var galaxyLocations []coord
	re := regexp.MustCompile("#")

	for row, line := range lines {
		galaxyPositions := re.FindAllStringIndex(line, -1)
		for _, pos := range galaxyPositions {
			col := pos[0]
			galaxyLocations = append(galaxyLocations, coord{row, col})
		}
	}

	// Check for empty rows
	for i := 0; i < len(lines); i++ {
		if !slices.ContainsFunc(galaxyLocations, func(c coord) bool {
			return c.row == i
		}) {
			emptyGalaxies["rows"] = append(emptyGalaxies["rows"], i)
		}
	}

	// Check for empty cols
	for i := 0; i < len(lines[0]); i++ {
		if !slices.ContainsFunc(galaxyLocations, func(c coord) bool {
			return c.col == i
		}) {
			emptyGalaxies["cols"] = append(emptyGalaxies["cols"], i)
		}
	}

	// Get distances between galaxies
	distance := 0
	for i := 0; i < len(galaxyLocations); i++ {
		for j := i + 1; j < len(galaxyLocations); j++ {
			rowDist := int(math.Abs(float64(galaxyLocations[i].row - galaxyLocations[j].row)))
			colDist := int(math.Abs(float64(galaxyLocations[i].col - galaxyLocations[j].col)))

			// Get how many empty galaxies between galaxies
			smallestGalaxyCol := int(math.Min(float64(galaxyLocations[i].col), float64(galaxyLocations[j].col)))
			largestGalaxyCol := int(math.Max(float64(galaxyLocations[i].col), float64(galaxyLocations[j].col)))

			smallestGalaxyRow := int(math.Min(float64(galaxyLocations[i].row), float64(galaxyLocations[j].row)))
			largestGalaxyRow := int(math.Max(float64(galaxyLocations[i].row), float64(galaxyLocations[j].row)))

			for k, v := range emptyGalaxies {
				if k == "rows" {
					for _, x := range v {
						if x > smallestGalaxyRow && x < largestGalaxyRow {
							rowDist = rowDist + (expansionMultiplier - 1)
						}
					}
				} else if k == "cols" {
					for _, x := range v {
						if x > smallestGalaxyCol && x < largestGalaxyCol {
							colDist = colDist + (expansionMultiplier - 1)
						}
					}
				}
			}
			distance = distance + rowDist + colDist
		}
	}
	return distance
}

func main() {
	partOne := getTotalGalaxyDist(2)
	partTwo := getTotalGalaxyDist(int(math.Pow10(6)))
	fmt.Println(partOne)
	fmt.Println(partTwo)
}
