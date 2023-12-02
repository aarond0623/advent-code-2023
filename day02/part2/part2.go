package part2

import (
	"regexp"
	"strconv"
)

func power(game string) int {
	// Initialize counts
	red, green, blue := 0, 0, 0
	// Regex to find red, green, blue
	re_red := regexp.MustCompile(`(\d+) red`)
	re_green := regexp.MustCompile(`(\d+) green`)
	re_blue := regexp.MustCompile(`(\d+) blue`)

	// Find all matches
	reds := re_red.FindAllStringSubmatch(game, -1)
	greens := re_green.FindAllStringSubmatch(game, -1)
	blues := re_blue.FindAllStringSubmatch(game, -1)

	// Find the minimum number of cubes needed
	for _, match := range reds {
		n, _ := strconv.Atoi(match[1])
		if n > red {
			red = n
		}
	}
	for _, match := range greens {
		n, _ := strconv.Atoi(match[1])
		if n > green {
			green = n
		}
	}
	for _, match := range blues {
		n, _ := strconv.Atoi(match[1])
		if n > blue {
			blue = n
		}
	}

	// Return the power
	return red * green * blue
}

func Solve(lines []string) int {
	// Initialize sum to 0
	var sum int = 0
	// Add powers to sum
	for _, line := range lines {
		sum += power(line)
	}
	return sum
}
