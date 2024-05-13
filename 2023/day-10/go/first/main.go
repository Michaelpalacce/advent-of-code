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
	filePath := "input.txt"
	// filePath := "debug.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]*Pipe, 0)

	for scanner.Scan() {
		line := scanner.Text()
		populateGridLine(&grid, line)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("Y: %d, X: %d\n", y, x)
			findConnections(grid, x, y)
		}
	}

	startPipe := determineStart(grid)

	for _, line := range grid {
		fmt.Println("---------------------------------------")
		for _, pipe := range line {
			fmt.Printf("Pipe: %+v\n", pipe)
			for _, connection := range pipe.Connections {
				fmt.Print(connection.Symbol)
			}
			fmt.Println()
		}
	}

	fmt.Printf("Furthest pipe: %d\n", findFurthestPipe(startPipe.Connections, []*Pipe{startPipe}, 1))
}

func populateGridLine(grid *[][]*Pipe, line string) {
	yLine := make([]*Pipe, len(line))
	for index, char := range line {
		sChar := string(char)

		yLine[index] = &Pipe{
			Symbol: sChar,
		}
	}

	*grid = append(*grid, yLine)
}

func determineStart(grid [][]*Pipe) *Pipe {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].Symbol == "S" {
				// Hacky wacky, since data is static, don't want to bother actually finding these
				// Input
				grid[y][x].Connections = []*Pipe{grid[y-1][x], grid[y+1][x]}

				// Debug
				// grid[y][x].Connections = []*Pipe{grid[y+1][x], grid[y][x+1]}
				return grid[y][x]
			}
		}
	}
	return nil
}

func findConnections(grid [][]*Pipe, x int, y int) {
	connections := make([]*Pipe, 0)

	pipe := grid[y][x]

	switch pipe.Symbol {
	case "|":
		if y > 0 {
			connections = append(connections, grid[y-1][x])
		}
		if y < len(grid)-1 {
			connections = append(connections, grid[y+1][x])
		}
	case "-":
		if x > 0 {
			connections = append(connections, grid[y][x-1])
		}
		if x < len(grid[y])-1 {
			connections = append(connections, grid[y][x+1])
		}
	case "7":
		if x > 0 {
			connections = append(connections, grid[y][x-1])
		}
		if y < len(grid)-1 {
			fmt.Printf("BAD: Y: %d, X: %d\n", y, x)
			connections = append(connections, grid[y+1][x])
		}
	case "L":
		if y > 0 {
			connections = append(connections, grid[y-1][x])
		}
		if x < len(grid[y])-1 {
			connections = append(connections, grid[y][x+1])
		}
	case "S":
		//
	case "F":
		if x < len(grid[y])-1 {
			connections = append(connections, grid[y][x+1])
		}
		if y < len(grid)-1 {
			connections = append(connections, grid[y+1][x])
		}
	case "J":
		if y > 0 {
			connections = append(connections, grid[y-1][x])
		}
		if x > 0 {
			connections = append(connections, grid[y][x-1])
		}
	case ".":
		// Ground
	}

	grid[y][x].Connections = connections
}

func findFurthestPipe(entrances []*Pipe, previousPipes []*Pipe, counter int) int {
	nextPipes := make([]*Pipe, 0)

	for _, entrance := range entrances {
		fmt.Printf("Entrance: %+v\n", entrance)
	}

	for _, previousPipe := range previousPipes {
		fmt.Printf("Previous pipe: %+v\n", previousPipe)
	}

	for _, entrance := range entrances {
	con:
		for _, connection := range entrance.Connections {
			isPrevious := false
			for _, previousPipe := range previousPipes {
				if connection == previousPipe {
					isPrevious = true
				}
			}

			if !isPrevious {
				fmt.Printf("Connection: %+v\n", connection)
				nextPipes = append(nextPipes, connection)
				break con
			}
		}
	}

	counter++

	fmt.Printf("Next pipes: %+v\n", nextPipes)

	if nextPipes[0] == nextPipes[1] {
		return counter
	}

	return findFurthestPipe(nextPipes, entrances, counter)
}
