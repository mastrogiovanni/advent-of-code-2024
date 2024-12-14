package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

var NextChar = map[byte]byte{
	'0': '1',
	'1': '2',
	'2': '3',
	'3': '4',
	'4': '5',
	'5': '6',
	'6': '7',
	'7': '8',
	'8': '9',
}

func Traverse(point utility.Point, graph *utility.CharGraph, targets map[utility.Point]int) {
	currentChar := graph.Get(point.X, point.Y)
	if currentChar == '9' {
		v, ok := targets[point]
		if ok {
			targets[point] = v + 1
		} else {
			targets[point] = 1
		}
		return
	}
	nextChar := NextChar[currentChar]
	for _, direction := range utility.Directions {
		pX := point.X + direction[0]
		pY := point.Y + direction[1]
		if graph.In(pX, pY) && graph.Get(pX, pY) == nextChar {
			Traverse(utility.Point{
				X: pX,
				Y: pY,
			}, graph, targets)
		}
	}
}

func Part1(graph *utility.CharGraph) {
	count := 0
	for x := 0; x < graph.Width; x++ {
		for y := 0; y < graph.Width; y++ {
			if graph.Get(x, y) == '0' {
				targets := make(map[utility.Point]int)
				Traverse(utility.Point{X: x, Y: y}, graph, targets)
				count = count + len(targets)
			}
		}
	}
	log.Println(count)
}

func Part2(graph *utility.CharGraph) {
	count := 0
	for x := 0; x < graph.Width; x++ {
		for y := 0; y < graph.Width; y++ {
			if graph.Get(x, y) == '0' {
				targets := make(map[utility.Point]int)
				Traverse(utility.Point{X: x, Y: y}, graph, targets)
				for _, v := range targets {
					count = count + v
				}
			}
		}
	}
	log.Println(count)
}

func main() {
	Part1(utility.NewGraph("cmd/es10/test.txt"))
	Part1(utility.NewGraph("cmd/es10/input.txt"))
	Part2(utility.NewGraph("cmd/es10/test.txt"))
	Part2(utility.NewGraph("cmd/es10/input.txt"))
}
