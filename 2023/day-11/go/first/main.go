package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pipe struct {
	Symbol      string
	Connections []*Pipe
}

func main() {
	// filePath := "input.txt"
	filePath := "debug.txt"

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
