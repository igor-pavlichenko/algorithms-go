package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsOnBST(t *testing.T) {
	assert := assert.New(t)

	assert.True(DfsOnBST(&tree, 29))
	assert.True(DfsOnBST(&tree, 7))
	assert.True(DfsOnBST(&tree, 10))

	assert.False(DfsOnBST(&tree, 1))
	assert.False(DfsOnBST(&tree, 6))
	assert.False(DfsOnBST(&tree, 17))
	assert.False(DfsOnBST(&tree, 31))
	assert.False(DfsOnBST(&tree, 99))
}
