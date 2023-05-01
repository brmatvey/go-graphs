package graph

func NewDirectedGraphCreator[T comparable](structure map[T][]T) DirectedGraphCreator[T] {
	return DirectedGraphCreator[T]{structure: structure}
}

type DirectedGraphCreator[T comparable] struct {
	structure map[T][]T
}

func NewWeightedGraphCreator[K, T comparable](structure map[T][]Length[T], uniqueKGen func() K) WeightedGraphCreator[K, T] {
	return WeightedGraphCreator[K, T]{structure: structure, uniqueKGen: uniqueKGen}
}

type WeightedGraphCreator[K, T comparable] struct {
	structure  map[T][]Length[T]
	uniqueKGen func() K
}

func NewLength[T comparable](to T, weight float64) Length[T] {
	return Length[T]{to: to, weight: weight}
}

type Length[T comparable] struct {
	to     T
	weight float64
}

func newPath[T comparable](from, to T) path[T] {
	return path[T]{from: from, to: to}
}

type path[T comparable] struct {
	from T
	to   T
}
