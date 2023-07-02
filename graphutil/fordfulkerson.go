package graphutil

import (
	"github.com/brmatvey/go-graphs/graph"
)

func FordFulkerson[K, T comparable](start, stop T, graph graph.WeightedGraph[K, T]) float64 {
	res := 0.0
	flows, paths := toFlowsAndPaths(graph)
	for {
		currentPath, err := findPathViaDfs[K, T](start, stop, paths)
		if err != nil {
			return res
		}

		minWeight := flows[newPath(currentPath[0], currentPath[1])]
		for i, j := 0, 1; j < len(currentPath); i, j = i+1, j+1 {
			pathCandidate := flows[newPath(currentPath[i], currentPath[j])]
			if pathCandidate < minWeight {
				minWeight = pathCandidate
			}
		}
		res += minWeight

		// modifying graph
		for i, j := 0, 1; j < len(currentPath); i, j = i+1, j+1 {
			from, to := currentPath[i], currentPath[j]
			// create reverse edges
			flows[newPath(to, from)] += minWeight
			if paths[to] == nil {
				paths[to] = make(map[T]struct{})
			}
			paths[to][from] = struct{}{}
			// loosen the old weight
			flows[newPath(from, to)] -= minWeight
			if flows[newPath(from, to)] == 0 {
				delete(flows, newPath(from, to))
				delete(paths[from], to)
			}
		}
	}
}
