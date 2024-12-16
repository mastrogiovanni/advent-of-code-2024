package main

import (
	"log"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

type Node struct {
	Id        int
	Position  utility.Point
	Neighbors []*Node
}

type TraversableGraph struct {
	Graph *utility.CharGraph
	Nodes map[int]*Node
}

type NodePath struct {
	Node *Node
	Next *NodePath
}

func NodeId(px, py int, w, h int) int {
	return py*w + px
}

func NewTraversableGraph(graph *utility.CharGraph) *TraversableGraph {
	tg := &TraversableGraph{
		Graph: graph,
		Nodes: make(map[int]*Node),
	}
	for x := 0; x < graph.Width; x++ {
		for y := 0; y < graph.Height; y++ {
			CreateNodes(tg, x, y)
		}
	}
	return tg
}

func CreateNodes(tg *TraversableGraph, px, py int) *Node {
	nodeId := NodeId(px, py, tg.Graph.Width, tg.Graph.Height)
	test, ok := tg.Nodes[nodeId]
	if ok {
		return test
	}
	resultNode := &Node{
		Id:       nodeId,
		Position: utility.Point{X: px, Y: py},
	}
	tg.Nodes[nodeId] = resultNode
	neighbors := make([]*Node, 0)
	for i := 0; i < 4; i++ {
		nx := px + utility.Directions[i][0]
		ny := py + utility.Directions[i][1]
		if !tg.Graph.In(nx, ny) {
			continue
		}
		if tg.Graph.Get(nx, ny) == '#' {
			continue
		}
		neighborId := NodeId(nx, ny, tg.Graph.Width, tg.Graph.Height)
		node, ok := tg.Nodes[neighborId]
		if !ok {
			node = CreateNodes(tg, nx, ny)
		}
		neighbors = append(neighbors, node)
	}
	resultNode.Neighbors = neighbors
	return resultNode
}

func Visit(tg *TraversableGraph, position, end utility.Point, visited map[int]bool, path []*Node) []*Node {
	node := tg.Nodes[NodeId(position.X, position.Y, tg.Graph.Width, tg.Graph.Height)]
	visited[node.Id] = true
	if position.X == end.X && position.Y == end.Y {
		// End
		return append(path, node)
	}
	var minDist []*Node = nil
	for _, neighbor := range node.Neighbors {
		if _, ok := visited[neighbor.Id]; ok {
			continue
		}
		d := Visit(tg, neighbor.Position, end, visited, append(path, node))
		if d == nil {
			continue
		}
		if minDist == nil {
			minDist = d
		} else {
			if len(minDist) < len(d) {
				minDist = d
			}
		}
	}
	visited[node.Id] = false
	return minDist
}

func main() {
	graph := utility.NewGraph("cmd/es16/test.txt")
	tg := NewTraversableGraph(graph)
	log.Println(tg)
	start := graph.Find('S')
	end := graph.Find('E')
	visited := make(map[int]bool)
	distance := Visit(tg, start, end, visited, make([]*Node, 0))
	for _, node := range distance {
		log.Println(node)
	}
}
