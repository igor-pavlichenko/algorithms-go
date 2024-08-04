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

func InsertOnBST(node *Node[int], value int) {
	var newNode *Node[int] = &Node[int]{value: value}

	if node.value < value {
		// value is bigger so we need to walk to the right side
		if node.right != nil {
			InsertOnBST(node.right, value)
		} else {
			// but if right is null - we found a place to insert
			node.right = newNode
		}
	} else if node.value >= value {
		// value is equal or smaller so we need to walk to the left side
		if node.left != nil {
			InsertOnBST(node.left, value)
		} else {
			node.left = newNode
		}
	}
}
