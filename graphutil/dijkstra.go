package graphutil

import (
	"errors"
	"fmt"

	"github.com/brmatvey/go-graphs/graph"
)

const max float64 = 1.7976931348623157e+308

func Dijkstra[K, T comparable](start T, weightedGraph graph.WeightedGraph[K, T]) (map[T]float64, error) {
	res, visited, nodes := make(map[T]float64), make(map[T]bool), weightedGraph.Nodes()

	for _, n := range nodes {
		res[n.Key()] = max
	}
	res[start] = 0.0

	for i := 0; i < len(nodes)+1; i++ {
		currentKey, minLen := (*T)(nil), max
		for _, n := range nodes {
			if !visited[n.Key()] && res[n.Key()] < minLen {
				key := n.Key()
				currentKey, minLen = &key, res[n.Key()]
			}
		}
		if currentKey == nil {
			break
		}
		visited[*currentKey] = true
		currentNode, ok := weightedGraph.Node(*currentKey)
		if !ok {
			return nil, errors.New(fmt.Sprintf("node %v is not found", *currentKey))
		}
		for _, n := range currentNode.Children() {
			if visited[n.Key()] {
				continue
			}
			e, ok := weightedGraph.FindEdge(currentNode.Key(), n.Key())
			if !ok {
				return nil, errors.New(fmt.Sprintf("edge from %v to %v is not found", currentNode.Key(), n.Key()))
			}
			if e.Weight() < 0 {
				return nil, errors.New("negative weight in edge in graph")
			}
			if res[n.Key()] > res[currentNode.Key()]+e.Weight() {
				res[n.Key()] = res[currentNode.Key()] + e.Weight()
			}
		}
	}

	return res, nil
}
