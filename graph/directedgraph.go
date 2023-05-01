package graph

import "errors"

type DirectedGraph[T comparable] interface {
	Node(T) (Node[T], bool)
	Nodes() []Node[T]
}

func NewDirectedGraphFromCreator[T comparable](creator DirectedGraphCreator[T]) (DirectedGraph[T], error) {

	nodesMap := make(map[T]Node[T])
	getNode := func(key T) Node[T] {
		if _, ok := nodesMap[key]; !ok {
			nodesMap[key] = NewNode(key)
		}
		return nodesMap[key]
	}

	for nodeKey, childrenKeys := range creator.structure {
		children, currentNode := make([]Node[T], 0), getNode(nodeKey)
		for _, childKey := range childrenKeys {
			childNode := getNode(childKey)
			children = append(children, childNode)
		}
		currentNode.AddChildren(children...)
		nodesMap[currentNode.Key()] = currentNode
	}

	nodes := make([]Node[T], 0)
	for _, n := range nodesMap {
		nodes = append(nodes, n)
	}

	return NewDirectedGraph(nodes...)
}

func NewDirectedGraph[T comparable](ns ...Node[T]) (DirectedGraph[T], error) {

	nodesMap := make(map[T]Node[T])
	for _, n := range ns {
		err := bfs(nodesMap, n, func(node Node[T]) error {
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	nodes := make([]Node[T], 0, len(nodesMap))
	for _, n := range nodesMap {
		nodes = append(nodes, n)
	}

	return &directedGraph[T]{
		nodesMap: nodesMap,
	}, nil
}

func bfs[T comparable](set map[T]Node[T], node Node[T], f func(node Node[T]) error) error {
	if foundNode, ok := set[node.Key()]; ok {
		if foundNode != node {
			return errors.New("repeated key")
		}
		return nil
	}
	set[node.Key()] = node
	if err := f(node); err != nil {
		return err
	}
	for _, n := range node.Children() {
		if err := bfs(set, n, f); err != nil {
			return err
		}
	}
	return nil
}

type directedGraph[T comparable] struct {
	nodesMap map[T]Node[T]
}

func (g *directedGraph[T]) Node(key T) (Node[T], bool) {
	n, ok := g.nodesMap[key]
	return n, ok
}

func (g *directedGraph[T]) Nodes() []Node[T] {
	nodes := make([]Node[T], 0, len(g.nodesMap))
	for _, n := range g.nodesMap {
		nodes = append(nodes, n)
	}
	return nodes
}
