package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BattleshipNode struct {
	xEnd    int
	numbers []int
	node    SimpleNode
}

type SymbolNode struct {
	symbol string
	node   SimpleNode
}

type SimpleNode struct {
	nodeType int
	x        int
	y        int
}

func (b *BattleshipNode) GetNodeType() int {
	return b.node.nodeType
}

func (b *BattleshipNode) GetX() int {
	return b.node.x
}

func (b *BattleshipNode) GetY() int {
	return b.node.y
}

func (b *BattleshipNode) getNum() int {
	// Convert each number to a string
	strNums := make([]string, len(b.numbers))
	for i, num := range b.numbers {
		strNums[i] = strconv.Itoa(num)
	}

	concatenatedStr := strings.Join(strNums, "")

	result, err := strconv.Atoi(concatenatedStr)
	if err != nil {
		log.Fatal("Could not convert number :0")
	}

	return result
}

func (s *SymbolNode) GetNodeType() int {
	return s.node.nodeType
}

func (s *SymbolNode) GetX() int {
	return s.node.x
}

func (s *SymbolNode) GetY() int {
	return s.node.y
}

func (n *SimpleNode) GetNodeType() int {
	return n.nodeType
}

func (n *SimpleNode) GetX() int {
	return n.x
}

func (n *SimpleNode) GetY() int {
	return n.y
}

type Node interface {
	GetX() int
	GetY() int
	GetNodeType() int
}

var (
	EMPTY_TYPE      = 1
	BATTLESHIP_TYPE = 2
	SYMBOL_TYPE     = 3
)

type Field struct {
	matrix [][]Node
}

func (f Field) getAllBattleShips() []*BattleshipNode {
	result := make([]*BattleshipNode, 0)

	for _, yLine := range f.matrix {
		for x := 0; x < len(yLine); x++ {
			xNode := yLine[x]
			if xNode.GetNodeType() == BATTLESHIP_TYPE {
				if v, ok := xNode.(*BattleshipNode); ok {
					result = append(result, v)
					x = v.xEnd
				} else {
					log.Fatal("error: Should Have been a battleship node :/")
				}
			}
		}
	}

	return result
}

func (f Field) HasSymbols(y int, xStart int, xEnd int) bool {

	fmt.Println(y, xStart, xEnd)
	if y <= 0 || y >= len(f.matrix) {
		return false
	}

	yLine := f.matrix[y]

	if xStart <= 0 {
		xStart = 0
	}

	if xEnd >= len(yLine)-1 {
		xEnd = len(yLine) - 1
	}
	fmt.Println(y, xStart, xEnd)

	nodes := yLine[xStart:xEnd+1]
	for _, node := range nodes {

        fmt.Println(node.GetNodeType())
        fmt.Println(node)
		if node.GetNodeType() == SYMBOL_TYPE {
			return true
		}
	}
	return false
}

func (f Field) HasSymbolsAround(battleship *BattleshipNode) bool {
	fmt.Println(battleship)
	batX := battleship.GetX()
	batY := battleship.GetY()
	batxEnd := battleship.xEnd

	fmt.Println(f.HasSymbols(batY, batX-1, batxEnd+1))

	return f.HasSymbols(batY, batX-1, batxEnd+1) || f.HasSymbols(batY-1, batX-1, batxEnd+1) || f.HasSymbols(batY+1, batX-1, batxEnd+1)
}

func populateFieldMatrix(line string, y int, field *Field) {
	yLine := make([]Node, len(line))
	for index, char := range line {
		sChar := string(char)
		if sChar == "." {
			yLine[index] = &SimpleNode{
				nodeType: EMPTY_TYPE,
				x:        index,
				y:        y,
			}
		} else if battleshipNumber, err := strconv.Atoi(sChar); err == nil {
			if index > 0 && yLine[index-1] != nil && yLine[index-1].GetNodeType() == BATTLESHIP_TYPE {
				prevNode := yLine[index-1]
				if v, ok := prevNode.(*BattleshipNode); ok {
					v.xEnd = index
					v.numbers = append(v.numbers, battleshipNumber)
					yLine[index] = prevNode
				} else {
					log.Fatal("error: Should Have been a battleship node :/")
				}
			} else {
				yLine[index] = &BattleshipNode{
					node: SimpleNode{
						x:        index,
						y:        y,
						nodeType: BATTLESHIP_TYPE,
					},
					numbers: []int{battleshipNumber},
					xEnd:    index,
				}
			}
		} else {
			// We naively assume it's a symbol :)
			yLine[index] = &SymbolNode{
				node: SimpleNode{
					x:        index,
					y:        y,
					nodeType: SYMBOL_TYPE,
				},
				symbol: sChar,
			}
		}
	}
	field.matrix = append(field.matrix, yLine)
}

func scanFileContents() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := Field{}
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		populateFieldMatrix(line, y, &field)
		y++
	}

	result := 0
	for _, battleship := range field.getAllBattleShips() {
		if ok := field.HasSymbolsAround(battleship); ok {
			fmt.Println("hasSymbolsAround", battleship.GetX(), battleship.GetY(), battleship.numbers)
			result += battleship.getNum()
		}
	}
	// debugNode := field.matrix[123][0]
	// if v, ok := debugNode.(*BattleshipNode); ok {
	// 	fmt.Println(field.HasSymbolsAround(v))
	// } else {
	// 	log.Fatal("error: Should Have been a battleship node :/")
	// }

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	scanFileContents()
}
