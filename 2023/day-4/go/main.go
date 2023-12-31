package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Winners int
	Copies  int
}

// One line is:
// Card {N}: 69 15 78 85 50 51 57 71 74 58 | 63 79  4 13 94 97 17 10 25 38 87 33 27 86 75 76 99 23 36 35 47 64 41 46 84
// Card {N} is the card number
// The first numbers are the winning numbers, the rest are the numbers on the card
func parseCard(line string) Card {
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
		if cardNumber == "" {
			continue
		}
		for _, winningNumber := range winningNumbers {
			if winningNumber == "" {
				continue
			}
			if cardNumber == winningNumber {
				cardNumberInt, _ := strconv.Atoi(cardNumber)

				winners = append(winners, cardNumberInt)
			}
		}
	}

	return Card{
		Winners: len(winners),
		Copies:  1,
	}
}

// populateCopies is a function that populates the copies
// of the next N cards where N is the number of winners.
// If the current card has 2 copies and 2 winners, the next
// 2 cards will have 2 copies each.
func populateCopies(cardBox []Card) {
	for index, card := range cardBox {
		if card.Winners > 0 {
			for i := 1; i <= card.Winners; i++ {
				if (index + i) >= len(cardBox) {
					break
				}
				cardBox[index+i].Copies += card.Copies
			}
		}
	}
}

// Scan the file and calculate the points
func scanFileContents() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := 0
	cardBox := make([]Card, 0)

	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		cardBox = append(cardBox, card)
	}

	populateCopies(cardBox)

	for _, card := range cardBox {
		output += card.Copies
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result is:", output)
}

func main() {
	scanFileContents()
}
