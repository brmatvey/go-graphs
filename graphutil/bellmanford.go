package graphutil

import (
	"errors"

	"github.com/brmatvey/go-graphs/graph"
)

func BellmanFord[K, T comparable](start T, weightedGraph graph.WeightedGraph[K, T]) (map[T]float64, error) {
	nodes, edges := weightedGraph.Nodes(), weightedGraph.Edges()
	res := make(map[T]float64)
	for _, n := range nodes {
		res[n.Key()] = max
	}
	res[start] = 0.0

	for i := 0; i < len(nodes)-1; i++ {
		for _, e := range edges {
			if res[e.To().Key()] > res[e.From().Key()]+e.Weight() {
				res[e.To().Key()] = res[e.From().Key()] + e.Weight()
			}
		}
	}

	for _, e := range edges {
		if res[e.To().Key()] != max && res[e.To().Key()] > res[e.From().Key()]+e.Weight() {
			return nil, errors.New("negative circular dependencies in graph")
		}
	}

	return res, nil
}
