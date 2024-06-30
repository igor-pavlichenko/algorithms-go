# Trees & terminology

-   **Root** - the most parent node. The first. Adam.
-   **Height** - the longest path from the root to the most child node.
-   **Binary Tree** - a tree that has at most 2 children & at least 0 children
-   **General Tree** - a tree with 0 or more children
-   **Binary Search Tree** - a tree that has a specific ordering to the nodes, with at most 2 children/ per node
-   **Leaves** - a node with no children
-   **Balanced Tree** - a tree is perfectly balanced when any node's left and right children have the same height
-   **Branching Factor** - the amount of children a tree has

# Traversing - O(n)

Visiting all the nodes of a tree (visiting means doing something with the value that a node holds, for example, printing it out)

Is done using recursion.

## Depth first search / traversal

-   visitNode()
-   recurseLeft()
-   recurseRight()

---

-   **Traversing pre-order** - Visiting the nodes before visiting the children. Root ends up in the beginning.
-   **Traversing in-order** - Visiting the nodes after visiting children on one side. Root ends up in the middle.
-   **Traversing post-order** - Visiting the nodes after visiting children on all sides. Root ends up in the end.
