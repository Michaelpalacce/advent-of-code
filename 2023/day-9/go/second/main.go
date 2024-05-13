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

	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		result += calculate(line)
	}

	fmt.Println(result)
}

func calculate(line string) int {
	splitLine := strings.Split(line, " ")
	numbers := make([][]int, 1)

	fmt.Println(splitLine)

	for i := 0; i < len(splitLine); i++ {
		number, _ := strconv.Atoi(splitLine[i])
		numbers[0] = append(numbers[0], number)
	}

	for !isAllZero(numbers[len(numbers)-1]) {
		numbers = append(numbers, predict(numbers[len(numbers)-1]))
	}

	fillInPredictions(numbers)

	fmt.Println(numbers)

	return numbers[0][0]
}

func predict(numbers []int) []int {
	nextLine := make([]int, 0)

	for i := 0; i < len(numbers)-1; i++ {
		nextLine = append(nextLine, numbers[i+1]-numbers[i])
	}

	return nextLine
}

func isAllZero(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != 0 {
			return false
		}
	}

	return true
}

func fillInPredictions(numbers [][]int) {
	for i := len(numbers) - 1; i > 0; i-- {
		numbers[i-1] = append([]int{numbers[i-1][0] - numbers[i][0]}, numbers[i-1]...)
	}
}
