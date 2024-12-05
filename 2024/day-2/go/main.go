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

	first(file)
	// second(file)
}

func first(file *os.File) {
	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if isSafe(line) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSafe(line string) bool {
	var err error
	levels := strings.Split(line, " ")
	intLevels := make([]int, len(levels))

	for i := 0; i < len(levels); i++ {
		intLevels[i], err = strconv.Atoi(strings.TrimSpace(levels[i]))
		if err != nil {
			log.Fatal(err)
		}
	}

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

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
