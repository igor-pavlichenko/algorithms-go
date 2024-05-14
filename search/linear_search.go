package search

/*
A linear search is the simplest method of searching. It sequentially checks each
element of the list until a match is found or the whole list has been searched.

Worst-case complexity: O(n)
Average complexity: O(n)
Best-case performance: O(1)
*/
func linear_search(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		element := haystack[i]
		if element == needle {
			return true
		}
	}

	return false
}
