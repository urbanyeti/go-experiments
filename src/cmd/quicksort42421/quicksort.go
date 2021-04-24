package main

/* Wrote this from memory without consulting notes
 */
import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{6, 1, 3, 5, 9, 2, 7, 4}
	fmt.Println(arr)
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)

	arr[right], arr[pivot] = arr[pivot], arr[right]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[right], arr[left] = arr[left], arr[right]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}
