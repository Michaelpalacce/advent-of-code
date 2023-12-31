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

// Parses the following input
// Time:        59     79     65     75
func parseInput(input string) int {
	input = strings.Replace(input, "Time: ", "", 1)
	input = strings.Replace(input, "Distance: ", "", 1)
	result := ""
	for _, v := range strings.Split(input, " ") {
		if v == "" {
			continue
		}

		result += v
	}

	resultInt, err := strconv.Atoi(result)
	if err != nil {
		log.Fatal(err)
	}

	return resultInt
}

func main() {
	filePath := "input.txt"

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

	race := Race{
		TimeMs: time,
		Milim:  distance,
		// MilimPerTimeMs: 0,
	}

	fmt.Println(race)
	lowerBound := race.CalculateLowerBound()
	upperBound := race.CalculateUpperBound(lowerBound)
	wins := upperBound - lowerBound + 1
	fmt.Printf("{%d:%d}. Total Wins: %d\n", lowerBound, upperBound, wins)
}
