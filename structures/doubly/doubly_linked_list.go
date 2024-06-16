package doubly

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
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
	list.tail.next = &newNode
	newNode.prev = list.tail
	list.tail = &newNode
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
