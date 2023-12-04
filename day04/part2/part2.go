package part2

import (
	"regexp"
	"slices"
	"strconv"
)

func CountWins(winning_numbers []int, card_numbers []int) int {
	// Number of winning numbers on this card
	var wins int = 0
	for _, n := range card_numbers {
		if slices.Contains(winning_numbers, n) {
			wins++
		}
	}

	return wins
}

func SpawnedCards(wins []int, cards []int, i int) int {
	if cards[i] != 0 {
		return cards[i]
	} else {
		for j := 1; j <= wins[i]; j++ {
			cards[i] += SpawnedCards(wins, cards, j+i)
		}
		cards[i] += 1
		return cards[i]
	}
}

func Solve(lines []string, win_qty int, card_qty int) int {
	// Initialize sum
	var sum int = 0

	// Set up a slice for the wins for each card, as well as an slice for how
	// many cards a card spawns, so we do not have to keep recalculating these
	wins := make([]int, len(lines))
	cards := make([]int, len(lines))

	// Regex for finding numbers
	re := regexp.MustCompile(`\d+`)

	// Find numbers on each line
	for i, line := range lines {
		matches := re.FindAllString(line, -1)
		var winning_numbers []int
		var card_numbers []int

		// Get the winning numbers and card numbers
		for _, match := range matches[1 : win_qty+1] {
			n, _ := strconv.Atoi(match)
			winning_numbers = append(winning_numbers, n)
		}
		for _, match := range matches[win_qty+1 : card_qty+win_qty+1] {
			n, _ := strconv.Atoi(match)
			card_numbers = append(card_numbers, n)
		}

		wins[i] = CountWins(winning_numbers, card_numbers)
	}

	// Get the amount of cards spawned by a card
	for i := 0; i < len(cards); i++ {
		if cards[i] != 0 { // Already done, don't need to revisit
			sum += cards[i]
		} else {
			sum += SpawnedCards(wins, cards, i)
		}
	}

	return sum
}
