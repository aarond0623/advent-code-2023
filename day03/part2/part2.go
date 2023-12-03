package part2

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

func GetPossibleGear(lines []string, row int, slice []int) [2]int {
	// Get the width and height of our input
	height := len(lines)
	width := len(lines[0])

	// Subtract and add 1 to row and slice to surround the part number, but don't
	// go out of range
	minRow := max(row-1, 0)
	maxRow := min(row+1, height-1)
	minCol := max(slice[0]-1, 0)
	maxCol := min(slice[1]+1, width) // The slice can go one over

	// Regex for finding gears (*)
	re := regexp.MustCompile(`\*`)

	// Return an index for a gear if we find one
	// Possible TODO? This won't find a gear if a number is adjacent to two gears
	for i := minRow; i <= maxRow; i++ {
		pos := re.FindStringIndex(lines[i][minCol:maxCol])
		if pos != nil {
			return [2]int{i, minCol + pos[0]}
		}
	}
	return [2]int{-1, -1}
}

func Solve(lines []string) int {
	// Initialize sum
	var sum int = 0

	// Map for gear coordinates to the numbers touching that gear
	gears := map[[2]int][]int{}

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
			pos := GetPossibleGear(lines, i, match)
			// If we find a gear, append the number to the gear's coordinate
			if pos != [2]int{-1, -1} {
				gears[pos] = append(gears[pos], partNumber)
			}
		}
	}

	// Only add power to the sum if the length of the part numbers is exactly 2
	for _, partNumbers := range gears {
		if len(partNumbers) == 2 {
			power := partNumbers[0] * partNumbers[1]
			sum += power
		}
	}

	return sum
}
