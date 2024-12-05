package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	// filePath := "debug.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// first(file)
	second(file)
}

func first(file *os.File) {
	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if isSafe(convertToInts(line)) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func convertToInts(line string) []int {
	var err error
	levels := strings.Split(line, " ")
	intLevels := make([]int, len(levels))

	for i := 0; i < len(levels); i++ {
		intLevels[i], err = strconv.Atoi(strings.TrimSpace(levels[i]))
		if err != nil {
			log.Fatal(err)
		}
	}

	return intLevels
}

func isSafe(intLevels []int) bool {
	// 0 means decreasing
	direction := 0

	if intLevels[1] > intLevels[0] {
		direction = 1
	}

	for i := 1; i < len(intLevels); i++ {
		if intLevels[i] == intLevels[i-1] {
			return false
		}

		if direction == 0 && intLevels[i] > intLevels[i-1] {
			return false
		} else if direction == 1 && intLevels[i] < intLevels[i-1] {
			return false
		}

		distance := intLevels[i] - intLevels[i-1]
		if distance > 3 || distance < -3 {
			return false
		}
	}

	return true
}

func second(file *os.File) {
	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if isSafeWithToleration(convertToInts(line)) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSafeWithToleration(intLevels []int) bool {
	// 0 means decreasing
	direction := 0

	if intLevels[1] > intLevels[0] {
		direction = 1
	}

	isMarkedUnsafe := false

	for i := 1; i < len(intLevels); i++ {
		if intLevels[i] == intLevels[i-1] {
			isMarkedUnsafe = true
			break
		}

		if direction == 0 && intLevels[i] > intLevels[i-1] {
			isMarkedUnsafe = true
			break
		} else if direction == 1 && intLevels[i] < intLevels[i-1] {
			isMarkedUnsafe = true
			break
		}

		distance := intLevels[i] - intLevels[i-1]
		if distance > 3 || distance < -3 {
			isMarkedUnsafe = true
			break
		}
	}

	if isMarkedUnsafe {
		for i := 0; i < len(intLevels); i++ {
			newLevels := remove(intLevels, i)

			if isSafe(newLevels) {
				fmt.Printf("%v is save with %v\n", intLevels, newLevels)
				return true
			}
		}

		return false
	}

	return true
}

func remove(slice []int, i int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:i])
	copy(newSlice[i:], slice[i+1:])
	return newSlice
}
