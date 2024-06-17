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
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.length)
	assert.Equal("[1] - [2] - [3]", list.ToString())
}

func TestPrepend(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(4)
	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)
	assert.Equal(4, list.length)
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

	list.InsertAt(9, list.length)
	assert.Equal("[9] - [9] - [0] - [0] - [9]", list.ToString())

	err1 := list.InsertAt(1, -1)
	assert.ErrorIs(err1, ErrIndexOutOfBounds)

	err2 := list.InsertAt(1, list.length+1)
	assert.ErrorIs(err2, ErrIndexOutOfBounds)
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	res1, err1 := list.Remove(2)
	assert.Equal(2, res1)
	assert.Nil(err1)
	assert.Equal(2, list.length)
	assert.Equal("[1] - [3]", list.ToString())

	res2, err2 := list.Remove(2)
	assert.Zero(res2)
	assert.ErrorIs(err2, ErrItemNotFound)
	assert.Equal("[1] - [3]", list.ToString())
}
