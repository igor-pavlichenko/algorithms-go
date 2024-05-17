package structures

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(node Node[T]) {

}

func (q *Queue[T]) Dequeue() T {

}

func (q *Queue[T]) Peek() T {

}
