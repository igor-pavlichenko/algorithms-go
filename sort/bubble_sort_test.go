package sort

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	assert := assert.New(t)

	unsortedArr1 := []int{453, 45, 9, 3423, 6, 1, 45345, 45454, 1212, 1211, 1004, 5, 8, 0}
	sortedArr1 := []int{0, 1, 5, 6, 8, 9, 45, 453, 1004, 1211, 1212, 3423, 45345, 45454}
	BubbleSort(unsortedArr1)
	assert.True(slices.Equal(unsortedArr1, sortedArr1))

	unsortedArr2 := []int{1, 0}
	sortedArr2 := []int{0, 1}
	BubbleSort(unsortedArr2)
	assert.True(slices.Equal(unsortedArr2, sortedArr2))

	unsortedArr3 := []int{1}
	sortedArr3 := []int{1}
	BubbleSort(unsortedArr3)
	assert.True(slices.Equal(unsortedArr3, sortedArr3))
}
