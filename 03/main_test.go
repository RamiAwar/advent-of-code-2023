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
	if Answer(lines) != testInput.want {
		t.Errorf("Valid sum of parts different from wanted: %v != %v", Answer(lines), testInput.want)
	}
}
