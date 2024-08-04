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

func TestInsertOnBST(t *testing.T) {
	assert := assert.New(t)

	var treeRef = &tree

	// root >= value
	assert.Nil(tree.left.right.right)
	InsertOnBST(treeRef, 20)
	assert.Equal(20, tree.left.right.right.value)

	// root < value
	assert.Nil(tree.right.left.left.left)
	InsertOnBST(treeRef, 21)
	assert.Equal(21, tree.right.left.left.left.value)
}
