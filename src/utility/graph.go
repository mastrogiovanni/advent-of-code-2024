package utility

import "fmt"

type CharGraph struct {
	Width  int
	Height int
	Rows   []string
}

type Point struct {
	X int
	Y int
}

var Directions = [][]int{
	{0, -1}, // 0: north
	{1, 0},  // 1: east
	{0, 1},  // 2: south
	{-1, 0}, // 3: west
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
