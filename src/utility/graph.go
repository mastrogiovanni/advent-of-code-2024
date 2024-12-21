package utility

import (
	"bufio"
	"fmt"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type CharGraph struct {
	Width  int
	Height int
	Rows   []string
}

type Point struct {
	X int
	Y int
}

func (p1 Point) ManhattanDistance(p2 Point) int {
	return max(p1.X, p2.X) - min(p1.X, p2.X) + max(p1.Y, p2.Y) - min(p1.Y, p2.Y)
}

func (p Point) Move(d Direction) Point {
	return Point{
		X: p.X + Directions[d][0],
		Y: p.Y + Directions[d][1],
	}
}

var Directions = [][]int{
	{0, -1}, // 0: north 	^
	{1, 0},  // 1: east 	>
	{0, 1},  // 2: south	v
	{-1, 0}, // 3: west		<
}

func NewGraphFromScanner(scanner *bufio.Scanner) *CharGraph {
	var fileLines []string
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}
		fileLines = append(fileLines, row)
	}
	width := len(fileLines[0])
	height := len(fileLines)
	return &CharGraph{
		Rows:   fileLines,
		Width:  width,
		Height: height,
	}
}

func NewGraph(fileName string) *CharGraph {
	rows := FileLines(fileName)
	width := len(rows[0])
	height := len(rows)
	return &CharGraph{
		Rows:   rows,
		Width:  width,
		Height: height,
	}
}

func (g *CharGraph) Find(symbol byte) Point {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Get(x, y) == symbol {
				return Point{X: x, Y: y}
			}
		}
	}
	return Point{-1, -1}
}

func (g *CharGraph) Dump() {
	fmt.Println()
	for _, row := range g.Rows {
		fmt.Println(row)
	}
}

func (g *CharGraph) Get(x, y int) byte {
	return g.Rows[y][x]
}

func (g *CharGraph) Set(x, y int, c byte) {
	row := g.Rows[y]
	bytes := []byte(row)
	bytes[x] = c
	g.Rows[y] = string(bytes)
}

func (g *CharGraph) In(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}
