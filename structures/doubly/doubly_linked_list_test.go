package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	assert := assert.New(t)

	list := NewDoublyLinkedList[int]()
	list.append(1)
	list.append(2)
	list.append(3)

	assert.Equal(3, list.length)
	assert.Equal("[1] - [2] - [3]", list.length)

}
