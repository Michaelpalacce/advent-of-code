package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	FiveOfAKind  = 1
	FourOfAKind  = 2
	FullHouse    = 3
	Straight     = 4
	ThreeOfAKind = 5
	TwoPair      = 6
	OnePair      = 7
	HighCard     = 8
)

type Result struct {
	Type     int
	HighCard int
}

type Hand struct {
	Hand    string
	Points  int
	Results []Result
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := make([]Hand, 0)

	for scanner.Scan() {
		line := scanner.Text()

		hand, err := ParseHandLine(line)
		if err != nil {
			log.Fatal(err)
		}

		hands = append(hands, hand)
	}

	for _, hand := range hands {
		fmt.Println(hand)
	}
}

// Example hand: 32T3K 765
// 32T3K = 2 of a kind
// 765 = Points
func ParseHandLine(line string) (Hand, error) {
	parsedLine := strings.Split(line, " ")

	points, err := strconv.Atoi(parsedLine[1])
	if err != nil {
		return Hand{}, err
	}

	hand := Hand{
		Hand:   parsedLine[0],
		Points: points,
	}

	return hand, nil
}

func (h *Hand) GetResults() []Result {
	if len(h.Results) == 0 {
		h.Results = h.calculateResults()
	}

	return h.Results
}

func (h *Hand) calculateResults() []Result {
	results := []Result{}

	pointer := make(map[string]int)

	for _, char := range h.Hand {
		if _, ok := pointer[string(char)]; ok {
			pointer[string(char)]++
		} else {
			pointer[string(char)] = 1
		}
	}

	for char, value := range pointer {
		result := Result{}
		switch value {
		case 5:
			result.Type = FiveOfAKind
			results = append(results, result)
		case 4:
			result.Type = FourOfAKind
			results = append(results, result)
		case 3:
			result.Type = ThreeOfAKind
			results = append(results, result)
		case 2:
			result.Type = OnePair
			results = append(results, result)
		case 1:
			result.Type = HighCard
			results = append(results, result)
		}

		switch char {
		case "A":
			result.HighCard = 14
		case "K":
			result.HighCard = 13
		case "Q":
			result.HighCard = 12
		case "J":
			result.HighCard = 11
		case "T":
			result.HighCard = 10
		default:
			result.HighCard, _ = strconv.Atoi(char)
		}
	}

	return results
}
