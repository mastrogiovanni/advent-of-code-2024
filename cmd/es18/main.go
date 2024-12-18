package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func Dump(end utility.Point, graph map[utility.Point]bool) {
	fmt.Println("\n-------------------")
	for y := 0; y < end.Y+1; y++ {
		fmt.Println()
		for x := 0; x < end.X+1; x++ {
			_, ok := graph[utility.Point{X: x, Y: y}]
			if ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
	}
}

func Visit(end utility.Point, graph map[utility.Point]bool) int {
	lenghts := make(map[utility.Point]int)
	visited := make(map[utility.Point]bool)
	queue := make([]utility.Point, 1)
	queue[0] = end
	lenghts[end] = 0

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
			if !(nx >= 0 && ny >= 0 && nx <= end.X && ny <= end.Y) {
				// Outside graph
				continue
			}
			np := utility.Point{X: nx, Y: ny}
			if _, ok := graph[np]; ok {
				// Cell is blocked
				continue
			}
			lold, ok := lenghts[np]
			if !ok {
				lenghts[np] = l + 1
			} else {
				lenghts[np] = min(lold, l+1)
			}
			queue = append(queue, np)
		}
	}

	l, ok := lenghts[utility.Point{X: 0, Y: 0}]
	if ok {
		return l
	} else {
		return -1
	}
}

func Part1(fileName string, end utility.Point, steps int) {
	scanner := utility.ScanFile(fileName)
	graph := make(map[utility.Point]bool)
	for scanner.Scan() {
		if steps == 0 {
			break
		}
		steps--
		row := scanner.Text()
		components := strings.Split(row, ",")
		x, _ := strconv.Atoi(components[0])
		y, _ := strconv.Atoi(components[1])
		graph[utility.Point{X: x, Y: y}] = true
	}
	log.Println(Visit(end, graph))
}

func Part2(fileName string, end utility.Point) {
	scanner := utility.ScanFile(fileName)
	graph := make(map[utility.Point]bool)
	lastPoint := utility.Point{}
	for scanner.Scan() {
		row := scanner.Text()
		components := strings.Split(row, ",")
		x, _ := strconv.Atoi(components[0])
		y, _ := strconv.Atoi(components[1])
		graph[utility.Point{X: x, Y: y}] = true
		l := Visit(end, graph)
		if l < 0 {
			lastPoint = utility.Point{X: x, Y: y}
			break
		}
	}
	log.Printf("%d,%d", lastPoint.X, lastPoint.Y)
}

func main() {
	// Part1("cmd/es18/test.txt", utility.Point{X: 6, Y: 6}, 12)
	// Part1("cmd/es18/input.txt", utility.Point{X: 70, Y: 70}, 1024)
	Part2("cmd/es18/test.txt", utility.Point{X: 6, Y: 6})
	Part2("cmd/es18/input.txt", utility.Point{X: 70, Y: 70})
}
