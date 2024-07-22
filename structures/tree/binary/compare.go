package binarytree

func Compare(a *Node[int], b *Node[int]) bool {
	// base cases
	if a == nil && b == nil { // structure check
		// if throughout our recursion we reach a nil node
		// on same position on both trees - they are equal
		return true
	}

	if a == nil || b == nil { // structure check
		// different nodes -> different trees
		return false
	}

	if a.value != b.value { // value check
		return false
	}

	// it will keep checking base cases on left subtree & right subtree and
	// if it reaches both nil ends returning true then they are equal
	return Compare(a.left, b.left) && Compare(a.right, b.right)
	// it's depth first because it will go all the way left first
}
