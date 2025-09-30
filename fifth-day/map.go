package main

import (
	"fmt"
	"maps"
	"reflect"
)

// map-> object in js
func main() {
	//creating map

	m := make(map[string]string)

	//stting elements
	m["name"] = "golang"

	//get element
	fmt.Println(m["name"])

	//IMP:if key does not exist then map return zero
	fmt.Println(m["new"])

	m1 := make(map[string]int)
	m1["age"] = 21
	m1["number"] = 234482
	fmt.Println(m1["age"])
	fmt.Println(len(m1))

	//delete
	delete(m, "name")
	fmt.Println(m)
	fmt.Println(m1)

	//clear ->remove all from map
	clear(m1)
	fmt.Println(m1)

	//other mthod for map declaration
	m2 := map[int]string{1: "name", 2: "age"}
	fmt.Println(m2)

	//k would store the value if m2 has 2
	k, ok := m2[2]
	fmt.Println(k)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	m3 := map[int]string{1: "name", 2: "age"}
	m4 := map[int]string{1: "name", 2: "age"}

	//maps are object so
	fmt.Println(maps.Equal(m3, m4))
	//reflect.TypeOf
	fmt.Println(reflect.TypeOf(m3))
}
