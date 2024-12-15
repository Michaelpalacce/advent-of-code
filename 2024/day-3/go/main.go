package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const REGEX string = `mul\([\d]+,[\d]+\)`

func main() {
	// filePath := "debug.txt"
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	first(file)
}

func first(file *os.File) {
	r := regexp.MustCompile(`mul\([\d]+,[\d]+\)`)

	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		found := r.FindAllString(line, -1)

		for _, f := range found {
			mult := string(f)

			// Remove mul( and )
			mult = mult[4 : len(mult)-1]

			// Split by ,
			nums := strings.Split(mult, ",")

			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)
			}

			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}

			result += num1 * num2
		}
	}

	fmt.Println(result)
}
