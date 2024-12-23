package main

import (
	"log"
	"slices"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
	"golang.org/x/exp/maps"
)

type Graph struct {
	Edges map[string][]string
	Nodes map[string]interface{}
}

func (graph *Graph) Connect(a, b string) {
	v, ok := graph.Edges[a]
	if ok {
		graph.Edges[a] = append(v, b)
	} else {
		graph.Edges[a] = []string{b}
	}
	vb, okb := graph.Edges[b]
	if okb {
		graph.Edges[b] = append(vb, a)
	} else {
		graph.Edges[b] = []string{a}
	}
	graph.Nodes[a] = struct{}{}
	graph.Nodes[b] = struct{}{}
}

func (graph *Graph) Connected(a, b string) bool {
	return slices.Contains(graph.Edges[a], b)
}

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]string),
		Nodes: make(map[string]interface{}),
	}
}

func IncreaseClique(graph *Graph, cliqueString string) []string {
	clique := strings.Split(cliqueString, ",")
	result := make([]string, 0)
	for _, node := range clique {
		for _, neighbor := range graph.Edges[node] {
			if slices.Contains(clique, neighbor) {
				continue
			}
			found := true
			for _, x := range clique {
				if !graph.Connected(x, neighbor) {
					found = false
					break
				}
			}
			if found {
				cliqueFound := make([]string, len(clique)+1)
				copy(cliqueFound[0:len(clique)], clique)
				cliqueFound[len(clique)] = neighbor
				slices.Sort(cliqueFound)
				result = append(result, strings.Join(cliqueFound, ","))
			}
		}
	}
	return result
}

func Cliques(graph *Graph, includeOnlyItemsWithT bool) []string {
	cliques := make(map[string]interface{})
	for node := range graph.Nodes {
		neighbors := graph.Edges[node]
		for i := 0; i < len(neighbors)-1; i++ {
			for j := i + 1; j < len(neighbors); j++ {
				a := neighbors[i]
				b := neighbors[j]
				if includeOnlyItemsWithT {
					if !(a[0] == 't' || b[0] == 't' || node[0] == 't') {
						continue
					}
				}
				if graph.Connected(a, b) {
					clique := []string{node, a, b}
					slices.Sort(clique)
					cliques[strings.Join(clique, ",")] = struct{}{}
				}
			}
		}
	}
	return maps.Keys(cliques)
}

func Part1(fileName string) {
	scanner := utility.ScanFile(fileName)
	graph := NewGraph()
	for scanner.Scan() {
		line := scanner.Text()
		computers := strings.Split(line, "-")
		a := computers[0]
		b := computers[1]
		graph.Connect(a, b)
	}
	log.Println(len(Cliques(graph, true)))
}

func Part2(fileName string) {
	scanner := utility.ScanFile(fileName)
	graph := NewGraph()
	for scanner.Scan() {
		line := scanner.Text()
		computers := strings.Split(line, "-")
		a := computers[0]
		b := computers[1]
		graph.Connect(a, b)
	}
	cliques := Cliques(graph, false)
	for {
		others := make(map[string]interface{})
		for _, clique := range cliques {
			biggerCliques := IncreaseClique(graph, clique)
			for _, bc := range biggerCliques {
				others[bc] = struct{}{}
			}
		}
		if len(others) > 0 {
			cliques = maps.Keys(others)
		} else {
			break
		}
	}
	log.Println(cliques[0])
}

func main() {
	Part1("cmd/es23/test.txt")
	Part1("cmd/es23/input.txt")
	Part2("cmd/es23/test.txt")
	Part2("cmd/es23/input.txt")
}
