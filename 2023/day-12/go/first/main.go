package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	Rows []Row
}

type Row struct {
	Positions   []Position
	Contingency []int
}

const (
	Onsen   = "."
	Broken  = "#"
	Unknown = "?"
)

type Position struct {
	X     int
	Y     int
	State string
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
	height := 0
	grid := Grid{
		Rows: make([]Row, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		grid.Rows = append(grid.Rows, processLine(line, height))
		height++
	}

	for _, row := range grid.Rows {
		fmt.Printf("Now working on row: %v\n", row)

		break
	}
}

func processLine(line string, height int) Row {
	positions := make([]Position, 0)

	lineParts := strings.Split(line, " ")

	if len(lineParts) != 2 {
		log.Fatalf("Invalid line: %s\n", line)
	}

	grid := lineParts[0]

	for index, char := range grid {
		sChar := string(char)

		positions = append(positions, Position{
			X:     index,
			Y:     height,
			State: sChar,
		})
	}

	contingencyParts := strings.Split(lineParts[1], ",")
	contingency := make([]int, 0)

	for _, part := range contingencyParts {
		partInt, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Invalid contingency part: %s\n", part)
		}

		contingency = append(contingency, partInt)
	}

	return Row{
		Positions:   positions,
		Contingency: contingency,
	}
}
