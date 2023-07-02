package graphutil

import (
	"errors"

	"github.com/brmatvey/go-data-structs/slice"
	"github.com/brmatvey/go-data-structs/stack"
	"github.com/brmatvey/go-graphs/graph"
)

func newPath[T comparable](from, to T) path[T] {
	return path[T]{From: from, To: to}
}

type path[T comparable] struct {
	From T
	To   T
}

func findPathViaDfs[K, T comparable](from, to T, paths map[T]map[T]struct{}) ([]T, error) {
	s, res, colors := stack.New[T](), stack.New[T](), make(map[T]int)
	s.Push(from)
	for !s.Empty() {
		currentNode := s.Peek()
		switch colors[currentNode] {
		case 0:
			colors[currentNode] = 1
			res.Push(currentNode)
			if currentNode == to {
				return stackToSlice(res), nil
			}
			for childNode := range paths[currentNode] {
				if colors[childNode] == 1 {
					continue
				}
				s.Push(childNode)
			}
		case 1:
			colors[currentNode] = 2
			res.Pop()
			s.Pop()
		case 2:
			s.Pop()
		}
	}
	return nil, errors.New("no path")
}

func stackToSlice[T any](s *stack.Stack[T]) []T {
	res := make([]T, 0)
	for !s.Empty() {
		res = append(res, s.Peek())
		s.Pop()
	}
	slice.Reverse(res)
	return res
}

func toFlowsAndPaths[K, T comparable](g graph.WeightedGraph[K, T]) (map[path[T]]float64, map[T]map[T]struct{}) {
	flows, paths := make(map[path[T]]float64), make(map[T]map[T]struct{})
	for _, e := range g.Edges() {
		flows[newPath[T](e.From().Key(), e.To().Key())] = e.Weight()
		if paths[e.From().Key()] == nil {
			paths[e.From().Key()] = make(map[T]struct{})
		}
		paths[e.From().Key()][e.To().Key()] = struct{}{}
	}
	return flows, paths
}
