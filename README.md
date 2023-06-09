# Simple implementation of trivial graph algorithms
You can use it, for instance, for educational, interview training etc.
Unfortunately, the simulator does not yet support the general case of a graph. Two vertices in the same direction are connected by at most one edge so far.
## Graph package
### Node
Node is simple generic implementation of graph node. It has knowledge about key (must be unique for graph) and children.
```go
type Node[T comparable] interface {
	Key() T
	Children() []Node[T]
	AddChildren(node ...Node[T])
}
```
For creating node manually call its constructor
```go
childNode := graph.NewNode(1)
parentNode := graph.NewNode(2, childNode)
```
Anyway, you don't need to create node manually if you use creator for creating graph (see below).
### Edge
Edge is simple generic implementation of graph edge. Has knowledge about its generic key, float64 weight and links to start and finish nodes.
```go
type Edge[K, T comparable] interface {
	Key() K
	Weight() float64
	From() Node[T]
	To() Node[T]
}
```
For creating edge manually call its constructor
```go
childNode := graph.NewNode(1)
parentNode := graph.NewNode(2, childNode)

simpleEdge := graph.NewEdge(1, 10.0, parentNode, childNode)
```
Anyway, you don't need to create edge manually if you use creator for creating weighted graph (see below).
### Directed graph
You have opportunities for creating simple directed graph via manual creating each node and its children
```go
childNode := graph.NewNode(1)
parentNode := graph.NewNode(2, childNode)

directedGraph, err := graph.NewDirectedGraph(parentNode)
if err != nil {
    t.Fatal("err must be nil")
}
```
However, you could use more simplistic form for setting dependencies in graph. In this case use map with keys and paths.
```go
dependencies := map[int][]int{
	1: {2, 3}, 
	2: {}, 
	3: {},
}
creator := graph.NewDirectedGraphCreator(dependencies)
directedGraph, err := graph.NewDirectedGraphFromCreator(creator)
if err != nil {
    t.Fatal("err must be nil")
}
```
### Weighted graph
You have opportunities for creating simple weighted graph via manual creating each node and its children, and also all required edges.
```go
childNode := graph.NewNode(1)
parentNode := graph.NewNode(2, childNode)

simpleEdge := graph.NewEdge(1, 10.0, parentNode, childNode)

nodes, edges := []graph.Node[int]{parentNode, childNode}, []graph.Edge[int, int]{simpleEdge}
weightedGraph, err := graph.NewWeightedGraph(nodes, edges)
if err != nil {
    t.Fatal("err must be nil")
}
```
However, you could use more simplistic form for setting dependencies in graph. In this case use map with keys and paths with weights.
You also need use generator for your unique generic edge key.
```go
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
creator := graph.NewWeightedGraphCreator(dependencies, uniqueKGen)
weightedGraph, err := graph.NewWeightedGraphFromCreator(creator)
if err != nil {
    t.Fatal("err must be nil")
}
```
## Graph util package
### Topological sort
The topological sort algorithm takes a directed graph and returns an array of the nodes where each node appears before all the nodes it points to. The ordering of the nodes in the array is called a topological ordering.
For using topological sort first of all create directed graph:
```go
dependencies := map[string][]string{
    "start":   {"eat", "smoking"},
    "eat":     {"commute"},
    "commute": {"work"},
    "smoking": {"work"},
    "work":    {},
}
creator := graph.NewDirectedGraphCreator(dependencies)
directedGraph, err := graph.NewDirectedGraphFromCreator(creator)
if err != nil {
    t.Fatal("err must be nil")
}
```
And call topological sort:
```go
sortedSeq, err := graphutil.TopologicalSort(graph)
if err != nil {
    t.Fatal(err)
}
```
Possible results in this case are:
```go
possibleSeqs := [][]string{
    {"start", "eat", "commute", "smoking", "work"},
    {"start", "eat", "smoking", "commute", "work"},
    {"start", "smoking", "eat", "commute", "work"},
}
```
If you have circular dependencies in graph, topological sort returns error.
### Bellman-Ford
The Bellman–Ford algorithm is an algorithm that computes shortest paths from a single source vertex to all of the other vertices in a weighted digraph. It is slower than Dijkstra's algorithm for the same problem, but more versatile, as it is capable of handling graphs in which some of the edge weights are negative numbers.
For using Bellman–Ford first of all create weighted graph:
```go
dependencies := map[int][]graph.Length[int]{
    1: {graph.NewLength(2, 6), graph.NewLength(3, 1)},
    2: {graph.NewLength(4, 7)},
    3: {graph.NewLength(5, 1), graph.NewLength(6, 2)},
    4: {graph.NewLength(7, 8)},
    5: {graph.NewLength(7, -1)},
    6: {graph.NewLength(8, 3)},
    7: {graph.NewLength(8, 10)},
    8: {},
}
count := 0
edgeKeyGen := func() int {
    count++
    return count
}
creator := graph.NewWeightedGraphCreator(dependencies, edgeKeyGen)
weightedGraph, err := graph.NewWeightedGraphFromCreator(creator)
if err != nil {
    t.Fatal("error must be nil")
}
```
And call BellmanFord:
```go
lengths, err := graphutil.BellmanFord(startNodeKey, weightedGraph)
if err != nil {
    t.Fatal("error must be nil")
}
```
Result of BellmanFord algo is hashmap with the shortest weight between start node and each other node in graph.
Algorithm returns error if negative circle is found.
### Dijkstra
Dijkstra is an algorithm for finding the shortest paths between nodes in a weighted graph.
Remember, that you mustn't have negative weight in your graph.
For using Dijkstra first of all create weighted graph:
```go
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
creator := graph.NewWeightedGraphCreator(dependencies, edgeKeyGen)
weightedGraph, err := graph.NewWeightedGraphFromCreator(creator)
if err != nil {
    t.Fatal("error must be nil")
}
```
And call Dijkstra:
```go
lengths, err := graphutil.Dijkstra(startNodeKey, weightedGraph)
if err != nil {
    t.Fatal("error must be nil")
}
```
Result of Dijkstra algo is hashmap with the shortest weight between start node and each other node in graph.
Algorithm returns error if negative weight is found.
### Ford-Fulkerson
The Ford–Fulkerson method or Ford–Fulkerson algorithm (FFA) is a greedy algorithm that computes the maximum flow in a flow network.
For using Ford–Fulkerson first of all create weighted graph:
```go
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
```
And call Ford–Fulkerson:
```go
flow := graphutil.FordFulkerson(1, 8, weightedGraph)
```
Result of Ford–Fulkerson algo is float64 value of max flow in network between start and finish nodes.
Ford–Fulkerson is noexcept method. If flow doesn't exist, method returns zero flow value.