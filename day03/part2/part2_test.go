package part2_test

import (
	"aoc2023/day03/part2"
	"reflect"
	"testing"
)

func TestIsPossibleRatio(t *testing.T) {
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
		want  []int
	}

	// Gear ratio test cases. Return true if an asterisk is adjacent
	testCases := []testCase{
		{row: 0, slice: []int{0, 3}, want: []int{1, 3}},
		{row: 0, slice: []int{5, 8}, want: nil},
		{row: 2, slice: []int{2, 4}, want: []int{1, 3}},
		{row: 2, slice: []int{6, 9}, want: nil},
		{row: 4, slice: []int{0, 3}, want: []int{4, 3}},
		{row: 5, slice: []int{7, 9}, want: nil},
		{row: 6, slice: []int{2, 5}, want: nil},
		{row: 7, slice: []int{6, 9}, want: []int{8, 5}},
		{row: 9, slice: []int{1, 4}, want: nil},
		{row: 9, slice: []int{5, 8}, want: []int{8, 5}},
	}

	for _, tc := range testCases {
		got := part2.GetPossibleGear(lines, tc.row, tc.slice)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("GetPossibleGear(%d, %v): want %v, got %v", tc.row, tc.slice, tc.want, got)
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
	var want int = 467835
	// What we actually get
	got := part2.Solve(testCase)

	if want != got {
		t.Errorf("Solve: want %d, got %d", want, got)
	}
}
