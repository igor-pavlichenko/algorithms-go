package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	assert.Equal("", list.ToString())

	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)
	assert.Equal("[1] - [2] - [3]", list.ToString())
}

func TestAppend(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(1)
	assert.Equal(1, list.head.value)
	assert.Equal(1, list.tail.value)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)
	assert.Equal(1, list.head.value)
	assert.Equal(3, list.tail.value)
	assert.Equal("[1] - [2] - [3]", list.ToString())
}

func TestPrepend(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(4)
	assert.Equal(4, list.head.value)
	assert.Equal(4, list.tail.value)
	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)
	assert.Equal(4, list.length)
	assert.Equal(1, list.head.value)
	assert.Equal(4, list.tail.value)
	assert.Equal("[1] - [2] - [3] - [4]", list.ToString())
}

func TestGetNodeAt(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	item, err := list.getNodeAt(1)
	assert.Equal(2, item.value)
	assert.Nil(err)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	item, err := list.Get(1)
	assert.Equal(2, item)
	assert.Nil(err)
}

func TestInsertAt(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	assert.Nil(list.InsertAt(9, 0))
	assert.Equal(1, list.length)
	assert.Equal("[9]", list.ToString())

	list.Append(0)
	list.Append(0)
	list.InsertAt(9, 1)
	assert.Equal("[9] - [9] - [0] - [0]", list.ToString())

	list.InsertAt(8, list.length)
	assert.Equal(8, list.tail.value)
	assert.Equal("[9] - [9] - [0] - [0] - [8]", list.ToString())

	err1 := list.InsertAt(1, -1)
	assert.ErrorIs(err1, ErrIndexOutOfBounds)

	err2 := list.InsertAt(1, list.length+1)
	assert.ErrorIs(err2, ErrIndexOutOfBounds)
}

func TestRemoveEmptyList(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	res1, err1 := list.Remove(2)
	assert.ErrorIs(err1, ErrListEmpty)
	assert.Zero(res1)
}

func TestRemoveNotFound(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(1)

	res1, err1 := list.Remove(2)
	assert.ErrorIs(err1, ErrItemNotFound)
	assert.Zero(res1)
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	assert.Equal(1, list.head.value)
	assert.Equal(5, list.tail.value)

	// removing head
	res1, err1 := list.Remove(1)
	assert.Nil(err1)
	assert.Equal(1, res1)
	assert.Equal(4, list.length)
	assert.Equal(2, list.head.value)
	assert.Equal("[2] - [3] - [4] - [5]", list.ToString())

	// removing tail
	res2, err2 := list.Remove(5)
	assert.Nil(err2)
	assert.Equal(5, res2)
	assert.Equal(3, list.length)
	assert.Equal(4, list.tail.value)
	assert.Equal("[2] - [3] - [4]", list.ToString())

	// removing until head & tail point to same node
	list.Remove(2)
	list.Remove(4)
	assert.Equal(1, list.length)
	assert.Equal(3, list.head.value)
	assert.Equal(3, list.tail.value)
	assert.Equal("[3]", list.ToString())
}

func TestRemoveAtEmptyList(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	res1, err1 := list.RemoveAt(0)
	assert.Zero(res1)
	assert.ErrorIs(err1, ErrListEmpty)
}

func TestRemoveAtOutOfBounds(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()
	list.Append(1)

	res1, err1 := list.RemoveAt(-1)
	assert.Zero(res1)
	assert.ErrorIs(err1, ErrIndexOutOfBounds)

	res2, err2 := list.RemoveAt(1)
	assert.Zero(res2)
	assert.ErrorIs(err2, ErrIndexOutOfBounds)
}

func TestRemoveAtMiddle(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)

	res, err := list.RemoveAt(1)
	assert.Equal(2, list.length)
	assert.Equal(2, res)
	assert.Nil(err)
	assert.Equal("[1] - [3]", list.ToString())
}

func TestRemoveAtHead(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)

	res, err := list.RemoveAt(0)
	assert.Nil(err)
	assert.Equal(1, res)
	assert.Equal(2, list.length)
	assert.Equal(2, list.head.value)
	assert.Equal("[2] - [3]", list.ToString())
}

func TestRemoveAtTail(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)

	res, err := list.RemoveAt(2)
	assert.Nil(err)
	assert.Equal(3, res)
	assert.Equal(2, list.length)
	assert.Equal(2, list.tail.value)
	assert.Equal("[1] - [2]", list.ToString())
}

func TestRemoveAtEverything(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)
	assert.Equal("[1] - [2] - [3]", list.ToString())

	list.RemoveAt(0)
	list.RemoveAt(0)
	list.RemoveAt(0)
	assert.Equal(0, list.length)
	assert.Nil(list.head)
	assert.Nil(list.tail)
	assert.Equal("", list.ToString())
}
