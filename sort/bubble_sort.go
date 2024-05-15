package sort

/*
Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm
that repeatedly steps through the input list element by element, comparing
the current element with the one after it, swapping their values if needed.

Worst-case complexity: O(n^2)
Best complexity: O(n)
Average complexity: O(n^2)
*/
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// remember to subtract "- 1 - i" because every iteration
		// the biggest number gets sorted to the end, so no need to keep comparing them
		for j := 0; j < len(arr)-1-i; j++ {
			current := arr[j]
			next := arr[j+1]
			if current > next {
				arr[j+1] = current
				arr[j] = next
			}
		}
	}
}
