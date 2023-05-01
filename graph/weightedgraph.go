package graph

import (
	"errors"
	"fmt"
)

type WeightedGraph[K, T comparable] interface {
	DirectedGraph[T]

	Edge(K) (Edge[K, T], bool)
	FindEdge(from, to T) (Edge[K, T], bool)
	Edges() []Edge[K, T]
}

func NewWeightedGraphFromCreator[K, T comparable](creator WeightedGraphCreator[K, T]) (WeightedGraph[K, T], error) {
	nodesMap, edgesMap := make(map[T]Node[T]), make(map[K]Edge[K, T])

	getNode := func(key T) Node[T] {
		if _, ok := nodesMap[key]; !ok {
			nodesMap[key] = NewNode(key)
		}
		return nodesMap[key]
	}

	for nodeKey, childrenSettings := range creator.structure {
		children, currentNode := make([]Node[T], 0), getNode(nodeKey)
		for _, settings := range childrenSettings {
			childNode := getNode(settings.to)
			children = append(children, childNode)
			e := NewEdge(creator.uniqueKGen(), settings.weight, currentNode, childNode)
			edgesMap[e.Key()] = e
		}
		currentNode.AddChildren(children...)
		nodesMap[currentNode.Key()] = currentNode
	}

	nodes, edges := make([]Node[T], 0), make([]Edge[K, T], 0)
	for _, n := range nodesMap {
		nodes = append(nodes, n)
	}
	for _, e := range edgesMap {
		edges = append(edges, e)
	}
	return NewWeightedGraph(nodes, edges)
}

func NewWeightedGraph[K, T comparable](ns []Node[T], edges []Edge[K, T]) (WeightedGraph[K, T], error) {
	graph, err := NewDirectedGraph[T](ns...)
	if err != nil {
		return nil, err
	}

	requiredWeights := make(map[path[T]]struct{})
	nodesMap := make(map[T]Node[T])
	for _, n := range graph.Nodes() {
		err = bfs(nodesMap, n, func(node Node[T]) error {
			for _, child := range node.Children() {
				requiredWeights[newPath(node.Key(), child.Key())] = struct{}{}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	actualWeights := make(map[path[T]]float64)
	edgesPaths, edgesMap := make(map[path[T]]Edge[K, T]), make(map[K]Edge[K, T])
	for _, e := range edges {
		if _, ok := edgesMap[e.Key()]; ok {
			return nil, errors.New(fmt.Sprintf("repeated edge key %v", e.Key()))
		}
		p := newPath(e.From().Key(), e.To().Key())
		actualWeights[p], edgesPaths[p], edgesMap[e.Key()] = e.Weight(), e, e
	}

	for requiredWeight := range requiredWeights {
		if _, ok := actualWeights[requiredWeight]; !ok {
			return nil, errors.New(fmt.Sprintf("path from %v to %v is required", requiredWeight.from, requiredWeight.to))
		}
	}

	for actualWeight := range actualWeights {
		if _, ok := requiredWeights[actualWeight]; !ok {
			return nil, errors.New(fmt.Sprintf("weight from %v to %v is required", actualWeight.from, actualWeight.to))
		}
	}

	return &weightedGraph[K, T]{
		DirectedGraph: graph,
		edgesPaths:    edgesPaths,
		edgesMap:      edgesMap,
	}, nil
}

type weightedGraph[K, T comparable] struct {
	DirectedGraph[T]

	edgesPaths map[path[T]]Edge[K, T]
	edgesMap   map[K]Edge[K, T]
}

func (w *weightedGraph[K, T]) Edge(key K) (Edge[K, T], bool) {
	e, ok := w.edgesMap[key]
	return e, ok
}

func (w *weightedGraph[K, T]) FindEdge(from, to T) (Edge[K, T], bool) {
	e, ok := w.edgesPaths[newPath(from, to)]
	return e, ok
}

func (w *weightedGraph[K, T]) Edges() []Edge[K, T] {
	edges := make([]Edge[K, T], 0, len(w.edgesMap))
	for _, e := range w.edgesMap {
		edges = append(edges, e)
	}
	return edges
}
