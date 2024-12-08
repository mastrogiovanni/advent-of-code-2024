package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type Point struct {
	X int
	Y int
}

func Part1(graph *utility.CharGraph) {
	symbols := make(map[Point]bool)
	antennas := make(map[byte][]Point)
	for y := 0; y < graph.Height; y++ {
		for x := 0; x < graph.Width; x++ {
			symbol := graph.Get(x, y)
			if symbol == '.' {
				continue
			}
			position := Point{x, y}
			symbols[position] = true
			items, ok := antennas[symbol]
			if !ok {
				antennas[symbol] = []Point{position}
			} else {
				antennas[symbol] = append(items, position)
			}
		}
	}
	antinodes := make(map[Point]bool)
	for _, v := range antennas {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				a := v[i]
				b := v[j]
				dx := a.X - b.X
				dy := a.Y - b.Y
				antinodeA := Point{a.X + dx, a.Y + dy}
				antinodeB := Point{b.X - dx, b.Y - dy}
				if graph.In(antinodeA.X, antinodeA.Y) {
					antinodes[antinodeA] = true
				}
				if graph.In(antinodeB.X, antinodeB.Y) {
					antinodes[antinodeB] = true
				}
			}
		}
	}
	log.Println(len(antinodes))
}

func main() {
	Part1(utility.NewGraph("cmd/es8/test.txt"))
	Part1(utility.NewGraph("cmd/es8/input.txt"))
}
