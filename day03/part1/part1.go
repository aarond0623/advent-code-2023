package part1

import (
	"regexp"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsPartNumber(lines []string, row int, slice []int) bool {
	// Get the width and height of our input
	height := len(lines)
	width := len(lines[0])

	// Subtract and add 1 to row and slice to surround the part number, but don't
	// go out of range
	minRow := max(row-1, 0)
	maxRow := min(row+1, height-1)
	minCol := max(slice[0]-1, 0)
	maxCol := min(slice[1]+1, width) // The slice can go one over

	// Regex for finding symbols (anything not a number or a period)
	re := regexp.MustCompile(`[^\d\.]`)

	for i := minRow; i <= maxRow; i++ {
		if re.MatchString(lines[i][minCol:maxCol]) {
			return true
		}
	}
	// Didn't find a symbol
	return false
}

func Solve(lines []string) int {
	// Initialize sum
	var sum int = 0

	// Regex for finding numbers
	re := regexp.MustCompile(`\d+`)

	// Find numbers on each line
	for i, line := range lines {
		matches := re.FindAllStringIndex(line, -1)
		// Did not find a number
		if matches == nil {
			continue
		}

		// match is a slice for a number on the line
		for _, match := range matches {
			partNumber, _ := strconv.Atoi(line[match[0]:match[1]])
			if IsPartNumber(lines, i, match) {
				sum += partNumber
			}
		}
	}
	return sum
}
