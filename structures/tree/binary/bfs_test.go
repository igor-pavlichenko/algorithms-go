package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	assert.True(BFS(tree, 29))
	assert.True(BFS(tree, 7))
	assert.False(BFS(tree, 8))
}
