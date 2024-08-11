package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree2 Node[int] = Node[int]{
	value: 20,
	right: &Node[int]{
		value: 50,
		right: &Node[int]{
			value: 100,
			right: nil,
			left:  nil,
		},
		left: &Node[int]{
			value: 30,
			right: nil,
			left: &Node[int]{
				value: 29,
				right: nil,
				left:  nil,
			},
		},
	},
	left: &Node[int]{
		value: 10,
		right: &Node[int]{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &Node[int]{
			value: 5,
			right: nil,
			left:  nil,
		},
	},
}

func TestCompare(t *testing.T) {
	assert := assert.New(t)

	tree := initializeTree()

	assert.True(Compare(tree, tree))
	assert.True(Compare(&tree2, &tree2))
	assert.True(Compare(nil, nil))
	assert.False(Compare(nil, tree))
	assert.False(Compare(tree, nil))
	assert.False(Compare(tree, &tree2))
}
