package part2_test

import (
	"aoc2023/day04/part2"
	"testing"
)

func TestCountWins(t *testing.T) {
	t.Parallel()

	type testCase struct {
		winning_numbers []int
		card_numbers    []int
		want            int
	}

	// Scratch card test cases.
	testCases := []testCase{
		{
			winning_numbers: []int{41, 48, 83, 86, 17},
			card_numbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			want:            4,
		},
		{
			winning_numbers: []int{13, 32, 20, 16, 61},
			card_numbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			want:            2,
		},
		{
			winning_numbers: []int{1, 21, 53, 59, 44},
			card_numbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			want:            2,
		},
		{
			winning_numbers: []int{41, 92, 73, 84, 69},
			card_numbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
			want:            1,
		},
		{
			winning_numbers: []int{87, 83, 26, 28, 32},
			card_numbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
			want:            0,
		},
		{
			winning_numbers: []int{31, 18, 13, 56, 72},
			card_numbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
			want:            0,
		},
	}

	for _, tc := range testCases {
		got := part2.CountWins(tc.winning_numbers, tc.card_numbers)
		if tc.want != got {
			t.Errorf("CountWins(%v, %v): want %d, got %d", tc.winning_numbers, tc.card_numbers, tc.want, got)
		}
	}
}

func TestSolve(t *testing.T) {
	t.Parallel()

	// Test case from Advent of Code
	testCase := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	// What we should get if the function works correctly
	var want int = 30
	// What we actually get
	got := part2.Solve(testCase, 5, 8)

	if want != got {
		t.Errorf("Solve: want %d, got %d", want, got)
	}
}
