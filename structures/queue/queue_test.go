package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Enqueue(55)
	queue.Enqueue(33)

	assert.Equal(2, queue.length)

	assert.Equal(55, queue.head.value)
	assert.Equal(33, queue.tail.value)
	assert.Equal(33, queue.head.next.value)

	peekVal, peekErr := queue.Peek()
	assert.Equal(55, peekVal)
	assert.Nil(peekErr)
}

func TestDequeue(t *testing.T) {
	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Enqueue(55)
	queue.Enqueue(33)
	value, err := queue.Dequeue()
	assert.Equal(55, value)
	assert.Nil(err)
	assert.Equal(1, queue.length)

	queue.Dequeue()
	assert.Equal(0, queue.length)

	peekVal, peekErr := queue.Peek()
	assert.Error(peekErr)
	assert.Equal(0, peekVal)

	dqVal, dqErr := queue.Dequeue()
	assert.Equal(0, dqVal)
	assert.Error(dqErr)
}

func TestPeek(t *testing.T) {
	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.head = &Node[int]{value: 11}
	queue.length++
	val, err := queue.Peek()
	assert.Equal(11, val)
	assert.Nil(err)
}
