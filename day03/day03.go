package main

import (
	"aoc2023/day03/part1"
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

	fmt.Println(part1.Solve(lines))
}
