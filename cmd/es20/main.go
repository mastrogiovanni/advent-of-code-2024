package main

import (
	"log"
	"slices"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type GraphVisitor struct {
	Lenghts map[utility.Point]int
	Visited map[utility.Point]bool
	Parents map[utility.Point]utility.Point
}

func FindShortestPath(start, stop utility.Point, walls map[utility.Point]bool, width, height int) *GraphVisitor {
	graphVisitor := &GraphVisitor{
		Lenghts: make(map[utility.Point]int),
		Visited: make(map[utility.Point]bool),
		Parents: make(map[utility.Point]utility.Point),
	}
	queue := make([]utility.Point, 1)
	queue[0] = stop
	graphVisitor.Lenghts[stop] = 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if _, ok := graphVisitor.Visited[item]; ok {
			continue
		}
		graphVisitor.Visited[item] = true
		l := graphVisitor.Lenghts[item]
		for i := 0; i < 4; i++ {
			nx := item.X + utility.Directions[i][0]
			ny := item.Y + utility.Directions[i][1]
			if !(nx >= 0 && ny >= 0 && nx < width && ny < height) {
				// Outside graph
				continue
			}
			np := utility.Point{X: nx, Y: ny}
			if _, ok := walls[np]; ok {
				// Cell is blocked
				continue
			}
			lold, ok := graphVisitor.Lenghts[np]
			if !ok {
				graphVisitor.Lenghts[np] = l + 1
				graphVisitor.Parents[item] = np
			} else {
				if lold < l+1 {
					graphVisitor.Lenghts[np] = lold
				} else {
					graphVisitor.Lenghts[np] = l + 1
					graphVisitor.Parents[item] = np
				}
			}
			queue = append(queue, np)
		}
	}
	return graphVisitor
}

func GetShortestPath(stop utility.Point, graphVisitor *GraphVisitor) []utility.Point {
	path := make([]utility.Point, 0)
	node := stop
	for {
		path = append(path, node)
		v, ok := graphVisitor.Parents[node]
		if !ok {
			break
		}
		node = v
	}
	slices.Reverse(path)
	return path
}

func GetSavesUsingCuts(path []utility.Point, picosecondsRule int) map[int]int {
	saves := make(map[int]int)
	for i := 0; i < len(path)-2; i++ {
		for j := i + 2; j < len(path); j++ {
			taxiCabDistance := path[i].ManhattanDistance(path[j])
			if !(taxiCabDistance > 1 && taxiCabDistance <= picosecondsRule) {
				continue
			}
			save := j - i - taxiCabDistance
			if save <= 0 {
				continue
			}
			v, ok := saves[save]
			if ok {
				saves[save] = v + 1
			} else {
				saves[save] = 1
			}
		}
	}
	return saves
}

func GetTotalSaves(saves map[int]int, minPath int) int {
	sum := 0
	for save, count := range saves {
		if save >= minPath {
			sum += count
		}
	}
	return sum
}

func Evaluate(start, stop utility.Point, walls map[utility.Point]bool, width, height int, picosecondsRule int, minPath int) {
	graphVisitor := FindShortestPath(start, stop, walls, width, height)
	path := GetShortestPath(stop, graphVisitor)
	saves := GetSavesUsingCuts(path, picosecondsRule)
	sum := GetTotalSaves(saves, minPath)
	log.Println(sum)
}

func Resolve(fileName string, picosecondRule int, minPath int) {
	graph := utility.NewGraph(fileName)
	start := graph.Find('S')
	stop := graph.Find('E')
	g := make(map[utility.Point]bool)
	for x := 0; x < graph.Width; x++ {
		for y := 0; y < graph.Height; y++ {
			if graph.Get(x, y) == '#' {
				g[utility.Point{X: x, Y: y}] = true
			}
		}
	}

	Evaluate(start, stop, g, graph.Width, graph.Height, picosecondRule, minPath)
}

func main() {
	Resolve("cmd/es20/test.txt", 2, 0)
	Resolve("cmd/es20/input.txt", 2, 100)
	Resolve("cmd/es20/test.txt", 20, 0)
	Resolve("cmd/es20/input.txt", 20, 100)
}
