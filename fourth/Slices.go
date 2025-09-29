package main

import (
	"fmt"
	"slices"
)

//slice =>dynamic array they can automatically resize
//all the drawback of array is removed
//most used
//usefull methods

func main() {
	//un-initialized slice is nil same as null but in go nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//make => inbuild function to make
	//in this size is 2
	// var nums1 = make([]int, 0, 5)

	//cap=>capacity max num of element can fit
	// fmt.Println(cap(nums1))
	// fmt.Println(len(nums1))
	//it is not nil as cap=2
	// fmt.Println(nums1 == nil)
	// nums1 := [3]int{}
	// nums1 = append(nums1, 3)
	// nums1 = append(nums1, 4)
	// nums1 = append(nums1, 1)

	// nums1[0] = 2
	// nums1[1] = 1
	// fmt.Println(nums1)
	// fmt.Println(len(nums1))
	// fmt.Println(cap(nums1))

	//copy function
	// var nums1 = make([]int, 0, 5)
	// fmt.Println(nums1)
	// nums1 = append(nums1, 2)
	// fmt.Println(nums1)
	// var nums2 = make([]int, len(nums1))
	// // nums1 = append(nums1, 2)
	// copy(nums2, nums1)
	// fmt.Println(nums1, nums2)

	//slice operator
	// var nums3 = []int{1, 2, 3}
	// fmt.Println(nums3[2:])

	//slice equal=> to compare array
	var num4 = []int{1, 2, 3}
	var num5 = []int{1, 2, 3}
	fmt.Println(slices.Equal(num4, num5))

	//2d array with slice
	var nums6 = [][]int{{2, 4, 6}, {3, 5, 7}}
	fmt.Println(nums6)
}
