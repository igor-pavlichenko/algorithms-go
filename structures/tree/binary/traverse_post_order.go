package binarytree

/**
 * we are given the head/root of the tree, and we need to visit all the nodes and
 * return them in "in order"
 */
func TraversePostOrder(head *Node[int]) *[]int {
	var path []int
	return walkPost(head, &path)
}

func walkPost(curr *Node[int], path *[]int) *[]int {
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
	walkPost(curr.left, path)
	walkPost(curr.right, path)

	// post
	*path = append(*path, curr.value)
	return path
}
