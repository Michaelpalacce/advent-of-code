package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	Destination int
	Source      int
	Range       int
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
	rawSeeds := parseNumbers(scanner.Text())

	// Split seeds into touples
	seedRanges := make([][]int, 0)
	for i := 0; i < len(rawSeeds); i += 2 {
		seedRanges = append(seedRanges, []int{rawSeeds[i], rawSeeds[i+1]})
	}

	// Arbitrarily read the next line
	scanner.Scan()

	maps := readMaps(scanner)

	nextRanges := seedRanges

	for _, currentMap := range maps {
		fmt.Println("currentMap", currentMap)
		currentRanges := nextRanges
		nextRanges = make([][]int, 0)
		for _, currentRange := range currentRanges {
			fmt.Println("b", currentRange)
			for _, mapping := range currentMap {
				a1, a2 := float64(mapping.Source), float64(mapping.Source+mapping.Range)
				b1, b2 := float64(currentRange[0]), float64(currentRange[0]+currentRange[1])

				// Check for overlap
				overlapStart := math.Max(a1, b1)
				overlapEnd := math.Min(a2, b2)

				if overlapStart < overlapEnd {
					fmt.Println("overlap", overlapStart, overlapEnd)
					overlapLength := overlapEnd - overlapStart
					overlappedRange := []int{int(mapping.Destination + (int(overlapStart) - mapping.Source)), int(overlapLength)}
					fmt.Println("overlappedRange", overlappedRange)
					nextRanges = append(nextRanges, overlappedRange)
					if overlapStart > b1 {
						nextRanges = append(nextRanges, []int{currentRange[0], int(overlapStart - b1)})
						fmt.Println("leftover", []int{currentRange[0], int(overlapStart - b1 - 1)})
					}
					if overlapEnd < b2 {
						nextRanges = append(nextRanges, []int{int(overlapEnd), int(b2 - overlapEnd)})

						fmt.Println("leftover", []int{int(overlapEnd), int(b2 - overlapEnd)})
					}
				}
			}

			if len(nextRanges) == 0 {
				fmt.Println("no overlap")

				nextRanges = append(nextRanges, currentRange)
			}
		}
	}

	lowestRangeStart := math.MaxInt64

	fmt.Println("nextRanges", len(nextRanges))
	for _, nextRange := range nextRanges {
		if nextRange[0] < lowestRangeStart {
			lowestRangeStart = nextRange[0]
		}
	}

	fmt.Println("lowestRangeStart", lowestRangeStart)
}

// Function to parse space-separated numbers from a string
func parseNumbers(input string) []int {
	var numbers []int
	for _, s := range strings.Fields(input) {
		num, err := strconv.Atoi(s)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func readMaps(scanner *bufio.Scanner) [][]Mapping {
	maps := make([][]Mapping, 0)
	for {
		oneMap := readMap(scanner)
		if len(oneMap) == 0 {
			break
		}
		maps = append(maps, oneMap)
	}

	return maps
}

// Function to read a map from the input
func readMap(scanner *bufio.Scanner) []Mapping {
	result := make([]Mapping, 0)

	// Read map lines
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		// Parse line into three numbers
		var dest, src, length int
		fmt.Sscanf(line, "%d %d %d", &dest, &src, &length)
		if length > 0 {
			result = append(result, Mapping{Source: src, Destination: dest, Range: length})
		}
	}

	return result
}
