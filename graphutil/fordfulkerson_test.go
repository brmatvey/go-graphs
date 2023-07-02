package graphutil_test

import (
	"testing"

	"github.com/brmatvey/go-graphs/graph"
	"github.com/brmatvey/go-graphs/graphutil"
)

func TestFordFulkerson(t *testing.T) {
	t.Run("test linear graph", func(t *testing.T) {
		//   1    2    3
		// 1 -> 2 -> 3 -> 4
		dependencies := map[int][]graph.Length[int]{
			1: {graph.NewLength(2, 1)},
			2: {graph.NewLength(3, 2)},
			3: {graph.NewLength(4, 3)},
			4: {},
		}
		count := 0
		edgeKeyGen := func() int {
			count++
			return count
		}

		weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, edgeKeyGen))
		if err != nil {
			t.Fatal("error must be nil")
		}

		flow := graphutil.FordFulkerson(1, 4, weightedGraph)
		if flow != 1 {
			t.Fatal("incorrect flow")
		}
	})

	t.Run("test random graph", func(t *testing.T) {
		dependencies := map[int][]graph.Length[int]{
			1: {graph.NewLength(2, 15), graph.NewLength(3, 1)},
			2: {graph.NewLength(4, 16)},
			3: {graph.NewLength(5, 3), graph.NewLength(6, 2)},
			4: {graph.NewLength(7, 8), graph.NewLength(6, 10)},
			5: {graph.NewLength(7, 5)},
			6: {graph.NewLength(8, 3)},
			7: {graph.NewLength(8, 10)},
			8: {},
		}
		count := 0
		edgeKeyGen := func() int {
			count++
			return count
		}

		weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, edgeKeyGen))
		if err != nil {
			t.Fatal("error must be nil")
		}

		flow := graphutil.FordFulkerson(1, 8, weightedGraph)
		if flow != 12 {
			t.Fatal("incorrect flow")
		}

	})

	t.Run("test random graph 2", func(t *testing.T) {
		dependencies := map[rune][]graph.Length[rune]{
			'A': {graph.NewLength('B', 7), graph.NewLength('C', 4)},
			'B': {graph.NewLength('C', 4), graph.NewLength('E', 2)},
			'C': {graph.NewLength('D', 4), graph.NewLength('E', 8)},
			'D': {graph.NewLength('F', 12)},
			'E': {graph.NewLength('D', 4), graph.NewLength('F', 5)},
		}
		count := 0
		edgeKeyGen := func() int {
			count++
			return count
		}

		weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, edgeKeyGen))
		if err != nil {
			t.Fatal("error must be nil")
		}

		flow := graphutil.FordFulkerson('A', 'F', weightedGraph)
		if flow != 10 {
			t.Fatal("incorrect flow")
		}
	})

	t.Run("test random graph 3 (see png)", func(t *testing.T) {
		dependencies := map[int][]graph.Length[int]{
			1:  {graph.NewLength(2, 6), graph.NewLength(3, 6), graph.NewLength(4, 8), graph.NewLength(5, 9)},
			2:  {graph.NewLength(3, 3), graph.NewLength(6, 4)},
			3:  {graph.NewLength(4, 4), graph.NewLength(7, 4)},
			4:  {graph.NewLength(5, 3), graph.NewLength(8, 5), graph.NewLength(9, 10)},
			5:  {graph.NewLength(9, 6)},
			6:  {graph.NewLength(3, 9), graph.NewLength(10, 5)},
			7:  {graph.NewLength(4, 10), graph.NewLength(6, 8), graph.NewLength(11, 5)},
			8:  {graph.NewLength(7, 8), graph.NewLength(12, 5), graph.NewLength(13, 12)},
			9:  {graph.NewLength(8, 7), graph.NewLength(13, 7)},
			10: {graph.NewLength(7, 10), graph.NewLength(14, 6)},
			11: {graph.NewLength(8, 12), graph.NewLength(10, 8), graph.NewLength(14, 9)},
			12: {graph.NewLength(11, 8), graph.NewLength(14, 7)},
			13: {graph.NewLength(12, 7), graph.NewLength(14, 6)},
			14: {},
		}
		count := 0
		edgeKeyGen := func() int {
			count++
			return count
		}

		weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, edgeKeyGen))
		if err != nil {
			t.Fatal("error must be nil")
		}

		flow := graphutil.FordFulkerson(1, 14, weightedGraph)
		if flow != 26 {
			t.Fatal("incorrect flow")
		}
	})

	t.Run("test with circles", func(t *testing.T) {
		dependencies := map[rune][]graph.Length[rune]{
			'A': {graph.NewLength('B', 7)},
			'B': {graph.NewLength('A', 7), graph.NewLength('C', 8)},
			'C': {},
		}
		count := 0
		edgeKeyGen := func() int {
			count++
			return count
		}

		weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, edgeKeyGen))
		if err != nil {
			t.Fatal("error must be nil")
		}

		flow := graphutil.FordFulkerson('A', 'C', weightedGraph)
		if flow != 7 {
			t.Fatal("incorrect flow")
		}
	})
}
