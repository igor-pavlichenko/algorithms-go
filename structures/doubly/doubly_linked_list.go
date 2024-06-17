package doubly

import (
	"errors"
	"fmt"
)

var (
	ErrListEmpty        = errors.New("list is empty")
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	ErrItemNotFound     = errors.New("item not found")
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// complexity: O(1)
func (list *DoublyLinkedList[T]) Append(item T) {
	newNode := Node[T]{value: item}

	list.length++
	if list.tail == nil { // list is empty
		list.head = &newNode
		list.tail = &newNode
		return
	}
	newNode.prev = list.tail
	list.tail.next = &newNode
	list.tail = &newNode
}

// complexity: O(1)
func (list *DoublyLinkedList[T]) Prepend(item T) {
	newNode := Node[T]{value: item}

	list.length++
	if list.head == nil {
		list.head = &newNode
		list.tail = &newNode
		return
	}
	newNode.next = list.head
	list.head.prev = &newNode
	list.head = &newNode
}

// complexity: O(n)
func (list *DoublyLinkedList[T]) getNodeAt(idx int) (*Node[T], error) {
	if list.length == 0 {
		return nil, ErrListEmpty
	}
	if idx < 0 || idx >= list.length {
		return nil, ErrIndexOutOfBounds
	}

	if idx == 0 {
		return list.head, nil
	}
	if idx == list.length-1 {
		return list.tail, nil
	}

	node := list.head
	for i := 0; i < idx; i++ {
		node = node.next
	}

	return node, nil
}

// complexity: O(n)
func (list *DoublyLinkedList[T]) Get(idx int) (T, error) {
	node, err := list.getNodeAt(idx)
	return node.value, err
}

// complexity: O(n)
func (list *DoublyLinkedList[T]) InsertAt(item T, idx int) error {
	if idx == 0 {
		list.Prepend(item)
		return nil
	}
	if idx == list.length {
		list.Append(item)
		return nil
	}

	if idx < 0 || idx > list.length {
		return ErrIndexOutOfBounds
	}

	node, err := list.getNodeAt(idx)
	if err != nil {
		return err
	}
	left := node.prev
	right := node

	newNode := Node[T]{value: item}
	newNode.prev = left
	newNode.next = right
	left.next = &newNode
	right.prev = &newNode

	list.length++

	return nil
}

// complexity: O(n)
func (list *DoublyLinkedList[T]) Remove(item T) (T, error) {
	if list.length < 1 {
		var zeroValue T
		return zeroValue, ErrListEmpty
	}

	curr := list.head
	for i := 0; i < list.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		var zeroValue T
		return zeroValue, ErrItemNotFound
	}

	left := curr.prev
	right := curr.next

	// update head if it was the head
	if curr == list.head {
		list.head = curr.next
	}
	// update tail if it was the tail
	if curr == list.tail {
		list.tail = curr.prev
	}

	// link the surrounding nodes
	if left != nil {
		left.next = right
	}
	if right != nil {
		right.prev = left
	}
	// unlink removed node
	curr.next = nil
	curr.prev = nil
	list.length--

	return curr.value, nil
}

func (list *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	node, err := list.getNodeAt(idx)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	list.length--
	if list.length == 0 {
		list.head = nil
		list.tail = nil
		return node.value, nil
	}

	left := node.prev
	right := node.next

	// update head if it was the head
	if node == list.head {
		list.head = node.next
	}
	// update tail if it was the tail
	if node == list.tail {
		list.tail = node.prev
	}

	// link the surrounding nodes
	if left != nil {
		left.next = right
	}
	if right != nil {
		right.prev = left
	}
	// unlink removed node
	node.next = nil
	node.prev = nil

	return node.value, nil
}

// complexity: O(n)
func (list *DoublyLinkedList[T]) ToString() string {
	str := ""
	curr := list.head
	for i := 0; i < list.length && curr != nil; i++ {
		str += "["
		str += fmt.Sprint(curr.value)
		str += "]"

		if i+1 < list.length {
			str += " - "
		}

		curr = curr.next
	}

	return str
}
