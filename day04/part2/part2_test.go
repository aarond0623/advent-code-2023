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

func TestSpawnedCards(t *testing.T) {
	t.Parallel()

	wins := []int{4, 2, 2, 1, 0, 0}

	type testCase struct {
		cards []int
		index int
		want  int
	}

	/* Scratch card test cases.
	* Want values were calculated by hand. Here is an example for card 1 showing
	* what each card spawns
	1 > (2 > (3 > (4 > (5), 5), 4 > (5)), 3 > (4 > (5), 5), 4 > (5), 5)
	*/
	testCases := []testCase{
		{cards: []int{0, 0, 0, 0, 0, 0}, index: 0, want: 15},
		{cards: []int{0, 0, 0, 0, 1, 0}, index: 0, want: 15},
		{cards: []int{0, 0, 0, 2, 1, 0}, index: 0, want: 15},
		{cards: []int{0, 0, 4, 2, 1, 0}, index: 0, want: 15},
		{cards: []int{0, 7, 4, 2, 1, 0}, index: 0, want: 15},
		{cards: []int{15, 7, 4, 2, 1, 0}, index: 0, want: 15},
		{cards: []int{0, 0, 0, 0, 0, 0}, index: 1, want: 7},
		{cards: []int{0, 0, 0, 0, 1, 0}, index: 1, want: 7},
		{cards: []int{0, 0, 0, 2, 1, 0}, index: 1, want: 7},
		{cards: []int{0, 0, 4, 2, 1, 0}, index: 1, want: 7},
		{cards: []int{0, 7, 4, 2, 1, 0}, index: 1, want: 7},
		{cards: []int{0, 0, 0, 0, 0, 0}, index: 2, want: 4},
		{cards: []int{0, 0, 0, 0, 1, 0}, index: 2, want: 4},
		{cards: []int{0, 0, 0, 2, 1, 0}, index: 2, want: 4},
		{cards: []int{0, 0, 4, 2, 1, 0}, index: 2, want: 4},
		{cards: []int{0, 0, 0, 0, 0, 0}, index: 3, want: 2},
		{cards: []int{0, 0, 0, 0, 1, 0}, index: 3, want: 2},
		{cards: []int{0, 0, 0, 2, 1, 0}, index: 3, want: 2},
		{cards: []int{0, 0, 0, 0, 0, 0}, index: 4, want: 1},
		{cards: []int{0, 0, 0, 0, 1, 0}, index: 4, want: 1},
		{cards: []int{0, 0, 0, 0, 0, 0}, index: 5, want: 1},
		{cards: []int{0, 0, 0, 0, 0, 1}, index: 5, want: 1},
	}

	for _, tc := range testCases {
		got := part2.SpawnedCards(wins, tc.cards, tc.index)
		if tc.want != got {
			t.Errorf("SpawnedCards(%v, %d): want %d, got %d", tc.cards, tc.index, tc.want, got)
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
