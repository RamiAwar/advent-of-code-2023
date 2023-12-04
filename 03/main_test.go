package main

import (
	"strings"
	"testing"
)

var testInput = struct {
	input string
	want  int
}{
	`........@1
..........
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
	4362,
}

var testInputGearRatios = struct {
	input string
	want  int
}{
	`
.........1
.........*
.........5
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
	467840,
}

func readInput(input string) [][]rune {
	lines := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			lines = append(lines, []rune(line))
		}
	}
	return lines
}

func TestValidSumOfParts(t *testing.T) {
	lines := readInput(testInput.input)
	got := Answer(lines)
	if got != testInput.want {
		t.Errorf("Valid sum of parts different from wanted: %v != %v", got, testInput.want)
	}
}

func TestValidSumOfGearRatios(t *testing.T) {
	lines := readInput(testInputGearRatios.input)
	got := Answer2(lines)
	if got != testInputGearRatios.want {
		t.Errorf("Valid sum of gear ratios different from wanted: %v != %v", got, testInputGearRatios.want)
	}
}
