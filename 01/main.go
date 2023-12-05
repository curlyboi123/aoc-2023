package main

import (
	"fmt"
	"os"
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

func part_two() int {
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
	for idx, line := range lines {
		var only_ints strings.Builder

		// Get index in string of first and last occurence of ints as strings in string
		first_num_as_string_idx := -1
		first_num := ""
		last_num_as_string_idx := 0
		last_num := ""
		for k, v := range number_string_to_int {
			// String contains integer as string
			num_as_string_idx := strings.Index(line, k)
			if num_as_string_idx >= 0 {
				if num_as_string_idx < first_num_as_string_idx || first_num_as_string_idx == -1 {
					first_num_as_string_idx = num_as_string_idx
					first_num = v
				}
				if num_as_string_idx > last_num_as_string_idx {
					last_num_as_string_idx = num_as_string_idx
					last_num = v
				}
			}
		}

		for _, char := range line {
			if unicode.IsDigit(char) {
				only_ints.WriteRune(char)
			}
		}
		output := only_ints.String()

		fmt.Print(output)

		foo := ""
		bar := ""
		if len(output) != 0 {
			first_num_as_int := string(output[0])
			last_num_as_int := string((output[len(output)-1]))

			if strings.Index(line, first_num_as_int) < first_num_as_string_idx {
				foo = first_num_as_int
			} else {
				foo = first_num
			}
			if strings.Index(line, last_num_as_int) > last_num_as_string_idx {
				bar = last_num_as_int
			} else {
				bar = last_num
			}
		} else {
			foo = first_num
			bar = last_num
		}

		fmt.Println(idx)

		combined, err := strconv.Atoi(foo + bar)
		if err != nil {
			panic(err)
		}

		total = total + combined
	}
	return total
}

func main() {
	// part_one_result := part_one()
	// fmt.Println(part_one_result)

	part_two_result := part_two()
	fmt.Println(part_two_result)

}
