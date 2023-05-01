package graphutil

import (
	"errors"

	"github.com/brmatvey/go-graphs/graph"

	"github.com/brmatvey/go-data-structs/slice"
	"github.com/brmatvey/go-data-structs/stack"
)

func TopologicalSort[T comparable](directedGraph graph.DirectedGraph[T]) ([]graph.Node[T], error) {
	nodes, colors := directedGraph.Nodes(), make(map[T]int)
	res := make([]graph.Node[T], 0, len(nodes))
	for _, nod := range nodes {
		s := stack.New[graph.Node[T]]()
		s.Push(nod)
		for !s.Empty() {
			currentNode := s.Peek()
			switch colors[currentNode.Key()] {
			case 0:
				colors[currentNode.Key()] = 1
				for _, child := range currentNode.Children() {
					if colors[child.Key()] == 1 {
						return nil, errors.New("cycle")
					}
					s.Push(child)
				}
			case 1:
				colors[currentNode.Key()] = 2
				res = append(res, currentNode)
				s.Pop()
			case 2:
				s.Pop()
			}
		}
	}
	slice.Reverse(res)
	return res, nil
}
