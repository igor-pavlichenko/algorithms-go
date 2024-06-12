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

func qs(
	arr []int,
	lo int, // abbreviation of LOW
	hi int, // abbreviation of HIGH, unusual thing - both LOW & HIGH are inclusive
) {

}

func partition(arr []int, lo int, hi int) int {

}
