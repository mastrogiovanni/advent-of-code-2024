package utility

import "fmt"

type CharGraph struct {
	Width  int
	Height int
	Rows   []string
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
