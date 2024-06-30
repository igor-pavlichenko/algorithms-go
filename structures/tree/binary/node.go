package binarytree

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}
