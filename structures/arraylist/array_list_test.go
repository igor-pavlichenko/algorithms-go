package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
