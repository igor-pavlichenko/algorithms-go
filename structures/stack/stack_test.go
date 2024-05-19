package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeek(t *testing.T) {
	assert := assert.New(t)

	stack := NewStack[int]()
	stack.head = &Node[int]{value: 11}
	stack.length++

	val, err := stack.Peek()
	assert.Equal(11, val)
	assert.Nil(err)
}

func TestPeekEmpty(t *testing.T) {
	assert := assert.New(t)

	stack := NewStack[int]()

	val, err := stack.Peek()
	assert.Equal(0, val)
	assert.Error(err)
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	assert.Equal(3, stack.length)

	val, err := stack.Peek()
	assert.Equal(3, val)
	assert.Nil(err)
}

func TestPopEmpty(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack[int]()

	assert.Equal(0, stack.length)

	val, err := stack.Pop()

	assert.Equal(0, stack.length)
	assert.Equal(0, val)
	assert.Error(err)
}

func TestPop(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	assert.Equal(5, stack.length)

	stack.Pop()
	val, err := stack.Pop()
	assert.Equal(3, stack.length)
	assert.Equal(4, val)
	assert.Nil(err)
}
