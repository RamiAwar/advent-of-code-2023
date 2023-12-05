package main

import (
	"log"
	"testing"
)

var testInputParsing = struct {
	input string
	want  Card
}{
	`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53`,
	Card{
		winning: map[uint64]struct{}{
			41: {},
			48: {},
			83: {},
			86: {},
			17: {},
		},
		score: 16,
	},
}

var testInput = struct {
	input []string
	want  uint64
}{
	[]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}, 8,
}

var testCountCardsInput = struct {
	input string
	want  []uint64
}{
	input: "test_input.txt",
	want: []uint64{
		1, 2, 4, 8, 14, 1,
	},
}

func TestParseCardFromString(t *testing.T) {
	got := parseCardFromString(testInputParsing.input)

	if !got.isEqual(&testInputParsing.want) {
		t.Errorf("Parsed game different from wanted game: %v != %v", got, testInputParsing.want)
	}
}

func TestAll(t *testing.T) {
	got, err := calculateScore("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	if got != testInput.want {
		t.Errorf("Got %v, want %v", got, testInput.want)
	}
}

func TestCountCards(t *testing.T) {
	got, err := countCards("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(testCountCardsInput.want); i++ {
		if got[i] != testCountCardsInput.want[i] {
			t.Errorf("Got %v, want %v", got[i], testCountCardsInput.want[i])
		}
	}
}
