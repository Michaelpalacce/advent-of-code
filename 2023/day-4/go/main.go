package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// One line is:
// Card {N}: 69 15 78 85 50 51 57 71 74 58 | 63 79  4 13 94 97 17 10 25 38 87 33 27 86 75 76 99 23 36 35 47 64 41 46 84
// Card {N} is the card number
// The first numbers are the winning numbers, the rest are the numbers on the card
func getWinners(line string) []int {
	sections := strings.Split(line, "|")
	winningSection := sections[0]

	winningSection = strings.Split(winningSection, ":")[1]
	winningSection = strings.TrimSpace(winningSection)
	winningNumbers := strings.Split(winningSection, " ")

	cardSection := sections[1]
	cardSection = strings.TrimSpace(cardSection)
	cardNumbers := strings.Split(cardSection, " ")

	winners := make([]int, 0)
	for _, cardNumber := range cardNumbers {
		for _, winningNumber := range winningNumbers {
			if cardNumber == winningNumber {
				cardNumberInt, _ := strconv.Atoi(cardNumber)

				winners = append(winners, cardNumberInt)
			}
		}
	}

	return winners
}

func calculatePoints(winners []int) int {
	points := 0
	for i := range winners {
		if i == 0 {
			points = 1
		} else {
			points *= 2
		}
	}

	return points
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
		winners := getWinners(line)
		output += calculatePoints(winners)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result is:", output)
}

func main() {
	scanFileContents()
}
