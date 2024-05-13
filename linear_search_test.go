package main

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	haystack := []int{
		1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420,
	}

	if !linear_search(haystack, 1) ||
		!linear_search(haystack, 69) ||
		!linear_search(haystack, 69420) {
		t.Fatal("failed to find the needle")
	}

	if linear_search(haystack, 1336) ||
		linear_search(haystack, 69421) ||
		linear_search(haystack, 0) {
		t.Fatal("found a completely unrelated needle")
	}

}
