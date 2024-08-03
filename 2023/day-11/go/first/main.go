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
	cosmicExpansion := 999999
	var filePath string
	flag.StringVar(&filePath, "file", "input.txt", "The file path to load the data from")
	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	galaxies := make([]Galaxy, 0)
	width := 0
	height := 0

	for scanner.Scan() {
		line := scanner.Text()

		if width == 0 {
			width = len(line)
		}

		galaxies = append(galaxies, processLine(line, height)...)
		height++
	}

	emptyX := make([]int, 0)
	emptyY := make([]int, 0)

	for x := range width {
		hasGalaxy := false
		for _, galaxy := range galaxies {
			if galaxy.x == x {
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			fmt.Printf("No galaxy at x: %d\n", x)
			emptyX = append(emptyX, x)
		}
	}

	for y := range height {
		hasGalaxy := false
		for _, galaxy := range galaxies {
			if galaxy.y == y {
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			fmt.Printf("No galaxy at y: %d\n", y)
			emptyY = append(emptyY, y)
		}
	}

	for i := range len(galaxies) {
		galaxy := &galaxies[i]

		xIncrease := 0
		for _, x := range emptyX {
			if galaxy.x > x {
				xIncrease += cosmicExpansion
			}
		}

		yIncrease := 0
		for _, y := range emptyY {
			if galaxy.y > y {
				yIncrease += cosmicExpansion
			}
		}

		galaxy.x += xIncrease
		galaxy.y += yIncrease
	}

	shortestPathAccumulator := 0

	for len(galaxies) > 0 {
		galaxy := galaxies[0]
		galaxies = galaxies[1:]

		for _, otherGalaxy := range galaxies {
			shortestPathAccumulator += shortestPath(galaxy, otherGalaxy)
		}
	}

	fmt.Printf("Shortest path: %d\n", shortestPathAccumulator)
}

// processLine will find the galaxies in the given line and return them as a slice
func processLine(line string, y int) []Galaxy {
	galaxies := make([]Galaxy, 0)
	for index, char := range line {
		if string(char) == "#" {
			galaxies = append(galaxies, Galaxy{
				x: index,
				y: y,
			})
		}
	}
	return galaxies
}

func shortestPath(galaxy1 Galaxy, galaxy2 Galaxy) int {
	return Abs((galaxy1.x - galaxy2.x)) + Abs((galaxy1.y - galaxy2.y))
}

func Abs[T int | int8 | int16 | int32 | int64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
