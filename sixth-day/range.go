package main

import (
	"fmt"
)

// it is used for iterating
func main() {
	nums := []int{1, 2, 3, 4, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	sum := 0
	for i, num := range nums {
		fmt.Println(i, num)
		sum += num
	}
	fmt.Println(sum)

	m := map[int]string{1: "john", 2: "john1"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//c=> gives unicode
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

	//this gives the string value
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
