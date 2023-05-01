package graph

type Edge[K, T comparable] interface {
	Key() K
	Weight() float64
	From() Node[T]
	To() Node[T]
}

func NewEdge[K, T comparable](key K, weight float64, from, to Node[T]) Edge[K, T] {
	return &edge[K, T]{
		key:    key,
		weight: weight,
		from:   from,
		to:     to,
	}
}

type edge[K, T comparable] struct {
	key    K
	weight float64
	from   Node[T]
	to     Node[T]
}

func (e *edge[K, T]) Key() K          { return e.key }
func (e *edge[K, T]) Weight() float64 { return e.weight }
func (e *edge[K, T]) From() Node[T]   { return e.from }
func (e *edge[K, T]) To() Node[T]     { return e.to }
