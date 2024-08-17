package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)

	heap := NewMinHeap()

	assert.Equal(0, heap.length)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	assert.Equal(8, heap.length)

	// 1, 3, 4, 7, 69, 5, 420, 8,
	assert.Equal(1, (*heap.data)[0])
	assert.Equal(3, (*heap.data)[1])
	assert.Equal(4, (*heap.data)[2])
	assert.Equal(7, (*heap.data)[3])
	assert.Equal(69, (*heap.data)[4])
	assert.Equal(5, (*heap.data)[5])
	assert.Equal(420, (*heap.data)[6])
	assert.Equal(8, (*heap.data)[7])
}

func TestPop(t *testing.T) {
	assert := assert.New(t)

	heap := NewMinHeap()
	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	assert.Equal(8, heap.length)
	assert.Equal(1, heap.Pop())
	assert.Equal(3, heap.Pop())
	assert.Equal(4, heap.Pop())
	assert.Equal(5, heap.Pop())
	assert.Equal(4, heap.length)
	assert.Equal(7, heap.Pop())
	assert.Equal(8, heap.Pop())
	assert.Equal(69, heap.Pop())
	assert.Equal(420, heap.Pop())
	assert.Equal(0, heap.length)
}
