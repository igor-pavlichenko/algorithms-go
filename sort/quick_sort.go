package sort

/*
Quicksort is a divide-and-conquer algorithm.
It works by selecting a 'pivot' element from the array
and partitioning the other elements into two sub-arrays,
according to whether they are less than or greater than the pivot.
For this reason, it is sometimes called partition-exchange sort.
The sub-arrays are then sorted recursively. This can be done in-place,
requiring small additional amounts of memory to perform the sorting.

Usually split into 2 functions:

  - partition
    produces the pivot index and moves
    the items that are <= to one side and
    the items that are > to the other side

  - qs
    calls partition
    gets the pivot
    recalls quick_sort
    handles base cases

Complexity: O(n*log(n))
*/
func QuickSort(arr []int) {

}

// lo - abbreviation of LOW
// hi - abbreviation of HIGH, unusual thing - both LOW & HIGH are inclusive
func qs(arr []int, lo int, hi int) {

}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi] // there are many variations for to pivot, this is the simplest
	idx := lo - 1    // will be incremented whenever we find an element >= pivot

	// now we need to walk from "lo" to "hi" but not including "hi" because it's the pivot
	// and do weak sort on this sub-array, meaning put everything that's less than pivot to the left
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}
	// by now we moved everything that is <= pivot to beginning
	// now we need to move our pivot to idx position
	// so that we hold true to the rule of
	// "everything to the left of pivot needs to be <= pivot and everything right needs to be > pivot"
	idx++
	arr[hi] = arr[idx] // move this element to the end where pivot was
	arr[idx] = pivot   // move pivot to it's final position

	return idx // return new index of pivot
}
