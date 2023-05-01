# Simple implementation of trivial graph algorithms
You can use it, for instance, for educational, interview training etc.
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
directedGraph, err := graph.NewDirectedGraphFromCreator(graph.NewDirectedGraphCreator(dependencies))
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

weightedGraph, err := graph.NewWeightedGraph([]graph.Node[int]{parentNode, childNode}, []graph.Edge[int, int]{simpleEdge})
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
weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, uniqueKGen))
if err != nil {
    t.Fatal("err must be nil")
}
```
## Graph package
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
directedGraph, err := graph.NewDirectedGraphFromCreator(graph.NewDirectedGraphCreator(dependencies))
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
If you have circular dependencies if graph, graphutil.TopologicalSort returns error.
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

weightedGraph, err := graph.NewWeightedGraphFromCreator(graph.NewWeightedGraphCreator(dependencies, edgeKeyGen))
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