package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func GetDirectionsFromLine(line string) (directions []rune) {
	for _, char := range line {
		directions = append(directions, char)
	}

	return directions
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func trimLastRune(s string) string {
	_, i := utf8.DecodeLastRuneInString(s)
	return s[:len(s)-i]
}

// constructorMap is a map of node values to node pointers, used to
// keep track of nodes that have already been created
var constructorMap = make(map[string]*Node)

func NewNode(value string) *Node {
	if node, ok := constructorMap[value]; ok {
		return node
	}

	node := &Node{
		Value: value,
	}

	constructorMap[value] = node

	return node
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  string
}

func (n *Node) AddChildren(nodeLine string) {
	nodeParts := strings.Split(nodeLine, "=")
	nodeChildren := strings.Split(strings.TrimSpace(nodeParts[1]), ",")
	leftValue := trimFirstRune(strings.TrimSpace(nodeChildren[0]))
	rightValue := trimLastRune(strings.TrimSpace(nodeChildren[1]))

	n.Left = NewNode(leftValue)
	n.Right = NewNode(rightValue)
	n.Left.Parent = n
	n.Right.Parent = n
}

func repeatSlice(original []rune, times int) []rune {
	var repeated []rune
	for i := 0; i < times; i++ {
		repeated = append(repeated, original...)
	}
	return repeated
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	directions := GetDirectionsFromLine(scanner.Text())

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		nodeParts := strings.Split(line, "=")
		nodeValue := strings.TrimSpace(nodeParts[0])

		node := NewNode(nodeValue)

		node.AddChildren(line)

		fmt.Printf("Node: %s, Left: %s, Right: %s\n", node.Value, node.Left.Value, node.Right.Value)

	}

	currentNode := constructorMap["AAA"]

	steps := 0

endWalk:
	for {
		for _, direction := range directions {
			fmt.Printf("Current node: %s, Current Direction: %c\n", currentNode.Value, direction)

			steps++

			if direction == 'L' {
				currentNode = currentNode.Left
			} else {
				currentNode = currentNode.Right
			}

			if currentNode.Value == "ZZZ" {
				fmt.Println("Reached the end of the tree")
				break endWalk
			}
		}
	}

	fmt.Println("Steps: ", steps)
}
