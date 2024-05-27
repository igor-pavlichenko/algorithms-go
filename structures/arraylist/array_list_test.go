package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAt(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[string](3)
	list.InsertAt("a", 0)
	list.InsertAt("b", 0)
	list.InsertAt("c", 0)
	list.InsertAt("d", 0)

	assert.Equal(4, list.length)
	assert.Equal(6, list.capacity)

	first, firstErr := list.Get(0)
	assert.Equal("d", first)
	assert.Nil(firstErr)

	last, lastErr := list.Get(3)
	assert.Equal("a", last)
	assert.Nil(lastErr)
}

func TestInsertAtError(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[string](3)
	res, err := list.InsertAt("a", 3)

	assert.Zero(res)
	assert.Error(err)
}

func TestAppend(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[int](3)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(2)

	assert.Equal(4, list.length)
	assert.Equal(6, list.capacity)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[int](3)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(2)

	val, err := list.Get(3)

	assert.Equal(2, val)
	assert.Nil(err)
}

func TestGetError(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[int](3)

	val0, err0 := list.Get(0)
	assert.Zero(val0)
	assert.Error(err0)

	list.Append(1)

	val1, err1 := list.Get(1)
	assert.Zero(val1)
	assert.Error(err1)

	val2, err2 := list.Get(-1)
	assert.Zero(val2)
	assert.Error(err2)
}

func TestToString(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[int](3)
	list.Append(1)
	list.Append(1)
	list.Append(1)
	list.Append(2)
	str := list.ToString()

	assert.Equal("[1] - [1] - [1] - [2] - [ ] - [ ] - [ ]", str)
}
