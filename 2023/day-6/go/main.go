package main

import (
	"bufio"
	"log"
	"os"
)

// -18
func main() {
	filePath := "debug.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
}
