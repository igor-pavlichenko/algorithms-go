package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfsOnBST(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	assert.True(DfsOnBST(tree, 29))
	assert.True(DfsOnBST(tree, 7))
	assert.True(DfsOnBST(tree, 10))

	assert.False(DfsOnBST(tree, 1))
	assert.False(DfsOnBST(tree, 6))
	assert.False(DfsOnBST(tree, 17))
	assert.False(DfsOnBST(tree, 31))
	assert.False(DfsOnBST(tree, 99))
}

func TestInsertOnBST(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	// root >= value
	assert.Nil(tree.left.right.right)
	InsertOnBST(tree, 20)
	assert.Equal(20, tree.left.right.right.value)

	// root < value
	assert.Nil(tree.right.left.left.left)
	InsertOnBST(tree, 21)
	assert.Equal(21, tree.right.left.left.left.value)
}

func TestDeleteOnBST_NodeWithNoChildren(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	// removing a node with no children
	assert.Equal(7, tree.left.left.right.value)
	DeleteOnBST(tree, 7)
	assert.Equal(5, tree.left.left.value)
	assert.Nil(tree.left.left.right)
}

func TestDeleteOnBST_NodeWithOneChild(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	// removing a node with one child
	assert.Equal(5, tree.left.left.value)
	DeleteOnBST(tree, 5)
	assert.Equal(7, tree.left.left.value)
	assert.Nil(tree.left.left.right)
}

func TestDeleteOnBST_NodeWithTwoChildren(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	// removing a node with two children
	assert.Equal(50, tree.right.value)
	assert.Equal(45, tree.right.left.right.value)
	DeleteOnBST(tree, 50)
	assert.Equal(45, tree.right.value)
	assert.Nil(tree.right.left.right)
}

func TestDeleteOnBST_RootNode(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	// removing a root node with two children
	assert.Equal(20, tree.value)
	assert.Equal(15, tree.left.right.value)
	DeleteOnBST(tree, 20)
	assert.Equal(15, tree.value)
	assert.Nil(tree.left.right)
}
