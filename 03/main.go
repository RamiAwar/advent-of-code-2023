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
	// Find the sum of all the parts
	// Part is a number that has a neighbor that is a symbol (not '.' though)
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
				// We have a completed number, need to reset vars
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

func isGear(value rune) bool {
	return value == '*'
}

// Create a structure like default dict of sets in python
type DefaultMap struct {
	m map[[2]int]map[[2]int]struct{}
}

func NewDefaultMap() *DefaultMap {
	return &DefaultMap{m: make(map[[2]int]map[[2]int]struct{})}
}

func (d *DefaultMap) SetSubset(key [2]int, subKey [2]int) {
	if d.m[key] == nil {
		d.m[key] = make(map[[2]int]struct{})
	}
	d.m[key][subKey] = struct{}{}
}

func Answer2(lines [][]rune) int {
	// Find the sum of all the gear ratios
	// Gear ratio is the product of ONLY two numbers that have a common neighbor of *
	// If gear connecting more than 2 numbers, ignore it
	sumOfGearRatios := 0

	// Track the gear edges: [gearRow, gearCol] -> [partStartRow, partStartCol]
	gearEdges := NewDefaultMap()
	partNumbers := make(map[[2]int]int)

	for i, line := range lines {

		start := -1
		hasGear := false

		for j, value := range line {
			if value >= '0' && value <= '9' {
				if start == -1 {
					start = j
				}

				// Find connected gears and track them in gearEdges
				// Can use the start position (row, col) to identify a number - that is unique per number!
				currentStartPosition := [2]int{i, start}
				if j > 0 && isGear(line[j-1]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i, j - 1}, currentStartPosition)
				}
				if j < len(line)-1 && isGear(line[j+1]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i, j + 1}, currentStartPosition)
				}
				if i > 0 && isGear(lines[i-1][j]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i - 1, j}, currentStartPosition)
				}
				if i < len(lines)-1 && isGear(lines[i+1][j]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i + 1, j}, currentStartPosition)
				}
				if i > 0 && j > 0 && isGear(lines[i-1][j-1]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i - 1, j - 1}, currentStartPosition)
				}
				if i < len(lines)-1 && j < len(line)-1 && isGear(lines[i+1][j+1]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i + 1, j + 1}, currentStartPosition)
				}
				if i > 0 && j < len(line)-1 && isGear(lines[i-1][j+1]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i - 1, j + 1}, currentStartPosition)
				}
				if i < len(lines)-1 && j > 0 && isGear(lines[i+1][j-1]) {
					hasGear = true
					gearEdges.SetSubset([2]int{i + 1, j - 1}, currentStartPosition)
				}

			} else if start != -1 {
				// We have a completed number
				// Time to add it to our map from start position -> complete number
				// Only care about it if we had any gears though
				if hasGear {
					partNumber := getPartNumber(line, start, j)
					currentStartPosition := [2]int{i, start}
					partNumbers[currentStartPosition] = partNumber
					hasGear = false
				}
				start = -1
			}
		}

		// Handle edge case where number is at the end of the line
		if hasGear {
			partNumber := getPartNumber(line, start, len(line))
			currentStartPosition := [2]int{i, start}
			partNumbers[currentStartPosition] = partNumber
			hasGear = false
		}
	}

	// Now that we have a gearEdges map, time to find all gears that are connecting two numbers exactly
	// Then we can multiply those numbers and add them to our sum
	for _, partStarts := range gearEdges.m {
		if len(partStarts) == 2 {
			// We have a gear connecting two numbers
			// Multiply them and add to sum
			product := 1
			for partStart := range partStarts {
				product *= partNumbers[partStart]
			}

			sumOfGearRatios += product
		}
	}

	return sumOfGearRatios
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Answer(lines))
	fmt.Println(Answer2(lines))
}
