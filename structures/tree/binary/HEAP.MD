# Heap data structure

is a binary tree where every child and grand child is smaller (**MaxHeap**), or larger (**MinHeap**) than the current node

---

- it is self balancing
- it can be used for priority
- there is no traversing the tree
- **insertion**
  - insert it at last spot
  - check if heap conditions are met
  - if not, bubble it up and check again
- **deletion** = pop = removing root node
  - after deleting a node, we have an empty spot in our tree
  - so we find the last node, delete it, insert it in the empty spot
  - heapify down
  - in MinHeap we get the smallest of our two children and compare it to that
  - if it's larger than the smallest child - swap & repeat
- nodes are inserted from left to right, and you cannot have empty spaces between nodes at the same level
- is stored using an array, indexing happens from left-right horizontally
- navigating to child index formula:
  - left child: 2 x currIndex + 1
  - right child: 2 x currIndex + 2
- navigating from any child to parent
  - currIndex / 2 (discarding the reminder)
- updating is trickier, we need to maintain a hash map with index to value pairs so we can understand if we need to bubble it up or down depending on new value. typically priority-queues do not consider updating
