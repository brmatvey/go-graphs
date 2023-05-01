package graph_test

import (
	"sort"
	"testing"

	"github.com/brmatvey/go-graphs/graph"
)

func TestWeightedGraph(t *testing.T) {
	t.Run("test creator", func(t *testing.T) {
		dependencies := map[int][]graph.Length[int]{
			1: {graph.NewLength(2, 2.0), graph.NewLength(3, 3.0)},
			2: {},
			3: {},
		}
		count := 0
		uniqueKGen := func() int {
			count++
			return count
		}
		weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, uniqueKGen))
		if err != nil {
			t.Fatal("err must be nil")
		}

		n1, ok := weightedGraph.Node(1)
		if !ok {
			t.Fatal("key must exist")
		}
		n2, ok := weightedGraph.Node(2)
		if !ok {
			t.Fatal("key must exist")
		}
		n3, ok := weightedGraph.Node(3)
		if !ok {
			t.Fatal("key must exist")
		}
		if len(n2.Children()) != 0 || len(n3.Children()) != 0 {
			t.Fatal("n2 and n3 must have no children")
		}

		e1, ok := weightedGraph.FindEdge(n1.Key(), n2.Key())
		if !ok {
			t.Fatal("edge must exist")
		}
		e2, ok := weightedGraph.FindEdge(n1.Key(), n3.Key())
		if !ok {
			t.Fatal("edge must exist")
		}

		validator(t, weightedGraph, []graph.Node[int]{n1, n2, n3}, []graph.Edge[int, int]{e1, e2})
	})

	t.Run("simple graph", func(t *testing.T) {
		n1, n2, n3 := graph.NewNode(1), graph.NewNode(2), graph.NewNode(3)
		n1.AddChildren(n2)
		n1.AddChildren(n3)

		e1, e2 := graph.NewEdge(1, 2.0, n1, n2), graph.NewEdge(2, 3.0, n1, n3)

		weightedGraph, err := graph.NewWeightedGraph([]graph.Node[int]{n1, n2, n3}, []graph.Edge[int, int]{e1, e2})
		if err != nil {
			t.Fatal("err must be nil")
		}

		validator(t, weightedGraph, []graph.Node[int]{n1, n2, n3}, []graph.Edge[int, int]{e1, e2})
	})

	t.Run("graph with circular dependencies", func(t *testing.T) {
		n1, n2 := graph.NewNode(1), graph.NewNode(2)
		n1.AddChildren(n2)
		n2.AddChildren(n1)

		e1, e2 := graph.NewEdge(1, 2.0, n1, n2), graph.NewEdge(2, 3.0, n2, n1)

		weightedGraph, err := graph.NewWeightedGraph([]graph.Node[int]{n1, n2}, []graph.Edge[int, int]{e1, e2})
		if err != nil {
			t.Fatal("err must be nil")
		}

		validator(t, weightedGraph, []graph.Node[int]{n1, n2}, []graph.Edge[int, int]{e1, e2})
	})

	t.Run("graph with repeated node's key", func(t *testing.T) {
		n1, n2 := graph.NewNode(1), graph.NewNode(1)
		n1.AddChildren(n2)
		n2.AddChildren(n1)

		e1, e2 := graph.NewEdge(1, 2.0, n1, n2), graph.NewEdge(2, 3.0, n2, n1)

		_, err := graph.NewWeightedGraph([]graph.Node[int]{n1, n2}, []graph.Edge[int, int]{e1, e2})
		if err == nil {
			t.Fatal("err must be not nil")
		}
	})

	t.Run("graph with repeated edge's key", func(t *testing.T) {
		n1, n2 := graph.NewNode(1), graph.NewNode(2)
		n1.AddChildren(n2)
		n2.AddChildren(n1)

		e1, e2 := graph.NewEdge(1, 2.0, n1, n2), graph.NewEdge(1, 3.0, n2, n1)

		_, err := graph.NewWeightedGraph([]graph.Node[int]{n1, n2}, []graph.Edge[int, int]{e1, e2})
		if err == nil {
			t.Fatal("err must be not nil")
		}
	})

	t.Run("graph with extra nodes amount", func(t *testing.T) {
		n1, n2, n3 := graph.NewNode(1), graph.NewNode(2), graph.NewNode(3)
		n1.AddChildren(n2)
		n1.AddChildren(n3)
		n2.AddChildren(n3)

		e1, e2 := graph.NewEdge(1, 2.0, n1, n2), graph.NewEdge(2, 3.0, n1, n3)

		_, err := graph.NewWeightedGraph([]graph.Node[int]{n1, n2, n3}, []graph.Edge[int, int]{e1, e2})
		if err == nil {
			t.Fatal("err must be not nil")
		}
	})

	t.Run("graph with extra edges amount", func(t *testing.T) {
		n1, n2, n3 := graph.NewNode(1), graph.NewNode(2), graph.NewNode(3)
		n1.AddChildren(n2)
		n1.AddChildren(n3)

		e1, e2 := graph.NewEdge(1, 2.0, n1, n2), graph.NewEdge(2, 3.0, n1, n3)
		e3 := graph.NewEdge(3, 4.0, n2, n3)

		_, err := graph.NewWeightedGraph([]graph.Node[int]{n1, n2, n3}, []graph.Edge[int, int]{e1, e2, e3})
		if err == nil {
			t.Fatal("err must be not nil")
		}
	})
}

func validator(t *testing.T, graph graph.WeightedGraph[int, int], expectedNodes []graph.Node[int], expectedEdges []graph.Edge[int, int]) {
	nodes := graph.Nodes()
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Key() < nodes[j].Key() })
	sort.SliceStable(expectedNodes, func(i, j int) bool { return expectedNodes[i].Key() < expectedNodes[j].Key() })
	for i, expected := range expectedNodes {
		if expected != nodes[i] {
			t.Fatal("must be equal")
		}
		nodeFromGraph, ok := graph.Node(nodes[i].Key())
		if ok != true {
			t.Fatal("must be true")
		}
		if nodeFromGraph != nodes[i] {
			t.Fatal("must be equal")
		}
	}

	edges := graph.Edges()
	sort.SliceStable(edges, func(i, j int) bool { return edges[i].Key() < edges[j].Key() })
	sort.SliceStable(expectedEdges, func(i, j int) bool { return expectedEdges[i].Key() < expectedEdges[j].Key() })
	for i, expected := range expectedEdges {
		if expected != edges[i] {
			t.Fatal("must be equal")
		}
		edgeFromGraph, ok := graph.Edge(edges[i].Key())
		if ok != true {
			t.Fatal("must be true")
		}
		if edgeFromGraph != edges[i] {
			t.Fatal("must be equal")
		}
	}
}
