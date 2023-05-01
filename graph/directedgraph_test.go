package graph_test

import (
	"sort"
	"testing"

	"github.com/brmatvey/go-graphs/graph"
)

func TestDirectedGraph(t *testing.T) {
	t.Run("test creator", func(t *testing.T) {
		creator := graph.NewDirectedGraphCreator(map[int][]int{1: {2, 3}, 2: {}, 3: {}})
		directedGraph, err := graph.NewDirectedGraphFromCreator(creator)
		if err != nil {
			t.Fatal("err must be nil")
		}

		n1, ok := directedGraph.Node(1)
		if !ok {
			t.Fatal("key must exist")
		}
		n2, ok := directedGraph.Node(2)
		if !ok {
			t.Fatal("key must exist")
		}
		n3, ok := directedGraph.Node(3)
		if !ok {
			t.Fatal("key must exist")
		}
		if len(n2.Children()) != 0 || len(n3.Children()) != 0 {
			t.Fatal("n2 and n3 must have no children")
		}

		n1Children := n1.Children()
		sort.SliceStable(n1Children, func(i, j int) bool { return n1Children[i].Key() < n1Children[j].Key() })
		for i, expected := range []graph.Node[int]{n2, n3} {
			if expected != n1Children[i] {
				t.Fatal("must be equal")
			}
			nodeFromGraph, ok := directedGraph.Node(n1Children[i].Key())
			if ok != true {
				t.Fatal("must be true")
			}
			if nodeFromGraph != n1Children[i] {
				t.Fatal("must be equal")
			}
		}
	})

	t.Run("simple graph", func(t *testing.T) {
		n1, n2, n3 := graph.NewNode(1), graph.NewNode(2), graph.NewNode(3)
		n1.AddChildren(n2)
		n1.AddChildren(n3)

		directedGraph, err := graph.NewDirectedGraph(n1)
		if err != nil {
			t.Fatal("err must be nil")
		}

		nodes := directedGraph.Nodes()
		sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Key() < nodes[j].Key() })
		for i, expected := range []graph.Node[int]{n1, n2, n3} {
			if expected != nodes[i] {
				t.Fatal("must be equal")
			}
			nodeFromGraph, ok := directedGraph.Node(nodes[i].Key())
			if ok != true {
				t.Fatal("must be true")
			}
			if nodeFromGraph != nodes[i] {
				t.Fatal("must be equal")
			}
		}
	})

	t.Run("simple graph with multiple nodes", func(t *testing.T) {
		n1, n2, n3 := graph.NewNode(1), graph.NewNode(2), graph.NewNode(3)
		n1.AddChildren(n2)
		n1.AddChildren(n3)

		directedGraph, err := graph.NewDirectedGraph(n1, n2, n3)
		if err != nil {
			t.Fatal("err must be nil")
		}

		nodes := directedGraph.Nodes()
		sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Key() < nodes[j].Key() })
		for i, expected := range []graph.Node[int]{n1, n2, n3} {
			if expected != nodes[i] {
				t.Fatal("must be equal")
			}
			nodeFromGraph, ok := directedGraph.Node(nodes[i].Key())
			if ok != true {
				t.Fatal("must be true")
			}
			if nodeFromGraph != nodes[i] {
				t.Fatal("must be equal")
			}
		}
	})

	t.Run("graph with circular dependencies", func(t *testing.T) {
		n1, n2 := graph.NewNode(1), graph.NewNode(2)
		n1.AddChildren(n2)
		n2.AddChildren(n1)

		directedGraph, err := graph.NewDirectedGraph(n1)
		if err != nil {
			t.Fatal("err must be nil")
		}

		nodes := directedGraph.Nodes()
		sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Key() < nodes[j].Key() })
		for i, expected := range []graph.Node[int]{n1, n2} {
			if expected != nodes[i] {
				t.Fatal("must be equal")
			}
			nodeFromGraph, ok := directedGraph.Node(nodes[i].Key())
			if ok != true {
				t.Fatal("must be true")
			}
			if nodeFromGraph != nodes[i] {
				t.Fatal("must be equal")
			}
		}
	})

	t.Run("graph with circular dependencies multiple nodes constructor", func(t *testing.T) {
		n1, n2 := graph.NewNode(1), graph.NewNode(2)
		n1.AddChildren(n2)
		n2.AddChildren(n1)

		directedGraph, err := graph.NewDirectedGraph(n1, n2)
		if err != nil {
			t.Fatal("err must be nil")
		}

		nodes := directedGraph.Nodes()
		sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Key() < nodes[j].Key() })
		for i, expected := range []graph.Node[int]{n1, n2} {
			if expected != nodes[i] {
				t.Fatal("must be equal")
			}
			nodeFromGraph, ok := directedGraph.Node(nodes[i].Key())
			if ok != true {
				t.Fatal("must be true")
			}
			if nodeFromGraph != nodes[i] {
				t.Fatal("must be equal")
			}
		}
	})

	t.Run("graph with repeated key", func(t *testing.T) {
		n1, n2, n3 := graph.NewNode(1), graph.NewNode(2), graph.NewNode(1)
		n1.AddChildren(n2)
		n1.AddChildren(n3)

		_, err := graph.NewDirectedGraph(n1)
		if err == nil {
			t.Fatal("err must be not nil")
		}

	})
}
