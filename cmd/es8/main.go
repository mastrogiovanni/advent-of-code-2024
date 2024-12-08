package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type Point struct {
	X int
	Y int
}

type AntennasAndSymbols struct {
	Antennas map[byte][]Point
	Symbols  map[Point]bool
}

func GetAntennasAndSymbols(graph *utility.CharGraph) *AntennasAndSymbols {
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
	return &AntennasAndSymbols{
		Antennas: antennas,
		Symbols:  symbols,
	}
}

func Part2(graph *utility.CharGraph) {
	antennasAndSymbols := GetAntennasAndSymbols(graph)
	antinodes := make(map[Point]bool)
	for _, v := range antennasAndSymbols.Antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				a := v[i]
				b := v[j]
				dx := a.X - b.X
				dy := a.Y - b.Y
				for z := 0; ; z++ {
					antinodeResonant := Point{a.X + z*dx, a.Y + z*dy}
					if graph.In(antinodeResonant.X, antinodeResonant.Y) {
						antinodes[antinodeResonant] = true
						graph.Set(antinodeResonant.X, antinodeResonant.Y, '#')
					} else {
						break
					}
				}
				for z := 0; ; z++ {
					antinodeResonant := Point{b.X - z*dx, b.Y - z*dy}
					if graph.In(antinodeResonant.X, antinodeResonant.Y) {
						antinodes[antinodeResonant] = true
						graph.Set(antinodeResonant.X, antinodeResonant.Y, '#')
					} else {
						break
					}
				}

			}
		}
	}
	// graph.Dump()
	log.Println(len(antinodes))
}

func Part1(graph *utility.CharGraph) {
	antennasAndSymbols := GetAntennasAndSymbols(graph)
	antinodes := make(map[Point]bool)
	for _, v := range antennasAndSymbols.Antennas {
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
	Part2(utility.NewGraph("cmd/es8/test.txt"))
	Part2(utility.NewGraph("cmd/es8/input.txt"))
}
