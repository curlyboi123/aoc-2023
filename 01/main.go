package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("input_long.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	total := 0
	for _, line := range lines {
		var only_ints strings.Builder

		for _, char := range line {
			if unicode.IsDigit(char) {
				only_ints.WriteRune(char)
			}
		}

		output := only_ints.String()
		first_num := string(output[0])
		last_num := string((output[len(output)-1]))

		combined, err := strconv.Atoi(first_num + last_num)
		if err != nil {
			panic(err)
		}

		total = total + combined
	}
	fmt.Println(total)
}
