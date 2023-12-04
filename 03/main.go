package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func readLines(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func isSymbol(value rune) bool {
	return value != '.' && (value < '0' || value > '9')
}

func getPartNumber(line []rune, start int, currentPosition int) int {

	// If value not a number and start is set and a neighbor exists, compile number from start to j and clear vars
	partNumber := 0
	for k := currentPosition - 1; k >= start; k-- {
		digit := int(line[k]-'0') * int(math.Pow(10, float64(currentPosition-k-1)))
		partNumber += digit
	}
	return partNumber
}

func Answer(lines [][]rune) int {
	sumOfParts := 0
	for i, line := range lines {

		start := -1
		isPart := false

		for j, value := range line {
			// Check if value is number
			// If so, set start position if not set and then check neighbors for symbols
			if value >= '0' && value <= '9' {
				if start == -1 {
					start = j
				}
				if j > 0 && isSymbol(line[j-1]) ||
					j < len(line)-1 && isSymbol(line[j+1]) ||
					i > 0 && isSymbol(lines[i-1][j]) ||
					i < len(lines)-1 && isSymbol(lines[i+1][j]) ||
					i > 0 && j > 0 && isSymbol(lines[i-1][j-1]) ||
					i < len(lines)-1 && j < len(line)-1 && isSymbol(lines[i+1][j+1]) ||
					i > 0 && j < len(line)-1 && isSymbol(lines[i-1][j+1]) ||
					i < len(lines)-1 && j > 0 && isSymbol(lines[i+1][j-1]) {
					isPart = true
				}

			} else if start != -1 {
				// Not a number value, need to reset vars
				if isPart {
					sumOfParts += getPartNumber(line, start, j)
					isPart = false
				}
				start = -1
			}
		}

		// Handle edge case where number is at the end of the line
		if isPart {
			sumOfParts += getPartNumber(line, start, len(line))
		}
	}
	return sumOfParts
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Answer(lines))
	// fmt.Println(Answer2(games))
}
