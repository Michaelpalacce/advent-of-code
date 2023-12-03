package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var dictionary = []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
var stringDict = map[string]rune{
	"zero":  48,
	"one":   49,
	"two":   50,
	"three": 51,
	"four":  52,
	"five":  53,
	"six":   54,
	"seven": 55,
	"eight": 56,
	"nine":  57,
}

func findAllOccurrences(fullString, substring string) []int {
	var occurrences []int
	startIndex := 0

	for {
		index := strings.Index(fullString[startIndex:], substring)

		if index == -1 {
			break // No more occurrences found
		}

		// Adjust the index based on the starting position
		adjustedIndex := startIndex + index
		occurrences = append(occurrences, adjustedIndex)

		// Move the starting index forward
		startIndex = adjustedIndex + 1
	}

	return occurrences
}

func getCalibrationValue(line string) int {
	var (
		first         rune
		firstPosition int
		last          rune
		lastPosition  int
	)

	for index, char := range line {
		contains := slices.Contains(dictionary, char)
		if contains {
			if first == 0 {
				first = char
				firstPosition = index
				last = char
				lastPosition = index
			} else {
				last = char
				lastPosition = index
			}
		}
	}

	for numberStr, numberRune := range stringDict {
		occurrences := findAllOccurrences(line, numberStr)
		if len(occurrences) > 0 {
			firstPossible := occurrences[0]
			if firstPossible < firstPosition {
				first = numberRune
				firstPosition = firstPossible
			}

			lastPossible := occurrences[len(occurrences)-1]
			if lastPossible > lastPosition {
				last = numberRune
				lastPosition = lastPossible
			}
		}
	}

	calibrationValue, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
	if err != nil {
		log.Fatalf("Could not convert line calibration value to int")
	}

	return calibrationValue
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
		output += getCalibrationValue(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result is:", output)
}

func main() {
	scanFileContents()
}
