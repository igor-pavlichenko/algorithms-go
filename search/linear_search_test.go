package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	haystack := []int{
		1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420,
	}

	if !LinearSearch(haystack, 1) ||
		!LinearSearch(haystack, 69) ||
		!LinearSearch(haystack, 69420) {
		t.Fatal("failed to find the needle")
	}

	if LinearSearch(haystack, 1336) ||
		LinearSearch(haystack, 69421) ||
		LinearSearch(haystack, 0) {
		t.Fatal("found a completely unrelated needle")
	}

}
