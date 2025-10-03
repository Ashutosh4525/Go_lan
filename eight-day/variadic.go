package main

import "fmt"

func sum(nums ...int) int {
	add := 0
	for _, num := range nums {
		add = add + num
	}
	return add
}

func main() {
	nums := []int{3, 4, 2, 8}
	fmt.Println(sum(nums...))
	fmt.Println(1, 2, 3, 4, 5, 5, 6, "Hello", true)
	fmt.Println(sum(8, 73, 432, 817))
}
