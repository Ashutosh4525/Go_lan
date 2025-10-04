package main

import "fmt"

//closures
// func counter() func() int {
// 	var count int = 0

// 	return func() int {
// 		count += 1
// 		return count
// 	}
// }

//pointers
// func changeNum(num int) {
// 	num = 5
// 	//the original value is 1 but in this we changed the value
// 	// but values can only be changed for no primitive type
// 	// so in this new memory is created and num refrences to that new memory as it stores 5
// 	fmt.Println("In changeNum ", num)
// }

//by refrence
func changeNum1(num *int) {
	*num = 5
	fmt.Println("In changeNum ", *num)
}
func main() {

	//closures
	// increment := counter()
	// fmt.Println(increment())
	// fmt.Println(increment())

	//pointers
	num := 1

	// changeNum(num)
	changeNum1(&num)
	fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main ", num)
}
