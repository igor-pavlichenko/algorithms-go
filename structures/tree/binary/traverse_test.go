package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func initializeTree() *Node[int] {
	return &Node[int]{
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
				right: &Node[int]{
					value: 45,
					right: nil,
					left:  nil,
				},
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
				right: &Node[int]{
					value: 7,
					right: nil,
					left:  nil,
				},
				left: nil,
			},
		},
	}
}

var tree Node[int] = Node[int]{
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
			right: &Node[int]{
				value: 45,
				right: nil,
				left:  nil,
			},
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
			right: &Node[int]{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}

func TestTraverseInOrder(t *testing.T) {
	assert := assert.New(t)

	expected := &[]int{
		5,
		7,
		10,
		15,
		20,
		29,
		30,
		45,
		50,
		100,
	}

	assert.Equal(expected, TraverseInOrder(&tree))
}

func TestTraversePreOrder(t *testing.T) {
	assert := assert.New(t)

	expected := &[]int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}

	assert.Equal(expected, TraversePreOrder(&tree))
}

func TestTraversePostOrder(t *testing.T) {
	assert := assert.New(t)

	expected := &[]int{
		7,
		5,
		15,
		10,
		29,
		45,
		30,
		100,
		50,
		20,
	}

	assert.Equal(expected, TraversePostOrder(&tree))
}
