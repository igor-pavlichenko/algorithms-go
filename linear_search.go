package main

func linear_search(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		element := haystack[i]
		if element == needle {
			return true;
		}
	}

	return false;
}