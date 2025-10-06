package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone int
}

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {

	// newCustomer := customer{
	// 	name:  "John",
	// 	phone: 972782361,
	// }
	// newOrder := order{
	// 	id:       1,
	// 	amount:   30.99,
	// 	status:   "recived",
	// 	customer: newCustomer,
	// }
	newOrder := order{
		id:     1,
		amount: 30.99,
		status: "recived",
		customer: customer{
			name:  "John",
			phone: 972782361,
		},
	}

	newOrder.customer.name = "Ashu"
	fmt.Println(newOrder)

}
