package part1

import (
	"regexp"
	"strconv"
)

func possible(game string, max_red int, max_green int, max_blue int) bool {
	// Regex to find red, green, blue
	re_red := regexp.MustCompile(`(\d+) red`)
	re_green := regexp.MustCompile(`(\d+) green`)
	re_blue := regexp.MustCompile(`(\d+) blue`)

	// Find all matches
	reds := re_red.FindAllStringSubmatch(game, -1)
	greens := re_green.FindAllStringSubmatch(game, -1)
	blues := re_blue.FindAllStringSubmatch(game, -1)

	// If any match is over max, return false
	for _, match := range reds {
		n, _ := strconv.Atoi(match[1])
		if n > max_red {
			return false
		}
	}
	for _, match := range greens {
		n, _ := strconv.Atoi(match[1])
		if n > max_green {
			return false
		}
	}
	for _, match := range blues {
		n, _ := strconv.Atoi(match[1])
		if n > max_blue {
			return false
		}
	}

	// Otherwise, this is a possible game
	return true
}

func Solve(lines []string) int {
	// Initialize sum to 0
	var sum int = 0
	// Iterate over lines and get the ID
	for _, line := range lines {
		re_id := regexp.MustCompile(`Game (\d+):`)
		id, _ := strconv.Atoi(re_id.FindStringSubmatch(line)[1])
		// If this game is possible, add the ID to the sum
		if possible(line, 12, 13, 14) {
			sum += id
		}
	}
	return sum
}
