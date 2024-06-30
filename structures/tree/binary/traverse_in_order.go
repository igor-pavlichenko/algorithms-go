package binarytree

/**
 * we are given the head/root of the tree, and we need to visit all the nodes and
 * return them in "in order"
 */
func TraverseInOrder(head *Node[int]) *[]int {
	var path []int
	return walk(head, &path)
}

func walk(curr *Node[int], path *[]int) *[]int {
	// base case is when our node is non-existent
	if curr == nil {
		return path
	}

	// recursion steps:
	// 1 - pre
	// 2 - recurse
	// 3 - post

	// pre

	// recurse
	walk(curr.left, path)
	// (kinda pre)
	*path = append(*path, curr.value)
	// recurse again
	walk(curr.right, path)

	// post
	return path
}
