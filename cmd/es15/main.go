package main

import (
	"fmt"
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Step(graph *utility.CharGraph, position utility.Point, direction int) utility.Point {
	// => @OO. goes in the direction looking for the first space or end or wall
	for i := 0; ; i++ {
		nextX := position.X + utility.Directions[direction][0]*i
		nextY := position.Y + utility.Directions[direction][1]*i
		if !graph.In(nextX, nextY) {
			// Don't move
			return position
		}
		nextSymbol := graph.Get(nextX, nextY)
		if nextSymbol == '#' {
			// Don't move
			return position
		}
		if nextSymbol == '.' {
			for j := i; j > 0; j = j - 1 {
				// Move up to here all the others
				lastX := position.X + utility.Directions[direction][0]*j
				lastY := position.Y + utility.Directions[direction][1]*j
				prevX := position.X + utility.Directions[direction][0]*(j-1)
				prevY := position.Y + utility.Directions[direction][1]*(j-1)
				graph.Set(lastX, lastY, graph.Get(prevX, prevY))
			}
			graph.Set(position.X, position.Y, '.')
			return utility.Point{
				X: position.X + utility.Directions[direction][0],
				Y: position.Y + utility.Directions[direction][1],
			}
		}
	}
}

// true means that was pushed
func Push(graph *utility.CharGraph, position utility.Point, direction int, dry bool) bool {

	// 0: north 	^
	// 1: east 		>
	// 2: south		v
	// 3: west		<

	if direction == 0 || direction == 2 {
		// north or south
		deltaY := 0
		if direction == 0 {
			deltaY = -1
		} else {
			deltaY = 1
		}

		deltaX := 0
		if graph.Get(position.X, position.Y) == '[' {
			deltaX = +1
		} else {
			if graph.Get(position.X, position.Y) == ']' {
				deltaX = -1
			} else {
				return false
			}
		}

		if !graph.In(position.X, position.Y+deltaY) {
			return false
		}
		if !graph.In(position.X+deltaX, position.Y+deltaY) {
			return false
		}
		if graph.Get(position.X, position.Y+deltaY) == '#' || graph.Get(position.X+deltaX, position.Y+deltaY) == '#' {
			return false
		}
		pushed1 := Push(graph, utility.Point{X: position.X, Y: position.Y + deltaY}, direction, true)
		pushed2 := Push(graph, utility.Point{X: position.X + deltaX, Y: position.Y + deltaY}, direction, true)
		if pushed1 != pushed2 {
			return false
		}
		Push(graph, utility.Point{X: position.X, Y: position.Y + deltaY}, direction, dry)
		Push(graph, utility.Point{X: position.X + deltaX, Y: position.Y + deltaY}, direction, dry)
		if graph.Get(position.X, position.Y+deltaY) == '.' && graph.Get(position.X+deltaX, position.Y+deltaY) == '.' {
			if !dry {
				graph.Set(position.X, position.Y+deltaY, graph.Get(position.X, position.Y))
				graph.Set(position.X+deltaX, position.Y+deltaY, graph.Get(position.X+deltaX, position.Y))
				graph.Set(position.X, position.Y, '.')
				graph.Set(position.X+deltaX, position.Y, '.')
			}
			return true
		}
	}

	if direction == 1 {
		// east
		if graph.Get(position.X, position.Y) == '[' {
			nextX := position.X + 2
			if !graph.In(nextX, position.Y) {
				return false
			}
			if graph.Get(nextX, position.Y) == '#' {
				return false
			}
			Push(graph, utility.Point{X: nextX, Y: position.Y}, direction, dry)
			if graph.Get(nextX, position.Y) == '.' {
				if !dry {
					graph.Set(position.X+2, position.Y, graph.Get(position.X+1, position.Y))
					graph.Set(position.X+1, position.Y, graph.Get(position.X, position.Y))
					graph.Set(position.X, position.Y, '.')
				}
				return true
			}
		}
	}

	if direction == 3 {
		// west
		if graph.Get(position.X, position.Y) == ']' {
			nextX := position.X - 2
			if !graph.In(nextX, position.Y) {
				return false
			}
			if graph.Get(nextX, position.Y) == '#' {
				return false
			}
			Push(graph, utility.Point{X: nextX, Y: position.Y}, direction, dry)
			if graph.Get(nextX, position.Y) == '.' {
				if !dry {
					graph.Set(position.X-2, position.Y, graph.Get(position.X-1, position.Y))
					graph.Set(position.X-1, position.Y, graph.Get(position.X, position.Y))
					graph.Set(position.X, position.Y, '.')
				}
				return true
			}
		}
	}

	return false
}

func Step2(graph *utility.CharGraph, position utility.Point, direction int) utility.Point {
	nextX := position.X + utility.Directions[direction][0]
	nextY := position.Y + utility.Directions[direction][1]
	if !graph.In(nextX, nextY) {
		return position
	}
	if graph.Get(nextX, nextY) == '#' {
		return position
	}
	if graph.Get(nextX, nextY) != '.' {
		// recursively push
		Push(graph, utility.Point{X: nextX, Y: nextY}, direction, false)
	}
	if graph.Get(nextX, nextY) == '.' {
		graph.Set(nextX, nextY, '@')
		graph.Set(position.X, position.Y, '.')
		return utility.Point{X: nextX, Y: nextY}
	}
	return position
}

func Value(graph *utility.CharGraph) {
	result := 0
	for y := 0; y < graph.Height; y++ {
		for x := 0; x < graph.Width; x++ {
			if graph.Get(x, y) == 'O' {
				result += 100*y + x
			}
		}
	}
	log.Println(result)
}

func Value2(graph *utility.CharGraph) {
	result := 0
	for y := 0; y < graph.Height; y++ {
		for x := 0; x < graph.Width; x++ {
			if graph.Get(x, y) == '[' {
				result += 100*y + x
			}
		}
	}
	log.Println(result)
}

func Part1(fileName string) {
	scanner := utility.ScanFile(fileName)
	graph := utility.NewGraphFromScanner(scanner)
	dirs := map[byte]int{
		'>': 1,
		'<': 3,
		'^': 0,
		'v': 2,
	}
	position := graph.Find('@')
	rows := utility.FileLinesFromScanner(scanner)
	for _, row := range rows {
		for _, symbol := range row {
			position = Step(graph, position, dirs[byte(symbol)])
		}
	}
	Value(graph)
}

func Part2(fileName string) {
	scanner := utility.ScanFile(fileName)
	graph := utility.NewGraphFromScanner(scanner)
	for y := 0; y < graph.Height; y++ {
		row := ""
		for _, symbol := range graph.Rows[y] {
			if symbol == '#' {
				row += "##"
			}
			if symbol == 'O' {
				row += "[]"
			}
			if symbol == '.' {
				row += ".."
			}
			if symbol == '@' {
				row += "@."
			}
		}
		graph.Rows[y] = row
	}
	graph.Width = graph.Width * 2
	dirs := map[byte]int{
		'>': 1,
		'<': 3,
		'^': 0,
		'v': 2,
	}
	position := graph.Find('@')
	rows := utility.FileLinesFromScanner(scanner)
	graph.Dump()
	for index, row := range rows {
		for j, symbol := range row {
			log.Println(index, "/", len(rows), ";", j, "/", len(row))
			fmt.Println(string(symbol), dirs[byte(symbol)])
			position = Step2(graph, position, dirs[byte(symbol)])
			graph.Dump()
		}
	}
	Value2(graph)
}

func main() {
	// Part1("cmd/es15/test.txt")
	// Part1("cmd/es15/input.txt")
	Part2("cmd/es15/test2.txt")
	// Part2("cmd/es15/input.txt")
}
