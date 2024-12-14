package main

import (
	"log"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type PointAndDirection struct {
	X         int
	Y         int
	Direction int
}

func HasLoop(graph *utility.CharGraph, posX, posY, direction int) bool {
	visited := make(map[PointAndDirection]bool)
	// count := 0
	for {
		key := PointAndDirection{X: posX, Y: posY, Direction: direction}
		_, ok := visited[key]
		if ok {
			// I was already in the same position with the same direction: is a loop
			return true
		}
		visited[key] = true
		dx := utility.Directions[direction][0]
		dy := utility.Directions[direction][1]
		nextX := posX + dx
		nextY := posY + dy
		if !graph.In(nextX, nextY) {
			// Exiting from the board: no loop
			return false
		}
		nextChar := graph.Get(nextX, nextY)
		if nextChar == '#' || nextChar == 'O' {
			// Next is a wall: change direction
			direction = (direction + 1) % 4
			// count++
			// if count > 1 {
			// 	log.Println("2", visited)
			// 	return false
			// }
		} else {
			posX = nextX
			posY = nextY
			// count = 0
			// key = fmt.Sprintf("%d,%d,%d", posX, posY, direction)
			// visited[key] = true
		}
	}
}

func Visited(graph *utility.CharGraph, posX, posY int) map[utility.Point]bool {
	visited := make(map[utility.Point]bool)
	direction := 0
	for {
		visited[utility.Point{X: posX, Y: posY}] = true
		dx := utility.Directions[direction][0]
		dy := utility.Directions[direction][1]
		if !graph.In(posX+dx, posY+dy) {
			break
		}
		symbol := graph.Get(posX+dx, posY+dy)
		if symbol == '#' || symbol == 'O' {
			direction = (direction + 1) % 4
		} else {
			posX = posX + dx
			posY = posY + dy
		}
	}
	return visited
}

func Candidates(graph *utility.CharGraph) map[utility.Point]bool {
	posX := -1
	posY := -1
	for y := 0; y < graph.Height; y++ {
		i := strings.Index(graph.Rows[y], "^")
		if i >= 0 {
			posX = i
			posY = y
		}
	}
	// startX := posX
	// startY := posY
	candidates := make(map[utility.Point]bool)
	direction := 0
	// count := 0
	for {
		dx := utility.Directions[direction][0]
		dy := utility.Directions[direction][1]
		testX := posX + dx
		testY := posY + dy
		if !graph.In(testX, testY) {
			break
		}
		nextChar := graph.Get(testX, testY)
		if nextChar == '#' || nextChar == 'O' {
			direction = (direction + 1) % 4
			// count = 0
		} else {
			posX = testX
			posY = testY
			// count++
			candidates[utility.Point{X: testX, Y: testY}] = true
			// if startX != testX && startY != testY /*&& count > 1*/ {
			// 	// 	nextX := testX + Directions[(direction+1)%4][0]
			// 	// 	nextY := testY + Directions[(direction+1)%4][1]
			// 	// 	nextNextChar := graph.Get(nextX, nextY)
			// 	// 	if nextNextChar != '#' && nextNextChar != 'O' {
			// 	// 	}
			// }
		}
	}
	return candidates
}

func Part2(graph *utility.CharGraph) {
	posX := -1
	posY := -1
	for y := 0; y < graph.Height; y++ {
		i := strings.Index(graph.Rows[y], "^")
		if i >= 0 {
			posX = i
			posY = y
		}
	}
	candidates := Candidates(graph)
	count := 0
	for k := range candidates {
		x := k.X
		y := k.Y
		if x == posX && y == posY {
			continue
		}
		if graph.Get(x, y) == '#' {
			continue
		}
		symbol := graph.Get(x, y)
		graph.Set(x, y, 'O')
		if HasLoop(graph, posX, posY, 0) {
			// graph.Dump()
			count = count + 1
		}
		graph.Set(x, y, symbol)
	}
	log.Println(count)
}

func Part1(graph *utility.CharGraph) {
	posX := -1
	posY := -1
	for y := 0; y < graph.Height; y++ {
		i := strings.Index(graph.Rows[y], "^")
		if i >= 0 {
			posX = i
			posY = y
		}
	}
	visited := Visited(graph, posX, posY)
	log.Println(len(visited))
}

func main() {
	Part1(utility.NewGraph("cmd/es6/test.txt"))
	Part1(utility.NewGraph("cmd/es6/input.txt"))
	Part2(utility.NewGraph("cmd/es6/test.txt"))
	Part2(utility.NewGraph("cmd/es6/input.txt"))
}
