package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	TimeMs int
	Milim  int
}

func (r *Race) CalculateLowerBound() int {
	velocity := 1
	for {
		distance := velocity * (r.TimeMs - velocity)
		if distance > r.Milim {
			return velocity
		}

		velocity++
		if velocity > r.TimeMs {
			return -1
		}
	}
}

func (r *Race) CalculateUpperBound(lowerBound int) int {
	return r.TimeMs - lowerBound
}

// func (r *Race) CalculatePossibleWins(race Race) []int {
// }

// Parses the following input
// Time:        59     79     65     75
func parseInput(input string) []int {
	input = strings.Replace(input, "Time: ", "", 1)
	input = strings.Replace(input, "Distance: ", "", 1)
	result := []int{}
	for _, v := range strings.Split(input, " ") {
		if v == "" {
			continue
		}

		i, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			panic(err)
		}

		result = append(result, i)
	}
	return result
}

func main() {
	filePath := "debug.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := parseInput(scanner.Text())

	scanner.Scan()
	distance := parseInput(scanner.Text())

	races := []Race{}

	for i := 0; i < len(time); i++ {
		races = append(races, Race{
			TimeMs: time[i],
			Milim:  distance[i],
			// MilimPerTimeMs: 0,
		})
	}

	result := 1 // So we can multiply

	for _, race := range races {
		fmt.Println(race)
		lowerBound := race.CalculateLowerBound()
		upperBound := race.CalculateUpperBound(lowerBound)
		wins := upperBound - lowerBound + 1
		fmt.Printf("{%d:%d}. Total Wins: %d\n", lowerBound, upperBound, wins)

		result *= wins
		// fmt.Printf("%d", race.CalculateLowerBound())
	}

	fmt.Println(result)
}
