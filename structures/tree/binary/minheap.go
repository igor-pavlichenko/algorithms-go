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

}

func (heap *MinHeap) heapifyDown(idx int) {

}

// - **insertion**
// - insert it at last spot
// - check if heap conditions are met
// - if not, bubble it up and check again
// complexity: O(log n)
func (heap *MinHeap) Insert(value int) {

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
