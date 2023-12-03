package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type GameDetails struct {
	GameId       int
	Possible     bool
	TotalRed     int
	TotalBlue    int
	TotalGreen   int
	HighestRed   int
	HighestBlue  int
	HighestGreen int
}

func (g *GameDetails) addValue(amount int, color string) {
	switch color {
	case "red":
		g.TotalRed += amount
		if amount > g.HighestRed {
			g.HighestRed = amount
		}
	case "green":
		g.TotalGreen += amount
		if amount > g.HighestGreen {
			g.HighestGreen = amount
		}
	case "blue":
		g.TotalBlue += amount
		if amount > g.HighestBlue {
			g.HighestBlue = amount
		}
	}

	if amount > constraints[color] {
		g.Possible = false
	}
}

func (g *GameDetails) getPower() int {
	return g.HighestBlue * g.HighestRed * g.HighestGreen
}

var constraints = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getGameDetails(line string) GameDetails {
	var (
		err         error
		gameDetails = GameDetails{
			Possible: true,
		}
	)

	fullGame := strings.Split(line, ":")

	// Game 1
	gameId := strings.Split(fullGame[0], " ")

	gameDetails.GameId, err = strconv.Atoi(gameId[1])
	if err != nil {
		log.Fatalf("Game not correct format. Expected: (Game INT), got %s", fullGame[0])
	}

	// 12 red, 2 green, 5 blue; 9 red, 6 green, 4 blue; 10 red, 2 green, 5 blue; 8 blue, 9 red
	sets := strings.Split(fullGame[1], ";")

	for _, set := range sets {
		set = strings.TrimSpace(set)
		// 12 red, 2 green, 5 blue
		values := strings.Split(set, ",")
		for _, value := range values {
			// 12 red
			result := strings.Split(strings.TrimSpace(value), " ")
			numericRes, err := strconv.Atoi(result[0])
			if err != nil {
				log.Fatalf("Result of a set is not a number")
			}
			gameDetails.addValue(numericRes, result[1])
		}
	}

	return gameDetails
}

func scanFileContents() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0

	for scanner.Scan() {
		line := scanner.Text()
		details := getGameDetails(line)
		output += details.getPower()
		// if details.Possible {
		// 	fmt.Printf("Game ID: %d is possible!\n", details.GameId)
		// 	output += details.GameId
		// } else {
		// 	fmt.Printf("Game ID: %d is not possible!\n", details.GameId)
		// }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result is:", output)
}

func main() {
	scanFileContents()
}
