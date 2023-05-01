package graphutil_test

import (
	"testing"

	"github.com/brmatvey/go-graphs/graph"
	"github.com/brmatvey/go-graphs/graphutil"
)

func TestDijkstra(t *testing.T) {
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

		lengths, err := graphutil.Dijkstra(1, weightedGraph)
		if err != nil {
			t.Fatal("error must be nil")
		}

		if lengths[1] != 0.0 || lengths[2] != 1.0 || lengths[3] != 3.0 || lengths[4] != 6.0 {
			t.Fatal("incorrect length")
		}
	})

	t.Run("test random graph", func(t *testing.T) {
		//   6   7    8       10
		// 1 ->2 -> 4 -> -> 7 -> -> -> 8
		// 1\      3    5 /    /
		//   \->3 -> 5 ->/    /
		//     2 \           / 3
		//        \-> 6 -> ->
		dependencies := map[int][]graph.Length[int]{
			1: {graph.NewLength(2, 6), graph.NewLength(3, 1)},
			2: {graph.NewLength(4, 7)},
			3: {graph.NewLength(5, 3), graph.NewLength(6, 2)},
			4: {graph.NewLength(7, 8)},
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

		lengths, err := graphutil.Dijkstra(1, weightedGraph)
		if err != nil {
			t.Fatal("error must be nil")
		}

		if lengths[1] != 0 || lengths[2] != 6 || lengths[3] != 1 || lengths[4] != 13 ||
			lengths[5] != 4 || lengths[6] != 3 || lengths[7] != 9 || lengths[8] != 6 {
			t.Fatal("incorrect length")
		}

	})
}
