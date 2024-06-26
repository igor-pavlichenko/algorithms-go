package stack

import "errors"

type Node[T any] struct {
	value T
	prev  *Node[T]
}

/*
The stack data structure is a linear data structure that accompanies a principle
known as LIFO (Last In First Out) or FILO (First In Last Out)

Search complexity:   O(n)
Insert complexity:   O(1)
Delete complexity:   O(1)
Space complexity:    O(n)
*/

type Stack[T any] struct {
	head   *Node[T]
	length int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	newNode := &Node[T]{value: item}
	if s.length < 1 {
		s.head = newNode
	} else {
		newNode.prev = s.head
		s.head = newNode
	}
	s.length++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.length < 1 {
		var zeroValue T
		return zeroValue, errors.New("Stack is empty")
	}

	s.length--
	oldHead := s.head
	s.head = s.head.prev
	return oldHead.value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.length < 1 {
		var value T
		return value, errors.New("Stack is empty")
	}

	return s.head.value, nil
}
