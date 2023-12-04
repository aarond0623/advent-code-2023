package part1

import (
	"regexp"
	"slices"
	"strconv"
)

func Score(winning_numbers []int, card_numbers []int) int {
	// Number of winning numbers on this card
	var wins int = 0
	for _, n := range card_numbers {
		if slices.Contains(winning_numbers, n) {
			wins++
		}
	}

	if wins == 0 {
		return 0
	} else {
		// Bit shift for 2 ^ (n - 1)
		return 1 << (wins - 1)
	}
}

func Solve(lines []string, win_qty int, card_qty int) int {
	// Initialize sum
	var sum int = 0

	// Regex for finding numbers
	re := regexp.MustCompile(`\d+`)

	// Find numbers on each line
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		var winning_numbers []int
		var card_numbers []int

		// There are always 10 winning numbers and 25 card numbers
		for _, match := range matches[1 : win_qty+1] {
			n, _ := strconv.Atoi(match)
			winning_numbers = append(winning_numbers, n)
		}
		for _, match := range matches[win_qty+1 : card_qty+win_qty+1] {
			n, _ := strconv.Atoi(match)
			card_numbers = append(card_numbers, n)
		}
		sum += Score(winning_numbers, card_numbers)
	}

	return sum
}
