package main

import (
	"log"
	"slices"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func FindShortestPath(start, stop utility.Point, walls map[utility.Point]bool, width, height int, parents map[utility.Point]utility.Point) int {
	lenghts := make(map[utility.Point]int)
	visited := make(map[utility.Point]bool)
	queue := make([]utility.Point, 1)
	queue[0] = stop
	lenghts[stop] = 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if _, ok := visited[item]; ok {
			continue
		}
		visited[item] = true
		l := lenghts[item]
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
			lold, ok := lenghts[np]
			if !ok {
				lenghts[np] = l + 1
				parents[item] = np
			} else {
				if lold < l+1 {
					lenghts[np] = lold
				} else {
					lenghts[np] = l + 1
					parents[item] = np
				}
			}
			queue = append(queue, np)
		}
	}

	l, ok := lenghts[start]
	if ok {
		return l
	} else {
		return -1
	}
}

func Evaluate(start, stop utility.Point, walls map[utility.Point]bool, width, height int, picosecondsRule int, minPath int) {
	parents := make(map[utility.Point]utility.Point)
	FindShortestPath(start, stop, walls, width, height, parents)
	path := make([]utility.Point, 0)
	node := stop
	for {
		path = append(path, node)
		v, ok := parents[node]
		if !ok {
			break
		}
		node = v
	}
	slices.Reverse(path)
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
	sum := 0
	for save, count := range saves {
		if save >= minPath {
			sum += count
		}
	}
	// log.Println(saves)
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
	Resolve("cmd/es20/input.txt", 20, 50)

}
