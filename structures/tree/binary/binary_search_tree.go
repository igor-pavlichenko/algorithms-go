package binarytree

/*
binary search tree, also called an ordered or sorted binary tree,
is a rooted binary tree data structure with the key of each internal node
being greater than all the keys in the respective node's left subtree
and less than the ones in its right subtree.
*/

// dfs on a binary-search-tree - complexity: O(log n) - O(n)
func DfsOnBST(node *Node[int], needle int) bool {
	// base cases
	if node == nil {
		return false
	}
	if node.value == needle {
		return true
	}

	//recurse
	// if current nodes value is smaller than needle then check the right side
	if node.value < needle {
		return DfsOnBST(node.right, needle)
	}
	// else, means its greater than, so need to check left side
	return DfsOnBST(node.left, needle)
}
