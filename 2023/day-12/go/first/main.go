package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Galaxy struct {
	x int
	y int
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "input.txt", "The file path to load the data from")
	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
