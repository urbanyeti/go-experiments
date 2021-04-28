package main

import (
	"fmt"
)

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 62, 70, 100}
	fmt.Println(binarySearch2(62, items))
}

func binarySearch(needle int, haystack []int) bool {
	// Set up low and high boundaries
	low, high := 0, len(haystack)-1

	// While low is less than high boundary
	for low <= high {
		// Find the midpoint
		mid := (low + high) / 2

		// If value is greater than midpoint, move the low boundary to select upper half
		// otherwise, lower the high boundary to select lower half
		if needle > haystack[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// We're done. Did we run out of array to check or is the low boundary not the value
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	// Return true
	return true
}

func binarySearch2(n int, arr []int) bool {
	if len(arr) <= 1 {
		if len(arr) == 0 {
			return false
		}
		if arr[0] == n {
			return true
		}
		return false
	}

	left, right := 0, len(arr)-1
	mid := (left + right) / len(arr)

	if n == arr[mid] {
		return true
	}

	if n > arr[mid] {
		return binarySearch2(n, arr[mid+1:])
	}

	return binarySearch2(n, arr[:mid])
}
