package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Subset struct {
	red   int
	green int
	blue  int
}

type Game struct {
	number  int
	subsets []Subset
}

func parseGameFromString(input string) *Game {
	var game Game

	re := regexp.MustCompile(`Game (\d+): (.+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		panic("Invalid input")
	}

	gameNumber, err := strconv.Atoi(matches[1])
	if err != nil {
		panic("Invalid game number")
	}
	game.number = gameNumber

	var subsets []Subset
	subsetStrings := strings.Split(matches[2], "; ")
	for _, subsetString := range subsetStrings {
		// Parse subset
		// ex. 6 red, 11 green
		colorParts := strings.Split(subsetString, ", ")
		var subset Subset
		for _, colorPart := range colorParts {
			// ex. 6 red
			parts := strings.Split(colorPart, " ")
			color := parts[1]

			// Parse number
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				panic("Invalid color count")
			}

			switch color {
			case "red":
				subset.red = count
			case "green":
				subset.green = count
			case "blue":
				subset.blue = count
			default:
				panic("Invalid color string")
			}
		}
		subsets = append(subsets, subset)
	}

	game.subsets = subsets
	return &game
}

// isEqual compares two Game objects for equality.
func (game1 Game) isEqual(game2 Game) bool {
	if game1.number != game2.number {
		return false
	}

	if len(game1.subsets) != len(game2.subsets) {
		return false
	}

	for i := range game1.subsets {
		if game1.subsets[i].red != game2.subsets[i].red ||
			game1.subsets[i].green != game2.subsets[i].green ||
			game1.subsets[i].blue != game2.subsets[i].blue {
			return false
		}
	}

	return true
}

func readGames(path string) ([]*Game, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var games []*Game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := parseGameFromString(scanner.Text())
		games = append(games, game)
	}
	return games, scanner.Err()
}

func Answer(games []*Game) int {
	// What games are possible given the following?
	// 12 red cubes, 13 green cubes, and 14 blue cubes
	// Sum their IDs
	var idSum int = 0
	for _, game := range games {
		hasOutOfBounds := false
		for _, subset := range game.subsets {
			if subset.red > 12 || subset.green > 13 || subset.blue > 14 {
				hasOutOfBounds = true
				break
			}
		}
		if !hasOutOfBounds {
			idSum += game.number
		}
	}
	return idSum
}

func Answer2(games []*Game) int {
	// Go over each game and get the minimum number of cubes needed for the game to be possible
	// Then store the sum of the product of these minimums
	productSum := 0
	for _, game := range games {
		minimums := []int{0, 0, 0}
		for _, subset := range game.subsets {
			if subset.red > minimums[0] {
				minimums[0] = subset.red
			}
			if subset.green > minimums[1] {
				minimums[1] = subset.green
			}
			if subset.blue > minimums[2] {
				minimums[2] = subset.blue
			}
		}
		product := minimums[0] * minimums[1] * minimums[2]
		productSum += product
	}
	return productSum
}

func main() {
	games, err := readGames("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Answer(games))
	fmt.Println(Answer2(games))
}
