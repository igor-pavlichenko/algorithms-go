package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var haystack = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)

	index, err := binarySearch(haystack, 0)
	assert.Equal(-1, index)
	assert.Error(err)

	index, err = binarySearch(haystack, 1)
	assert.Equal(0, index)
	assert.Nil(err)

	index, err = binarySearch(haystack, 69)
	assert.Equal(3, index)
	assert.Nil(err)

	index, err = binarySearch(haystack, 81)
	assert.Equal(5, index)
	assert.Nil(err)

	index, err = binarySearch(haystack, 1336)
	assert.Equal(-1, index)
	assert.Error(err)

	index, err = binarySearch(haystack, 69420)
	assert.Equal(10, index)
	assert.Nil(err)

	index, err = binarySearch(haystack, 69421)
	assert.Equal(-1, index)
	assert.Error(err)
}
