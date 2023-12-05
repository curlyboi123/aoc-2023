package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func get_file_content_by_line() []string {
	file_path := os.Args[1]
	content, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func part_one() int {
	lines := get_file_content_by_line()
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
	return total
}

func part_two() {
	lines := get_file_content_by_line()

	number_string_to_int := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	total := 0
	for _, line := range lines {
		fmt.Println(line)
		// Sanitise input
		line = strings.ReplaceAll(line, "oneeight", "oneeight")
		line = strings.ReplaceAll(line, "threeight", "threeeight")
		line = strings.ReplaceAll(line, "fiveight", "fiveeight")
		line = strings.ReplaceAll(line, "nineight", "nineeight")
		line = strings.ReplaceAll(line, "twone", "twoone")
		line = strings.ReplaceAll(line, "sevenine", "sevennine")
		line = strings.ReplaceAll(line, "eightwo", "eighttwo")

		fmt.Println(line)
		re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
		match_slices := re.FindAllStringSubmatchIndex(line, -1)

		first_match := line[match_slices[0][0]:match_slices[0][1]]
		fmt.Println(first_match)
		if len(first_match) > 1 {
			first_match = number_string_to_int[first_match]
		}

		last_match := line[match_slices[len(match_slices)-1][0]:match_slices[len(match_slices)-1][1]]
		fmt.Println(last_match)
		if len(last_match) > 1 {
			last_match = number_string_to_int[last_match]
		}

		combined := first_match + last_match
		fmt.Println(combined)

		i, err := strconv.Atoi(combined)
		if err != nil {
			panic(err)
		}
		total = total + i
		fmt.Println()
	}
	fmt.Println(total)
}

func main() {
	// part_one_result := part_one()
	// fmt.Println(part_one_result)

	part_two()
}
