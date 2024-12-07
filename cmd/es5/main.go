package main

import (
	"bufio"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type PairGraph struct {
	Graph   map[string][]string
	Scanner *bufio.Scanner
}

func Check(graph map[string][]string, path []string) int {
	for i := 1; i < len(path); i++ {
		// log.Printf("%v %v\n", path[i-1], path[i])
		l, ok := graph[path[i]]
		if !ok {
			continue
		}
		if slices.Contains(l, path[i-1]) {
			return 0
		}
	}
	return 1
}

func FirstBroken(graph map[string][]string, path []string) int {
	for i := 1; i < len(path); i++ {
		// log.Printf("%v %v\n", path[i-1], path[i])
		l, ok := graph[path[i]]
		if !ok {
			continue
		}
		if slices.Contains(l, path[i-1]) {
			return i - 1
		}
	}
	return -1
}

func Repair(graph map[string][]string, path []string) {
	i := FirstBroken(graph, path)
	if i < 0 {
		return
	} else {
		a := path[i]
		path[i] = path[i+1]
		path[i+1] = a
		Repair(graph, path)
	}
}

func Part1(graph *PairGraph) {
	sum := 0
	for graph.Scanner.Scan() {
		line := graph.Scanner.Text()
		path := strings.Split(line, ",")
		if Check(graph.Graph, path) == 1 {
			v, e := strconv.Atoi(path[len(path)/2])
			if e != nil {
				log.Fatal(e)
			}
			sum += v
		}
	}
	log.Println(sum)
}

func Part2(graph *PairGraph) {
	sum := 0
	for graph.Scanner.Scan() {
		line := graph.Scanner.Text()
		path := strings.Split(line, ",")
		if Check(graph.Graph, path) == 0 {
			Repair(graph.Graph, path)
			v, e := strconv.Atoi(path[len(path)/2])
			if e != nil {
				log.Fatal(e)
			}
			sum += v
		}
	}
	log.Println(sum)
}

func NewPairGraph(scanner *bufio.Scanner) *PairGraph {
	graph := make(map[string][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pair := strings.Split(line, "|")
		if l, ok := graph[pair[0]]; ok {
			graph[pair[0]] = append(l, pair[1])
		} else {
			graph[pair[0]] = []string{pair[1]}
		}
	}
	return &PairGraph{
		Graph:   graph,
		Scanner: scanner,
	}
}

func main() {
	Part1(NewPairGraph(utility.ScanFile("cmd/es5/test.txt")))
	Part1(NewPairGraph(utility.ScanFile("cmd/es5/input.txt")))
	Part2(NewPairGraph(utility.ScanFile("cmd/es5/test.txt")))
	Part2(NewPairGraph(utility.ScanFile("cmd/es5/input.txt")))
}
