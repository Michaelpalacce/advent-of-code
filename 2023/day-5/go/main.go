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
	filePath := "debug.txt"

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

	// Read seed-to-soil map
	seedToSoilMap := readMap(scanner)
	soilToFertilizerMap := readMap(scanner)
	fertilizerToWaterMap := readMap(scanner)
	waterToLightMap := readMap(scanner)
	lightToTemperatureMap := readMap(scanner)
	temperatureToHumidityMap := readMap(scanner)
	humidityToLocationMap := readMap(scanner)
	lowestLocation := -1
	for _, seedRange := range seedRanges {
		fmt.Println("seed", seedRange)
		soil := findOverlapAndReturnDestination(seedRange, seedToSoilMap)
		fmt.Println("soil", soil)
		fertilizer := findOverlapAndReturnDestination(soil, soilToFertilizerMap)
		fmt.Println("fertilizer", fertilizer)
		water := findOverlapAndReturnDestination(fertilizer, fertilizerToWaterMap)
		fmt.Println("water", water)
		light := findOverlapAndReturnDestination(water, waterToLightMap)
		fmt.Println("light", light)
		temperature := findOverlapAndReturnDestination(light, lightToTemperatureMap)
		fmt.Println("temperature", temperature)
		humidity := findOverlapAndReturnDestination(temperature, temperatureToHumidityMap)
		fmt.Println("humidity", humidity)
		location := findOverlapAndReturnDestination(humidity, humidityToLocationMap)
		fmt.Println("location", location)

		if lowestLocation == -1 {
			lowestLocation = location[0]
		}

		if location[0] < lowestLocation {
			lowestLocation = location[0]
		}
	}

	fmt.Println("lowestLocation", lowestLocation)
}

// The only thing missing for the solutino is that I need to still return the unmatched ones :/
func findOverlapAndReturnDestination(input []int, mappings []Mapping) []int {
	for _, mapping := range mappings {
		a1, a2 := float64(mapping.Source), float64(mapping.Source+mapping.Range)
		b1, b2 := float64(input[0]), float64(input[0]+input[1])

		// Check for overlap
		overlapStart := math.Max(a1, b1)
		overlapEnd := math.Min(a2, b2)

		if overlapStart < overlapEnd {
			fmt.Println("overlap", overlapStart, overlapEnd)
			overlapLength := overlapEnd - overlapStart
			return []int{int(mapping.Destination +(int(overlapStart) - mapping.Source)), int(overlapLength)}
		}
	}

	return input
}

func findMappingAndReturnDestination(input int, mappings []Mapping) int {
	for _, mapping := range mappings {
		if mapping.Source < input && mapping.Source+mapping.Range > input {
			return mapping.Destination + (input - mapping.Source)
		}
	}
	return input
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
