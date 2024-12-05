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

func main() {
	// filePath := "input.txt"
	filePath := "debug.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	first(file)
	// first(file)
}

func first(file *os.File) {
	scanner := bufio.NewScanner(file)

	var leftSide []int
	var rightSide []int

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, "   ")

		leftNumber, err := strconv.Atoi(lineParts[0])
		if err != nil {
			log.Fatal(err)
		}

		leftSide = append(leftSide, leftNumber)

		rightNumber, err := strconv.Atoi(lineParts[1])
		if err != nil {
			log.Fatal(err)
		}

		rightSide = append(rightSide, rightNumber)
	}

	leftIndexes := getIndexedArray(leftSide)
	rightIndexes := getIndexedArray(rightSide)

	fmt.Println(leftSide)
	fmt.Println(leftIndexes)

	fmt.Println("============================")
	fmt.Println(rightSide)
	fmt.Println(rightIndexes)

	size := 0.0

	for i := 0; i < len(leftIndexes); i++ {
		fmt.Println(leftIndexes[i], rightIndexes[i])
		size += math.Abs(float64(leftIndexes[i] - rightIndexes[i]))
	}

	fmt.Println(int(size))
}

func getIndexedArray(arr []int) []int {
	indexedArr := make([]int, 0, len(arr))
	sortedArr := sortArray(arr)
	indexMap := make(map[int]bool)

	for _, currentValue := range sortedArr {
		for j, val := range arr {
			if val == currentValue && !indexMap[j] {
				indexedArr = append(indexedArr, j)
				indexMap[j] = true
				break
			}
		}
	}

	return indexedArr
}

func sortArray(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	for i := 0; i < len(newArr); i++ {
		for j := 0; j < len(newArr); j++ {
			if newArr[i] < newArr[j] {
				newArr[i], newArr[j] = newArr[j], newArr[i]
			}
		}
	}
	return newArr
}

func second(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
