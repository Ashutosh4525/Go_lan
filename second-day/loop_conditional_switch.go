package main

// for => only construct in go for looping
func main() {
	//while
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {

	// }

	//classic for loop
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	//break same as all programming language
	// for i := 0; i < 3; i++ {
	// 	break
	// 	fmt.Println(i)
	// }

	//continue same as all other lang
	// for i := 0; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// in go v-1.22 new feature "range"
	// for i := range 3 {
	// 	fmt.Println(i)
	// }

	//if-else=>else or else if should be immediatly after the ended curve bracket
	// age := 13

	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("person is minor")
	// }

	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else if age >= 12 {
	// 	fmt.Print("teen")
	// } else {
	// 	fmt.Println("kid")
	// }

	// var role = "admin"
	// var hasPer = false

	// if role == "admin" && hasPer {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("no")
	// }

	// declaring variable in if
	//but it should be in short-hand syntax
	// if age := 18; age >= 18 {
	// 	fmt.Print("persion is adult ", age)
	// } else if age >= 12 {
	// 	fmt.Print("teen", age)
	// }

	//go does not have ternary operator use if-else

	// simple switch
	//in this break is not neseccary and default too
	//
	// i := 2
	// switch i {
	// case 1:
	// 	fmt.Print("one")
	// case 2:
	// 	fmt.Print("two")
	// default:
	// 	fmt.Print("other")
	// }

	//multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Print("weekend")
	// default:
	// 	fmt.Print("workday")
	// }

	//type switch
	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Print("integer")
	// 	case string:
	// 		fmt.Print("string")
	// 	default:
	// 		fmt.Print("other ", t)
	// 	}
	// }

	// whoAmI(2.3)
}
