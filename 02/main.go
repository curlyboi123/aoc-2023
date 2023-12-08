package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_file_content_by_line() []string {
	file_path := os.Args[1]
	content, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}

func part_one() {
	file_lines := get_file_content_by_line()
	total := 0
	for _, line := range file_lines {
		ball_totals := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		number_and_results := strings.Split(line, ":")
		game_number, err := strconv.Atoi(strings.Split(number_and_results[0], " ")[1])
		if err != nil {
			panic(err)
		}
		game_results := strings.Split(number_and_results[1], ";")

		game_possible := true
		for _, round := range game_results {
			round_results := strings.Split(round, ",")
			for _, result := range round_results {
				s := strings.Split(strings.TrimSpace(result), " ")
				amount, colour := s[0], s[1]
				i, err := strconv.Atoi(amount)
				if err != nil {
					panic(err)
				}
				ball_totals[colour] = i
			}
			if ball_totals["red"] > 12 || ball_totals["green"] > 13 || ball_totals["blue"] > 14 {
				game_possible = false
				break
			}
		}
		if game_possible {
			total = total + game_number
		}
	}
	fmt.Println(total)
}

func part_two() {
	file_lines := get_file_content_by_line()

	total := 0
	for _, line := range file_lines {
		fewest_possible_balls := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		results := strings.Split(strings.Split(line, ":")[1], ";")

		for _, round := range results {
			round_results := strings.Split(round, ",")
			for _, result := range round_results {
				s := strings.Split(strings.TrimSpace(result), " ")
				colour := s[1]
				amount, err := strconv.Atoi(s[0])
				if err != nil {
					panic(err)
				}
				if amount > fewest_possible_balls[colour] {
					fewest_possible_balls[colour] = amount
				}
			}
		}
		power := 1
		for _, v := range fewest_possible_balls {
			power = power * v
		}
		total = total + power
	}
	fmt.Println(total)
}

func main() {
	part_one()
	part_two()
}
