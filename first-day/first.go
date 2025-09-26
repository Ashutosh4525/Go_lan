package main

import "fmt"

// func main() {
// 	//first program
// 	fmt.Println("Hello World")
// }

// func main() {
// 	//simple values
// 	//int
// 	//in this operation can be also performed
// 	fmt.Println(1)

// 	//string
// 	fmt.Println("Hello")

// 	//bool
// 	fmt.Println(true)

// 	//float
// 	fmt.Println(10.5)
// }

// func main() {
// 	//when var declared it should be printed too as it gives error
// 	//this method can be used if you just want to declare the variable withot the value
// 	// var name string = "golang"

// 	//it can also decide datatype of the declared variable just like JS
// 	//if want to specify we can declare datatype with variable declaration
// 	// var name = "golang"

// 	//shorthand syntax - this is used when you want to declare value immediately
// 	name := "golang"
// 	fmt.Println(name)
// }

////const - it has same properties as of Js const

func main() {
	// const name = "hello"
	// fmt.Println(name)

	// const grouping - for declaring multiple const variable
	const (
		port = 2000
		host = "localhost"
	)
	fmt.Println(host, port)
}
