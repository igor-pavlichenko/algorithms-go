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
