package part1_test

import (
	"aoc2023/day03/part1"
	"testing"
)

func TestIsPartNumber(t *testing.T) {
	t.Parallel()

	lines := []string{
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

	type testCase struct {
		row   int
		slice []int
		want  bool
	}

	// Part number test cases. Should be true if any symbol is adjacent.
	testCases := []testCase{
		{row: 0, slice: []int{0, 3}, want: true},
		{row: 0, slice: []int{5, 8}, want: false},
		{row: 2, slice: []int{2, 4}, want: true},
		{row: 2, slice: []int{6, 9}, want: true},
		{row: 4, slice: []int{0, 3}, want: true},
		{row: 5, slice: []int{7, 9}, want: false},
		{row: 6, slice: []int{2, 5}, want: true},
		{row: 7, slice: []int{6, 9}, want: true},
		{row: 9, slice: []int{1, 4}, want: true},
		{row: 9, slice: []int{5, 8}, want: true},
	}

	for _, tc := range testCases {
		got := part1.IsPartNumber(lines, tc.row, tc.slice)
		if tc.want != got {
			t.Errorf("IsPartNumber(%d, {%d, %d}): want %t, got %t", tc.row, tc.slice[0], tc.slice[1], tc.want, got)
		}
	}
}

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
		t.Errorf("Solve: want %d, got %d", want, got)
	}
}
