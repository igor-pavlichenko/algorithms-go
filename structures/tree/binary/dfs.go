package binarytree

func DFS(node *Node[int], needle int) bool {
	// base case is when our node is non-existent
	if node == nil {
		return false
	}

	// or it's value matches needle
	if node.value == needle {
		return true
	}

	return DFS(node.left, needle) || DFS(node.right, needle)
}
