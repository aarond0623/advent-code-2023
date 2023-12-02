package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(line string) int {
	// Returns an integer found by finding the first and last digits in a string
	// Match first digit
	re_digit1 := regexp.MustCompile(`^\D*(\d)`)
	// Match last digit
	re_digit2 := regexp.MustCompile(`(\d)\D*$`)
	// Get matches
	digit1 := re_digit1.FindStringSubmatch(line)[1]
	digit2 := re_digit2.FindStringSubmatch(line)[1]
	// Add to sum
	value, _ := strconv.Atoi(digit1 + digit2)
	return value
}

func part2(line string) int {
	// Returns an integer found by finding the first number on a line, either a
	// digit or spelled out
	// Map of the names of the digits
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	// Build the query based on the above keys
	re_query1 := `\d`
	for key, _ := range digits {
		re_query1 += `|`
		re_query1 += key
	}

	/* Due to overlapping matches like "twone" or "oneight", we can't trust the
	 * last match with the above query; the first match consumes part of the next
	 * number. Therefore, we'll reverse both the string and the digit names and
	 * find the first occurrence in the reversed string.
	 */
	re_query2 := `\d`
	for key, _ := range digits {
		re_query2 += `|`
		re_query2 += reverse(key)
	}
	// Find the matches
	re_1 := regexp.MustCompile(re_query1)
	re_2 := regexp.MustCompile(re_query2)
	// Get the digits. If we can't find it in our map, it must be a digit already
	digit1, ok := digits[re_1.FindString(line)]
	if !ok {
		digit1 = re_1.FindString(line)
	}
	digit2, ok := digits[reverse(re_2.FindString(reverse(line)))]
	if !ok {
		digit2 = re_2.FindString(reverse(line))
	}

	value, _ := strconv.Atoi(digit1 + digit2)
	return value
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Initialize our sum to 0
	var sum_part1 int = 0
	var sum_part2 int = 0
	// Read input from stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sum_part1 += part1(scanner.Text())
		sum_part2 += part2(scanner.Text())
	}

	fmt.Println(sum_part1)
	fmt.Println(sum_part2)
}
