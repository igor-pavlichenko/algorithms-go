package search

import (
	"errors"
)

/*
Binary search is an efficient algorithm for finding an item from a sorted list of items.
It works by repeatedly dividing in half the portion of the list that could contain the item,
until you've narrowed down the possible locations to just one.

Worst-case complexity: O(log n)
Best complexity: O(1)
*/
func binarySearch(haystack []int, needle int) (int, error) {
	low := 0              // start is inclusive
	high := len(haystack) // end is exclusive, meaning we will go up to, but not including it

	for low < high {
		mid := low + (high-low)/2
		value := haystack[mid]

		if value == needle {
			return mid, nil
		} else if value > needle { // then throw out everything to the right
			high = mid
		} else if value < needle { // then throw away everything to the left
			low = mid + 1 // no need to check mid again
		}
	}

	return -1, errors.New("needle not found")
}
