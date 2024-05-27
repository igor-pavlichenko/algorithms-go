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

func TestPrepend(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[string](3)
	list.Prepend("a")
	list.Prepend("b")
	list.Prepend("c")

	assert.Equal(3, list.length)
	assert.Equal(3, list.capacity)

	first, firstErr := list.Get(0)
	assert.Equal("c", first)
	assert.Nil(firstErr)

	last, lastErr := list.Get(2)
	assert.Equal("a", last)
	assert.Nil(lastErr)
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

func TestRemoveAt(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[string](3)
	list.Append("a")
	list.Append("b")
	list.Append("c")
	removed, removeErr := list.RemoveAt(1)

	assert.Equal(2, list.length)
	assert.Equal(3, list.capacity)
	assert.Equal("b", removed)
	assert.Nil(removeErr)

	item, err := list.Get(1)
	assert.Equal("c", item)
	assert.Nil(err)

	zVal, errBounds := list.Get(2)
	assert.Zero(zVal)
	assert.Error(errBounds)

	assert.Equal("[a] - [c] - [ ]", list.ToString())
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	list := NewArrayList[string](3)
	list.Append("a")
	list.Append("b")
	list.Append("c")
	removed, removeErr := list.Remove("b")

	assert.Equal(2, list.length)
	assert.Equal(3, list.capacity)
	assert.Equal("b", removed)
	assert.Nil(removeErr)

	item, err := list.Get(1)
	assert.Equal("c", item)
	assert.Nil(err)

	zVal, errBounds := list.Get(2)
	assert.Zero(zVal)
	assert.Error(errBounds)

	assert.Equal("[a] - [c] - [ ]", list.ToString())
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

	assert.Equal("[1] - [1] - [1] - [2] - [ ] - [ ]", str)
}
