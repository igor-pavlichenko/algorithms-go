package structures

import "errors"

type Node[T any] struct {
	value T
	next  *Node[T]
}

/*
A queue is a linear data structure that stores the elements sequentially.
It uses the FIFO (First In First Out) approach for accessing elements.

Search complexity:   O(n)
Insert complexity:   O(1)
Delete complexity:   O(1)
Space complexity:    O(n)
*/
type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	var newNode *Node[T] = &Node[T]{value: item}
	if q.length < 1 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.length++
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.length < 1 {
		var value T
		return value, errors.New("Queue is empty")
	}

	return q.head.value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.length < 1 {
		var value T
		// can't return nill because that's not everything's zero-value
		// so I need to declare a variable of type T that will get the correct zero-value
		// however, depending on the purpose of our Queue, we might want to use pointers instead
		return value, errors.New("Queue is empty")
	}

	return q.head.value, nil
}
