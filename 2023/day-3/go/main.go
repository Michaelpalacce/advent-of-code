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

func (f Field) getAllSymbols() []*SymbolNode {
	result := make([]*SymbolNode, 0)

	for _, yLine := range f.matrix {
		for x := 0; x < len(yLine); x++ {
			xNode := yLine[x]
			if xNode.GetNodeType() == SYMBOL_TYPE {
				if v, ok := xNode.(*SymbolNode); ok {
					result = append(result, v)
				} else {
					log.Fatal("error: Should Have been a symbol node :/")
				}
			}
		}
	}

	return result
}

func (f Field) GetBattleshipsInRange(y int, xStart int, xEnd int) []*BattleshipNode {
	if y < 0 || y >= len(f.matrix) {
		return []*BattleshipNode{}
	}
	fmt.Printf("y: %d, xS: %d, xE: %d\n", y, xStart, xEnd)

	yLine := f.matrix[y]

	if xStart <= 0 {
		xStart = 0
	}

	if xEnd >= len(yLine)-1 {
		xEnd = len(yLine) - 1
	}

	results := make([]*BattleshipNode, 0)
	for x := xStart; x < xEnd+1; x++ {
		xNode := yLine[x]
		if xNode.GetNodeType() == BATTLESHIP_TYPE {
			if v, ok := xNode.(*BattleshipNode); ok {
				results = append(results, v)
				x = v.xEnd
			} else {
				log.Fatal("error: Should Have been a symbol node :/")
			}
		}
	}
	return results
}

func (f Field) GetBattleshipsAround(symbol SymbolNode) []*BattleshipNode {
	batX := symbol.GetX()
	batY := symbol.GetY()

	sameLineBattleships := f.GetBattleshipsInRange(batY, batX-1, batX+1)
	upperLineBattleships := f.GetBattleshipsInRange(batY-1, batX-1, batX+1)
	belowLineBattleships := f.GetBattleshipsInRange(batY+1, batX-1, batX+1)
	fmt.Println(upperLineBattleships)

	return append(append(sameLineBattleships, belowLineBattleships...), upperLineBattleships...)
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
	for _, symbol := range field.getAllSymbols() {
		fmt.Println(symbol)
		battleships := field.GetBattleshipsAround(*symbol)

		if len(battleships) == 2 {
			for _, bat := range battleships {
				fmt.Println(bat)
			}
			result += battleships[0].getNum() * battleships[1].getNum()
		}
	}

	// debugNode := field.matrix[1][12]
	// fmt.Println(debugNode.GetNodeType())
	// fmt.Println(debugNode)
	// if v, ok := debugNode.(*SymbolNode); ok {
	// 	field.GetBattleshipsAround(*v)
	// } else {
	// 	log.Fatal("error: Should Have been a symbol node :/")
	// }

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	scanFileContents()
}
