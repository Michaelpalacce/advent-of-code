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
	oneOfAKindBitMask   = 0b00000001
	twoOfAKindBitMask   = 0b00000010
	threeOfAKindBitMask = 0b00000100
	fourOfAKindBitMask  = 0b00001000
	fiveOfAKindBitMask  = 0b00010000
	fullHouseBitMask    = oneOfAKindBitMask | threeOfAKindBitMask
)

type Result struct {
	ResultMask int
	HighCard   int
}

type Hand struct {
	Hand   string
	Points int
	result *Result
}

func (h *Hand) GetResult() *Result {
	if h.result == nil {
		h.result = h.calculateResults()
	}

	return h.result
}

// calculateResults calculates the results of the hand
// and returns a result
// uses bitwise operations to calculate the results
func (h *Hand) calculateResults() *Result {
	pointer := make(map[string]int)

	for _, char := range h.Hand {
		if _, ok := pointer[string(char)]; ok {
			pointer[string(char)]++
		} else {
			pointer[string(char)] = 1
		}
	}

	bitwiseResult := 0b00000000
	for char, value := range pointer {
		if char == "J" {
			continue
		}

		switch value {
		case 5:
			bitwiseResult = bitwiseResult | fiveOfAKindBitMask
		case 4:
			bitwiseResult = bitwiseResult | fourOfAKindBitMask
		case 3:
			bitwiseResult = bitwiseResult | threeOfAKindBitMask
		case 2:
			bitwiseResult = bitwiseResult + oneOfAKindBitMask
		}
	}

	return &Result{
		ResultMask: bitwiseResult,
		HighCard:   0,
	}
}

func (h *Hand) Compare(otherHand Hand) int {
	fmt.Printf("Comparing %+v to %+v\n", h, otherHand)
	handResult := h.GetResult()
	otherHandResult := otherHand.GetResult()

	fmt.Printf("Hand: %s, Result: %d\n", h.Hand, handResult.ResultMask)
	fmt.Printf("Other Hand: %s, Result: %d\n", otherHand.Hand, otherHandResult.ResultMask)

	if handResult.ResultMask > otherHandResult.ResultMask {
		fmt.Printf("Hand %s is better than %s\n", h.Hand, otherHand.Hand)
		return 1
	} else if handResult.ResultMask < otherHandResult.ResultMask {
		fmt.Printf("Hand %s is worse than %s\n", h.Hand, otherHand.Hand)
		return -1
	}

	for i := 0; i < len(h.Hand); i++ {
		value := cardToValue(h.Hand[i])
		otherValue := cardToValue(otherHand.Hand[i])
		fmt.Printf("Comparing %s to %s\n", string(h.Hand[i]), string(otherHand.Hand[i]))
		fmt.Printf("Comparing %d to %d\n", value, otherValue)

		if value > otherValue {
			return 1
		} else if value < otherValue {
			return -1
		}
	}

	return 0
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

func cardToValue(card byte) int {
	valueInt, err := strconv.Atoi(string(card))
	if err != nil {
		switch card {
		case 'T':
			valueInt = 10
		case 'Q':
			valueInt = 12
		case 'K':
			valueInt = 13
		case 'A':
			valueInt = 14
		case 'J':
			valueInt = 0
		}
	}
	return valueInt
}

func quickSort(arr []Hand, low, high int) []Hand {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []Hand) []Hand {
	return quickSort(arr, 0, len(arr)-1)
}

func partition(arr []Hand, low, high int) ([]Hand, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].Compare(pivot) < 0 {
			// if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
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

	output := 0

	sortedHands := quickSortStart(hands)
	for i := len(sortedHands) - 1; i >= 0; i-- {
		output += sortedHands[i].Points * (i + 1)
	}

	fmt.Println(output)
}
