package main

import "fmt"

// Queue-Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// New method initializes with Order with priority, quantity, product,customerName
// func (order *Order) New(priority, quantity int, product, customerName string) {
// 	order.priority = priority
// 	order.quantity = quantity
// 	order.product = product
// 	order.customerName = customerName
// }

// Add method adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		for i, addedOrder := range *queue {
			fmt.Println("addedOrder : ", addedOrder) // first
			fmt.Println("order : ", order)           // second
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				fmt.Println("(*queue)[:i] : ", (*queue)[:i])
				fmt.Println("Queue{order} : ", Queue{order})
				fmt.Println("(*queue)[i:]...)... : ", (*queue)[i:])
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	queue := make(Queue, 0)
	order1 := Order{
		priority:     2,
		quantity:     20,
		product:      "Computer",
		customerName: "Greg White",
	}
	order2 := Order{
		priority:     4,
		quantity:     10,
		product:      "Monitor",
		customerName: "John Smith",
	}

	queue.Add(&order1)
	queue.Add(&order2)
	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
