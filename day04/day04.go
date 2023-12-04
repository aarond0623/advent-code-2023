package main

import (
	"aoc2023/day04/part1"
	"aoc2023/day04/part2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from stdin line by line
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(part1.Solve(lines, 10, 25))
	fmt.Println(part2.Solve(lines, 10, 25))
}
