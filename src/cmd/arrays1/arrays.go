package main

/* Playing around with arrays in Go because I keep not understanding when to pass slices
as refrence or value.

Use pointer if you want to modify the structure, the size, or the location in memory of the slice and pass on change to those who call the function.

Useful links:
https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b
*/
import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Array val before: %v\n", arr)
	updateArrayValue(arr)
	resizeArrayValue(arr)
	fmt.Printf("Array val after: %v\n", arr)

	fmt.Printf("Array ref before: %v\n", arr2)
	updateArrayRef(&arr2)
	resizeArrayRef(&arr2)
	fmt.Printf("Array ref after: %v\n", arr2)

}

func updateArrayValue(arr []int) {
	arr[0] = -999
}

func updateArrayRef(arr *[]int) {
	(*arr)[0] = -999
}

func resizeArrayValue(arr []int) {
	arr = append(arr, 999)
}

func resizeArrayRef(arr *[]int) {
	*arr = append(*arr, 999)
}
