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

// complexity depends on height: O(h)
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

// complexity depends on height: O(h)
func DeleteOnBST(tree *Node[int], target int) *Node[int] {
	tree = delete(tree, target)
	return tree
}
func delete(node *Node[int], target int) *Node[int] {
	// base case 1 - empty tree
	if node == nil {
		return node
	}

	if target < node.value {
		// base case 2 - target < value
		// so we recurse to the left subtree
		node.left = delete(node.left, target)
		return node
	} else if target > node.value {
		// base case 3 - target > value
		// so we recurse to the right subtree
		node.right = delete(node.right, target)
		return node
	} else {
		// base case 4 - target == value

		// we found the node to be deleted
		// and we know that this code will be run in recursive calls
		// meaning that parent.left or parent.right will be assigned
		// the return value of this block

		if node.left == nil && node.right == nil {
			// case 1 - no children
			return nil // node deleted by simply assigning nil
		} else if node.left == nil {
			// case 2 - single child
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			// case 3 - two children
			// find a leaf node to replace the current node's value
			replacement := getLargestOnSmallSide(node.left)
			node.left = delete(node.left, replacement.value) // Update `node.left` with the result of the deletion
			node.value = replacement.value
			return node
		}
	}
}

func getLargestOnSmallSide(left *Node[int]) *Node[int] {
	curr := left

	for curr.right != nil {
		curr = curr.right
	}

	return curr
}
