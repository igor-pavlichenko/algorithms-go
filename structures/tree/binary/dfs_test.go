package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	assert.True(DFS(tree, 20))
	assert.True(DFS(tree, 50))
	assert.True(DFS(tree, 100))
	assert.True(DFS(tree, 30))
	assert.True(DFS(tree, 45))
	assert.True(DFS(tree, 29))
	assert.True(DFS(tree, 10))
	assert.True(DFS(tree, 15))
	assert.True(DFS(tree, 5))
	assert.True(DFS(tree, 7))

	assert.False(DFS(tree, 0))
	assert.False(DFS(tree, 1))
	assert.False(DFS(tree, 99))
	assert.False(DFS(tree, 60))
}
