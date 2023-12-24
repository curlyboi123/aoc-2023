package main

import (
	"fmt"
	"os"
	"strings"
)

func getFileContentByLine() []string {
	filePath := os.Args[1]
	content, _ := os.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func main() {
	lines := getFileContentByLine()
	_, seedsStr, _ := strings.Cut(lines[0], ":")
	seeds := strings.Fields(seedsStr)
	fmt.Println(seeds)
	for _, line := range lines {
		fmt.Println(line)
	}
}
