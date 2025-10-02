package main

import "fmt"

//in go function it is first class citizen
//third int is to say how many return value so type of the return value
func add(a int, b int) int { //in bracket we can write it as(a,b int)
	return a + b
}

func getLanguages() (string, string, string, int) {
	return "golang", "javascript", "c", 10
}

// func proccess(fn func(a int) int) {
// 	fn(1)
// }

func proccess() func(a int) int {
	return func(a int) int {
		return 4
	}
}

// this is main function called automatically by go
func main() {
	result := add(3, 2)
	fmt.Println(result)

	fmt.Println(getLanguages())
	// lang1, lang2, lang3, proficiency := getLanguages()
	// fmt.Println(lang1, lang2, lang3, proficiency)

	//if you dont want a input in function not to print you can use ,_
	lang1, _, _, proficiency := getLanguages()
	fmt.Println(lang1, proficiency)

	// fn := func(a int) int {
	// 	return 2
	// }

	//wrong
	// fn := proccess()
	// fn(7)
	// fmt.Println(fn)
}
