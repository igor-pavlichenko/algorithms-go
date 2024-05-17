package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Enqueue(Node[int]{value: 55})
	queue.Enqueue(Node[int]{value: 33})
	assert.Equal(2, queue.length)
	assert.Equal(55, queue.Peek())
}

func TestDequeue(t *testing.T) {
	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Enqueue(Node[int]{value: 55})
	queue.Enqueue(Node[int]{value: 33})
	value := queue.Dequeue()
	assert.Equal(55, value)
	assert.Equal(1, queue.length)
	assert.Equal(33, queue.Peek())
}

func TestPeek(t *testing.T) {
	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.head = &Node[int]{value: 11}
	assert.Equal(11, queue.Peek())
}
