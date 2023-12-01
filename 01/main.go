package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetDigits(input string) int {
	var first_digit int

	for _, c := range input {
		if unicode.IsDigit(c) {
			first_digit = int(c - '0')
			break
		}
	}

	// Loop in reverse and find last digit
	var last_digit int
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(input[i])) {
			last_digit = int(input[i] - '0')
			break
		}
	}

	// Efficiently construct integer from the two digits
	var number int = first_digit*10 + last_digit
	return number
}

func GetDigitsIncludingWords(input string) int {
	digit_words := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// Loop through input and find first digit or digit word
	var first_digit int
	var found bool = false
	for i, c := range input {
		if found {
			break
		}

		if unicode.IsDigit(c) {
			first_digit = int(c - '0')
			break
		} else {
			// Check for digit word
			for word, digit := range digit_words {
				// First check first char matching
				// Efficient early stopping
				if rune(word[0]) == c {
					// Check if word fits in input to begin with
					if i+len(word) > len(input) {
						continue
					}

					// Check for full word match
					if input[i:i+len(word)] == word {
						first_digit = digit
						found = true
						break
					}
				}
			}
		}
	}

	// Loop in reverse and find last digit or digit word
	var last_digit int
	found = false
	for i := len(input) - 1; i >= 0; i-- {
		if found {
			break
		}

		c := rune(input[i])
		if unicode.IsDigit(c) {
			last_digit = int(c - '0')
			break
		} else {
			// Check for digit word
			for word, digit := range digit_words {
				// First check first char matching
				// Efficient early stopping
				if rune(word[len(word)-1]) == c {
					// Check if word fits in input to begin with
					if i-len(word)+1 < 0 {
						continue
					}

					// Check for full word match
					if input[i-len(word)+1:i+1] == word {
						last_digit = digit
						found = true
						break
					}
				}
			}
		}
	}

	return first_digit*10 + last_digit
}

func Answer(input []string) int {
	var sum int

	for _, line := range input {
		sum += GetDigits(line)
	}

	return sum
}

func Answer2(input []string) int {
	var sum int
	for _, line := range input {
		sum += GetDigitsIncludingWords(line)
	}
	return sum
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Answer(lines))
	fmt.Println(Answer2(lines))
}
