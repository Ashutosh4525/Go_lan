package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web services")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type ofresponse%T\n", res)
	// fmt.Println("respone: ", res)

	//response data
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Enter reading response ", err)
		return
	}
	fmt.Println("response ", string(data))
}
