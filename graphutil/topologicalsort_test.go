package graphutil_test

import (
	"fmt"
	"testing"

	"github.com/brmatvey/go-graphs/graph"
	"github.com/brmatvey/go-graphs/graphutil"

	"github.com/brmatvey/go-data-structs/slice"
)

func TestTopologicalSort(t *testing.T) {

	t.Run("simple sequence", func(t *testing.T) {

		seq := []string{"start", "running", "finish"}
		graph := generateList(seq...)
		sortedSeq, err := graphutil.TopologicalSort(graph)
		if err != nil {
			t.Fatal(err)
		}

		if len(seq) != len(sortedSeq) {
			t.Fatal("inconsistent len")
		}

		for i, expected := range seq {
			if expected != sortedSeq[i].Key() {
				t.Fatal(fmt.Sprintf("inconsistent key %d %s %s", i, expected, sortedSeq[i].Key()))
			}
		}
	})

	t.Run("circled sequence", func(t *testing.T) {
		_, err := graphutil.TopologicalSort(generateLoopedGraph())
		if err == nil {
			t.Fatal("err must be not nil")
		}
	})

	t.Run("branched struct", func(t *testing.T) {
		graph, checker := branchedGraph()
		sortedSeq, err := graphutil.TopologicalSort(graph)
		if err != nil {
			t.Fatal(err)
		}

		if !checker(sortedSeq) {
			t.Fatal("inconsistent order")
		}
	})
}

func generateList(list ...string) graph.DirectedGraph[string] {
	slice.Reverse(list)
	child := (graph.Node[string])(nil)
	for _, leaf := range list {
		if child == nil {
			child = graph.NewNode(leaf)
			continue
		}
		child = graph.NewNode(leaf, child)
	}
	slice.Reverse(list)
	graph, _ := graph.NewDirectedGraph(child)
	return graph
}

func generateLoopedGraph() graph.DirectedGraph[string] {
	v1, v2 := graph.NewNode("start"), graph.NewNode("finish")
	v1.AddChildren(v2)
	v2.AddChildren(v1)
	graph, _ := graph.NewDirectedGraph(v1, v2)
	return graph
}

func branchedGraph() (graph.DirectedGraph[string], func(actual []graph.Node[string]) bool) {
	// start -> eat -> commute -> work
	//      \->   smoking    -/
	dependencies := map[string][]string{
		"start":   {"eat", "smoking"},
		"eat":     {"commute"},
		"commute": {"work"},
		"smoking": {"work"},
		"work":    {},
	}

	checker := func(actual []graph.Node[string]) bool {
		actualSeq := make([]string, len(actual))
		for i, node := range actual {
			actualSeq[i] = node.Key()
		}

		possibleSeqs := [][]string{
			{"start", "eat", "commute", "smoking", "work"},
			{"start", "eat", "smoking", "commute", "work"},
			{"start", "smoking", "eat", "commute", "work"},
		}

		return slice.IsEqual(actualSeq, possibleSeqs[0]) ||
			slice.IsEqual(actualSeq, possibleSeqs[1]) ||
			slice.IsEqual(actualSeq, possibleSeqs[2])
	}
	graph, _ := graph.NewDirectedGraphFromCreator(graph.NewDirectedGraphCreator(dependencies))
	return graph, checker
}
