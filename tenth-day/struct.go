package main

import (
	"fmt"
	"time"
)

// complex structure
//grouping multiple filed
//structs destructures itself automatically

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// golang dosen't have class and constructor so we use function
func newOrder(id string, amount float32, status string) *order { //pointer for getting orignal values and returning value is order as struct
	//intialize in variable
	myOrder2 := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder2
}

// in this o.status becomes a copy of the given instance
// so we use "*" for getting the pointer and therfore changing the value
// if value changing occurs then always use "*"
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}
func main() {
	//if you don't set any field, default value is zero
	//int=>0, float=>0, string=>"", bool=>false
	// var order order=
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	//use of chnageStatus
	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)

	fmt.Println(myOrder.getAmount())
	//add value to the key in struct
	myOrder.createdAt = time.Now()

	//getting a field
	fmt.Println(myOrder.status)

	myOrder1 := order{
		id:        "2",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
	}
	myOrder.status = "paid"
	fmt.Println("Order struct", myOrder)
	fmt.Println("Order one", myOrder1)

	//geting value and passing it by using func
	myOrder2 := newOrder("3", 200, "unpaid")
	fmt.Println(myOrder2)
	fmt.Println(myOrder2.status)

	//struct
	language := struct {
		name   string
		isGood bool
	}{"goland", true}

	fmt.Println(language)
	fmt.Println(language.isGood)
}
