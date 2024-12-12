package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type AreaAndPerimeter struct {
	Area      int
	Perimeter int
	Sides     int
}

var Directions = [][]int{
	{0, -1}, // 0: north
	{1, 0},  // 1: east
	{0, 1},  // 2: south
	{-1, 0}, // 3: west
}

func Expand(graph *utility.CharGraph, point utility.Point, globalVisited map[utility.Point]bool) AreaAndPerimeter {
	result := AreaAndPerimeter{Area: 0, Perimeter: 0}
	symbol := graph.Get(point.X, point.Y)
	visited := make(map[utility.Point]bool)
	for toVisit := []utility.Point{point}; len(toVisit) > 0; {
		point := toVisit[0]
		toVisit = toVisit[1:]
		if _, ok := visited[point]; ok {
			continue
		}
		visited[point] = true
		globalVisited[point] = true
		result.Area += 1
		result.Perimeter += 4
		for direction := 0; direction < 4; direction++ {
			if !graph.In(point.X+Directions[direction][0], point.Y+Directions[direction][1]) {
				continue
			}
			adjacent := utility.Point{X: point.X + Directions[direction][0], Y: point.Y + Directions[direction][1]}
			if graph.Get(adjacent.X, adjacent.Y) != symbol {
				continue
			}
			if _, ok := visited[adjacent]; ok {
				result.Perimeter -= 2
				continue
			}
			toVisit = append(toVisit, adjacent)
		}
	}
	result.Sides = result.Perimeter - GetSidesAdjustment(graph, visited)
	return result
}

func GetSidesAdjustment(graph *utility.CharGraph, visited map[utility.Point]bool) int {
	toRemove := 0
	for point := range visited {
		for _, direction := range []int{0, 1} {
			adjacent := utility.Point{X: point.X + Directions[direction][0], Y: point.Y + Directions[direction][1]}
			if !graph.In(point.X, point.Y) || !graph.In(adjacent.X, adjacent.Y) {
				continue
			}
			_, ok1 := visited[point]
			_, ok2 := visited[adjacent]
			if !ok1 || !ok2 {
				continue
			}
			if direction == 0 {
				// check if in direction 1 and 3 there are not visited
				pointCheck1 := utility.Point{X: point.X + Directions[1][0], Y: point.Y + Directions[1][1]}
				adjacentCheck1 := utility.Point{X: adjacent.X + Directions[1][0], Y: adjacent.Y + Directions[1][1]}
				_, ok3 := visited[pointCheck1]
				_, ok4 := visited[adjacentCheck1]
				if !ok3 && !ok4 {
					toRemove += 1
				}

				pointCheck3 := utility.Point{X: point.X + Directions[3][0], Y: point.Y + Directions[3][1]}
				adjacentCheck3 := utility.Point{X: adjacent.X + Directions[3][0], Y: adjacent.Y + Directions[3][1]}
				_, ok5 := visited[pointCheck3]
				_, ok6 := visited[adjacentCheck3]
				if !ok5 && !ok6 {
					toRemove += 1
				}
			} else if direction == 1 {
				// check if in direction 2 and 0 there are not visited
				pointCheck2 := utility.Point{X: point.X + Directions[2][0], Y: point.Y + Directions[2][1]}
				adjacentCheck2 := utility.Point{X: adjacent.X + Directions[2][0], Y: adjacent.Y + Directions[2][1]}
				_, ok3 := visited[pointCheck2]
				_, ok4 := visited[adjacentCheck2]
				if !ok3 && !ok4 {
					toRemove += 1
				}

				pointCheck0 := utility.Point{X: point.X + Directions[0][0], Y: point.Y + Directions[0][1]}
				adjacentCheck0 := utility.Point{X: adjacent.X + Directions[0][0], Y: adjacent.Y + Directions[0][1]}
				_, ok5 := visited[pointCheck0]
				_, ok6 := visited[adjacentCheck0]
				if !ok5 && !ok6 {
					toRemove += 1
				}
			}
		}
	}
	return toRemove
}

func Resolver(fileName string) {
	graph := utility.NewGraph(fileName)
	total := 0
	totalDiscounted := 0
	globalVisited := make(map[utility.Point]bool)
	for x := 0; x < graph.Width; x++ {
		for y := 0; y < graph.Height; y++ {
			point := utility.Point{X: x, Y: y}
			if _, ok := globalVisited[point]; ok {
				continue
			}
			aileAreaAndPerimeter := Expand(graph, point, globalVisited)
			total += aileAreaAndPerimeter.Area * aileAreaAndPerimeter.Perimeter
			totalDiscounted += aileAreaAndPerimeter.Area * aileAreaAndPerimeter.Sides
		}
	}
	log.Println(fileName)
	log.Println("Total", total)
	log.Println("Discounted", totalDiscounted)
}

func main() {
	Resolver("cmd/es12/test1.txt")
	Resolver("cmd/es12/test2.txt")
	Resolver("cmd/es12/test3.txt")
	Resolver("cmd/es12/test4.txt")
	Resolver("cmd/es12/test5.txt")
	Resolver("cmd/es12/input.txt")
}
