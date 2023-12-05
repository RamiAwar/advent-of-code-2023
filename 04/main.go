package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning map[uint64]struct{}
	score   uint64
	matches uint64
}

func (c *Card) isEqual(other *Card) bool {
	if len(c.winning) != len(other.winning) {
		return false
	}
	for i := range c.winning {
		if c.winning[i] != other.winning[i] {
			return false
		}
	}

	return true
}

func parseCardFromString(input string) *Card {
	var card Card
	card.winning = make(map[uint64]struct{})

	// Fast string split
	i := strings.IndexByte(input, ':')
	input = input[i+1:]

	j := strings.IndexByte(input, '|')
	winning, chosen := input[:j], input[j+1:]

	// Parse space separated numbers as ints, ignoring leading and trailing whitespace
	winningNumbers := strings.Fields(winning)
	chosenNumbers := strings.Fields(chosen)

	for _, number := range winningNumbers {
		intNumber, err := strconv.ParseUint(number, 10, 32)
		if err != nil {
			panic("Invalid winning number")
		}
		card.winning[intNumber] = struct{}{}
	}

	var currentScore uint64 = 0
	var matches uint64 = 0
	for _, number := range chosenNumbers {
		intNumber, err := strconv.ParseUint(number, 10, 32)
		if err != nil {
			panic("Invalid chosen number")
		}

		// If chosen number is in winning numbers, bit shift left by 1 (powers of 2)
		if _, ok := card.winning[intNumber]; ok {
			matches += 1
			if currentScore == 0 {
				currentScore = 1
			} else {
				currentScore = currentScore << 1
			}
		}
	}

	card.score = currentScore
	card.matches = matches
	return &card
}

func calculateScore(path string) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var totalScore uint64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := parseCardFromString(scanner.Text())
		totalScore += card.score
	}
	return totalScore, scanner.Err()
}

func countCards(path string) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var cardCounts []uint64 = make([]uint64, 205)
	scanner := bufio.NewScanner(file)
	var i int = 0
	for scanner.Scan() {
		card := parseCardFromString(scanner.Text())
		cardCounts[i] += 1
		for j := 1; j <= int(card.matches); j++ {
			// Increment counts of current card up until the number of matches
			if cardCounts[i] > 1 {
				cardCounts[i+j] += cardCounts[i]
			} else {
				cardCounts[i+j] += 1
			}
		}
		i += 1
	}

	// Calculate sum of cards
	sum := uint64(0)
	for _, cardCount := range cardCounts {
		sum += cardCount
	}

	return sum, scanner.Err()
}

func main() {
	score, err := calculateScore("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	score2, err := countCards("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
	fmt.Println(score2)
}
