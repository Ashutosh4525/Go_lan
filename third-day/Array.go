package main

import "fmt"

//numbered sequence of specific length
func main() {
	//array declaration
	//if creating array with length then array will be empty
	//if int-then zero,bool-then false, string-then empty space
	var nums [4]int

	//array length
	fmt.Println(len(nums))

	//adding value
	nums[0] = 1

	//printing by index
	fmt.Println(nums[0])

	//printing the array
	fmt.Println(nums)

	//declare it in single line with short hand
	nums1 := [3]int{1, 2, 3}

	fmt.Println(nums1)

	//2d arrays
	arr2d := [2][2]int{{3, 4}, {5, 3}}
	fmt.Println(arr2d)

	//In Go, arrays are primarily used when the exact size and type of the collection of elements are known at compile time and are not expected to change during the program's execution.
	//- fixed size, that is predictable
	//- Memory optimization
	//- Contant time access
}
