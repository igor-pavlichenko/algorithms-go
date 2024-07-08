package binarytree

import "github.com/igor-pavlichenko/algorithms-go/structures/queue"

func BFS(head *Node[int], needle int) bool {
	q := queue.NewQueue[*Node[int]]()
	q.Enqueue(head)

	for q.Len() > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			return false
		}
		left := curr.left
		right := curr.right

		// search
		if curr.value == needle {
			return true
		}

		// keep traversing by enqueueing left & right of the node
		if left != nil {
			q.Enqueue(left)
		}
		if right != nil {
			q.Enqueue(right)
		}
		// by doing this the loop will keep iterating
	}

	// if we get to this point - we have traversed the whole tree
	// and didn't find the needle
	return false
}
