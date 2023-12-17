package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
