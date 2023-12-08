package main

import (
	"bufio"
	"fmt"
	"log"
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
	seeds := parseNumbers(scanner.Text())

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
	for _, seed := range seeds {
		fmt.Println("seed", seed)
		soil := findMappingAndReturnDestination(seed, seedToSoilMap)
		fmt.Println("soil", soil)
		fertilizer := findMappingAndReturnDestination(soil, soilToFertilizerMap)
		fmt.Println("fertilizer", fertilizer)
		water := findMappingAndReturnDestination(fertilizer, fertilizerToWaterMap)
		fmt.Println("water", water)
		light := findMappingAndReturnDestination(water, waterToLightMap)
		fmt.Println("light", light)
		temperature := findMappingAndReturnDestination(light, lightToTemperatureMap)
		fmt.Println("temperature", temperature)
		humidity := findMappingAndReturnDestination(temperature, temperatureToHumidityMap)
		fmt.Println("humidity", humidity)
		location := findMappingAndReturnDestination(humidity, humidityToLocationMap)
		fmt.Println("location", location)

		if lowestLocation == -1 {
			lowestLocation = location
		}

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println("lowestLocation", lowestLocation)
}

func findMappingAndReturnDestination(input int, mappings []Mapping) int {
	for _, mapping := range mappings {
		// fmt.Println("")
		// fmt.Println("")
		// fmt.Println("")
		// fmt.Println("mapping", mapping)
		// fmt.Println("input", input)
		// fmt.Println("mapping.Source", mapping.Source)
		// fmt.Println("mapping.Range", mapping.Range)
		// fmt.Println("mapping.Destination", mapping.Destination)
		// fmt.Println("mapping.Source < input", mapping.Source < input)
		// fmt.Println("mapping.Source+mapping.Range > input", mapping.Source+mapping.Range > input)
		// fmt.Println("")
		// fmt.Println("")
		// fmt.Println("")
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
