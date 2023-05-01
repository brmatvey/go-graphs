package graph

type Node[T comparable] interface {
	Key() T
	Children() []Node[T]
	AddChildren(node ...Node[T])
}

func NewNode[T comparable](key T, children ...Node[T]) Node[T] {
	return &node[T]{
		key:      key,
		children: children,
	}
}

type node[T comparable] struct {
	key      T
	children []Node[T]
}

func (n *node[T]) Key() T                      { return n.key }
func (n *node[T]) Children() []Node[T]         { return n.children }
func (n *node[T]) AddChildren(node ...Node[T]) { n.children = append(n.children, node...) }
