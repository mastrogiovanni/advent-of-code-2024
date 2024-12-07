package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

// dx = 1, dy = 0
// dx = 0, dy = 1
// dx = 1, dy = 1
func SearchD(graph *utility.CharGraph, x, y, dx, dy int) int {
	if x+dx*3 >= graph.Width || x+dx*3 < 0 {
		return 0
	}
	if y+dy*3 >= graph.Height || y+dy*3 < 0 {
		return 0
	}
	if graph.Get(x+dx, y+dy) != 'M' {
		return 0
	}
	if graph.Get(x+2*dx, y+2*dy) != 'A' {
		return 0
	}
	if graph.Get(x+3*dx, y+3*dy) != 'S' {
		return 0
	}
	return 1
}

func Check(graph *utility.CharGraph, x, y, dx, dy int) int {
	if !graph.In(x-dx, 0) || !graph.In(x+dx, 0) || !graph.In(y-dy, 0) || !graph.In(y+dy, 0) {
		return 0
	}
	if graph.Get(x+dx, y+dy) == 'M' && graph.Get(x-dx, y-dy) == 'S' {
		return 1
	}
	if graph.Get(x+dx, y+dy) == 'S' && graph.Get(x-dx, y-dy) == 'M' {
		return 1
	}
	return 0
}

func SearchD2(graph *utility.CharGraph, x, y int) int {
	count := 0
	d := Check(graph, x, y, 1, 1)
	d += Check(graph, x, y, 1, -1)
	if d == 2 {
		count = count + 1
	}
	return count
}

func Search1(graph *utility.CharGraph, x, y int) int {
	count := 0
	count += SearchD(graph, x, y, 1, 0)
	count += SearchD(graph, x, y, 0, 1)
	count += SearchD(graph, x, y, 1, 1)
	count += SearchD(graph, x, y, -1, 1)
	count += SearchD(graph, x, y, -1, 0)
	count += SearchD(graph, x, y, 0, -1)
	count += SearchD(graph, x, y, -1, -1)
	count += SearchD(graph, x, y, 1, -1)
	return count
}

func Part1(graph *utility.CharGraph) {
	count := 0
	for x := range graph.Width {
		for y := range graph.Height {
			if graph.Get(x, y) == 'X' {
				count += Search1(graph, x, y)
			}
		}
	}
	log.Println(count)
}

func Part2(graph *utility.CharGraph) {
	count := 0
	for x := range graph.Width {
		for y := range graph.Height {
			if graph.Get(x, y) == 'A' {
				count += SearchD2(graph, x, y)
			}
		}
	}
	log.Println(count)
}

func main() {
	Part1(utility.NewGraph("cmd/es4/test.txt"))
	Part1(utility.NewGraph("cmd/es4/input.txt"))
	Part2(utility.NewGraph("cmd/es4/test.txt"))
	Part2(utility.NewGraph("cmd/es4/input.txt"))
}
