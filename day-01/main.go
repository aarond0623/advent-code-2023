package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	/* PART 1 */
	// Initialize our sum to 0
	var sum int = 0
	// Read input from stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Match first digit
		re_digit1 := regexp.MustCompile(`^\D*(\d)`)
		// Match last digit
		re_digit2 := regexp.MustCompile(`(\d)\D*$`)
		// Get matches
		digit1 := re_digit1.FindStringSubmatch(scanner.Text())[1]
		digit2 := re_digit2.FindStringSubmatch(scanner.Text())[1]
		// Add to sum
		value, _ := strconv.Atoi(digit1 + digit2)
		sum += value
	}

	fmt.Println(sum)
}
