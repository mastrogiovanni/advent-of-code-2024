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

type VisitContext struct {
	Target  utility.Point
	Visited map[int]bool
	Path    []*Node
	Cost    int
	Best    *Solution
}

type Solution struct {
	Path []*Node
	Cost int
}

func (tg *TraversableGraph) NodeId(px, py int) int {
	return py*tg.Graph.Width + px
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
	nodeId := tg.NodeId(px, py)
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
		neighborId := tg.NodeId(nx, ny)
		node, ok := tg.Nodes[neighborId]
		if !ok {
			node = CreateNodes(tg, nx, ny)
		}
		neighbors = append(neighbors, node)
	}
	resultNode.Neighbors = neighbors
	return resultNode
}

func IsTurn(path []*Node, index int) bool {
	if index+1 >= len(path) {
		return false
	}
	if index-1 < 0 {
		return false
	}
	dx1 := path[index].Position.X - path[index-1].Position.X
	dx2 := path[index+1].Position.X - path[index].Position.X
	if dx1 != dx2 {
		return true
	}
	dy1 := path[index].Position.Y - path[index-1].Position.Y
	dy2 := path[index+1].Position.Y - path[index].Position.Y
	return dy1 != dy2
}

// var bestSolution *Solution = nil

// func Dump(context *VisitContext) {
// 	context.Cost -= 1
// 	dx := context.Path[1].Position.X - context.Path[0].Position.X
// 	if dx == 0 {
// 		// vertical
// 		context.Cost += 1000
// 	} else {
// 		if dx < 0 {
// 			context.Cost += 2000
// 		}
// 	}
// 	if bestSolution == nil {
// 		bestSolution = &Solution{
// 			Path: context.Path,
// 			Cost: context.Cost,
// 		}
// 	} else {
// 		if bestSolution.Cost > context.Cost {
// 			bestSolution = &Solution{
// 				Path: context.Path,
// 				Cost: context.Cost,
// 			}
// 		}
// 	}
// }

func Next(neighbors []*Node, past *Node) *Node {
	if neighbors[0] == past {
		return neighbors[1]
	} else {
		return neighbors[0]
	}
}

// Return the path from position to end and its cost
func Visit(tg *TraversableGraph, position utility.Point, context *VisitContext) *Solution {

	// nodeId := tg.NodeId(position.X, position.Y)
	// node := tg.Nodes[nodeId]
	// queue := list.New()
	// queue.PushBack(node)

	// parents := make(map[int]*Node)
	// parents[node.Id] = nil

	// for queue.Len() > 0 {
	// 	currentNode := queue.Front().Value.(*Node)
	// 	queue.Remove(queue.Front())

	// 	// compare if node is equals to target
	// 	if currentNode.Position.X == context.Target.X && currentNode.Position.Y == context.Target.Y {
	// 		// the target has been looked
	// 		// reconstructing the path
	// 		var route []*Node
	// 		for len(currentNode.Value) > 0 {
	// 			// recreating route
	// 			route = append([]string{currentNode}, route...)
	// 			// changing pointer
	// 			currentNode.Value = parents[currentNode.Value]
	// 		}

	// 		// returning path result
	// 		return strings.Join(route, "->")
	// 	}

	// 	for _, neighbor := range node.Neighbors {
	// 		// check if the neighbor has not already been visited
	// 		if _, visited := context.Visited[neighbor.Id]; !visited {
	// 			// add neighbor to parents map associated to current node value
	// 			parents[neighbor.Id] = currentNode
	// 			// add neighbor to the end of the queue
	// 			queue.PushBack(neighbor)
	// 		}
	// 	}
	// }

	// log.Println(len(context.Path))
	nodeId := tg.NodeId(position.X, position.Y)
	node := tg.Nodes[nodeId]
	path := append(context.Path, node)
	context.Visited[nodeId] = true
	toUnvisit := make([]int, 0)
	toUnvisit = append(toUnvisit, nodeId)

	defer func() {
		for _, nodeId := range toUnvisit {
			delete(context.Visited, nodeId)
		}
	}()

	// +1000 if this node was reached after a turn
	turnCost := 1
	if IsTurn(path, len(path)-2) {
		turnCost = 1000 + 1
	}

	// If Visited the target node
	if position.X == context.Target.X && position.Y == context.Target.Y {
		// Dump(&VisitContext{
		// 	Path: path,
		// 	Cost: context.Cost + turnCost,
		// })
		dx := context.Path[1].Position.X - context.Path[0].Position.X
		cost := context.Cost + turnCost - 1
		if dx == 0 {
			// vertical
			cost += 1000
		} else {
			if dx < 0 {
				cost += 2000
			}
		}
		solution := &Solution{
			Path: path,
			Cost: cost,
		}
		if context.Best == nil {
			context.Best = solution
		} else {
			if solution.Cost < context.Best.Cost {
				context.Best = solution
			}
		}
		return context.Best
	}

	for _, neighbor := range node.Neighbors {

		if _, ok := context.Visited[neighbor.Id]; ok {
			continue
		}

		if context.Best != nil && context.Best.Cost < context.Cost+turnCost {
			continue
		}

		solution := Visit(tg, neighbor.Position, &VisitContext{
			Target:  context.Target,
			Visited: context.Visited,
			Path:    path,
			Cost:    context.Cost + turnCost,
		})
		if solution == nil {
			continue
		}

		if context.Best == nil {
			context.Best = solution
		} else {
			if solution.Cost < context.Best.Cost {
				context.Best = solution
			}
		}

	}

	return context.Best

}

func main() {
	graph := utility.NewGraph("cmd/es16/input.txt")
	tg := NewTraversableGraph(graph)
	start := graph.Find('S')
	end := graph.Find('E')
	context := &VisitContext{
		Target:  end,
		Visited: make(map[int]bool),
	}
	log.Println("Starting")
	solution := Visit(tg, start, context)
	log.Println("Solution")
	for _, item := range solution.Path {
		log.Println(item)
	}
	log.Println(solution.Cost)

}
