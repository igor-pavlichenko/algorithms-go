package binarytree

type MinHeap struct {
	length int
	data   *[]int
}

func NewMinHeap() *MinHeap {
	arr := []int{}

	return &MinHeap{
		length: 0,
		data:   &arr,
	}
}

func parentIdx(idx int) int {
	return idx / 2
}

func leftChildIdx(idx int) int {
	return 2*idx + 1
}

func rightChildIdx(idx int) int {
	return 2*idx + 2
}

func (heap *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return // because we reached the top
	}

	curr := (*heap.data)[idx]
	parentIdx := parentIdx(idx)
	parent := (*heap.data)[parentIdx]

	if parent > curr {
		// swap current <-> parent
		(*heap.data)[idx] = parent
		(*heap.data)[parentIdx] = curr
		// keep heapifying up
		heap.heapifyUp(parentIdx)
	}
}

func (heap *MinHeap) heapifyDown(idx int) {
	leftIdx := leftChildIdx(idx)
	rightIdx := rightChildIdx(idx)

	// base cases
	// we reached the end, or
	// the first child-index we must check is outside of the array
	if idx >= heap.length || leftIdx >= heap.length {
		return
	}

	currValue := (*heap.data)[idx]
	leftValue := (*heap.data)[leftIdx]
	rightValue := (*heap.data)[rightIdx]

	// find the smallest of 2 children, and if our value is smaller - swap them
	if leftValue < rightValue && leftValue < currValue {
		(*heap.data)[idx] = leftValue
		(*heap.data)[leftIdx] = currValue
		// keep heapifying down
		heap.heapifyDown(leftIdx)
	} else if rightValue < leftValue && rightValue < currValue {
		(*heap.data)[idx] = rightValue
		(*heap.data)[rightIdx] = currValue
		// keep heapifying down
		heap.heapifyDown(rightIdx)
	}

}

// - **insertion**
// - insert it at last spot
// - check if heap conditions are met
// - if not, bubble it up and check again
// complexity: O(log n)
func (heap *MinHeap) Insert(value int) {
	(*heap.data) = append((*heap.data), value)
	heap.heapifyUp(heap.length)
	heap.length++
}

// - **deletion**
// - after deleting a node, we have an empty spot in our tree
// - so we find the last node, delete it, insert it in the empty spot
// - heapify down
// - in MinHeap we get the smallest of our two children and compare it to that
// - if it's larger than the smallest child - swap & repeat
// complexity: O(log n)
func (heap *MinHeap) Pop() int {
	return -1
}
