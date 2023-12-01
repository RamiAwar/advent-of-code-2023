package main

import (
	"testing"
)

var testInput = []struct {
	input string
	want  int
}{
	{"1abc2", 12},
	{"pqr3stu8vwx", 38},
	{"a1b2c3d4e5f", 15},
	{"treb7uchet", 77},
}

// Extend simpleTestInput with more test cases
var testInput2 = []struct {
	input string
	want  int
}{
	{"1abc2", 12},
	{"pqr3stu8vwx", 38},
	{"a1b2c3d4e5f", 15},
	{"treb7uchet", 77},
	{"one", 11},
	{"two1nine", 29},
	{"eightwothree", 83},
	{"abcone2threexyz", 13},
	{"xtwone3four", 24},
	{"4nineightseven2", 42},
	{"zoneight234", 14},
	{"7pqrstsixteen", 76},
}

var realInput, _ = readLines("input.txt")

func TestGetDigits(t *testing.T) {
	for _, row := range testInput {
		got := GetDigits(row.input)
		if got != row.want {
			t.Errorf("GetDigits(%q) = %d, want %d", row.input, got, row.want)
		}
	}
}

func TestGetDigitsIncludingWords(t *testing.T) {
	for _, row := range testInput2 {
		got := GetDigitsIncludingWords(row.input)
		if got != row.want {
			t.Errorf("GetDigitsIncludingWords(%q) = %d, want %d", row.input, got, row.want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Answer(realInput)
	}
}
