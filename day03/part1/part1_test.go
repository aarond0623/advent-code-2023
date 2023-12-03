package part1_test

import (
	"aoc2023/day03/part1"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	// Test case from Advent of Code
	testCase := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	// What we should get if the function works correctly
	var want int = 4361
	// What we actually get
	got := part1.Solve(testCase)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
