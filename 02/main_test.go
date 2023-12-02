package main

import (
	"testing"
)

var testInputParsing = []struct {
	input string
	want  Game
}{
	{
		"Game 2: 6 red, 11 green; 4 blue, 4 green, 5 red; 11 green, 6 blue, 6 red",
		Game{
			number: 2,
			subsets: []Subset{
				{red: 6, green: 11, blue: 0},
				{red: 5, green: 4, blue: 4},
				{red: 6, green: 11, blue: 6},
			},
		},
	},
}

var testInput = struct {
	input []string
	want  int
}{
	[]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}, 8,
}

var testMinimumSubsetProductSumInput = struct {
	input []string
	want  int
}{
	[]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}, 2286,
}

func TestParseGameFromString(t *testing.T) {
	for _, row := range testInputParsing {
		got := parseGameFromString(row.input)

		if !got.isEqual(row.want) {
			t.Errorf("Parsed game different from wanted game: %v != %v", got, row.want)
		}
	}
}

func TestValidGameCount(t *testing.T) {
	var games []*Game
	for _, row := range testInput.input {
		game := parseGameFromString(row)
		games = append(games, game)
	}

	if Answer(games) != testInput.want {
		t.Errorf("Valid game count different from wanted count: %v != %v", Answer(games), testInput.want)
	}
}

func TestMinimumSubsetProductSum(t *testing.T) {
	var games []*Game
	for _, row := range testMinimumSubsetProductSumInput.input {
		game := parseGameFromString(row)
		games = append(games, game)
	}

	if Answer2(games) != testMinimumSubsetProductSumInput.want {
		t.Errorf("Valid game count different from wanted count: %v != %v", Answer(games), testMinimumSubsetProductSumInput.want)
	}
}
