package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateAmounts(n int) <-chan int {
	amounts := make(chan int) // creates an output channel
	go func() {
		defer close(amounts) // closes
		for i := 0; i < n; i++ {
			amounts <- rand.Intn(100) + 1
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return amounts // return the output channel

}

func main() {
	sales := generateAmounts(50)    // generates 50 amounts on the sales channel
	expenses := generateAmounts(40) // generates 40 amounts on the expenses channel
	endOfDayAmount := 0

	for sales != nil || expenses != nil { // continues to loop while there is a non-nil channel
		select {
		case sale, moreData := <-sales: // consumes the next amount and channel open flag from the sales channel
			if moreData {
				fmt.Println("Sale of : ", sale)
				endOfDayAmount += sale // adds the sales amount to the total end-of-day balance
			} else {
				sales = nil // if the channel has been cloesd , marks the channel as nil , disabling this select case
			}
		case expense, moreData := <-expenses: // consumes the next amount and channel open floag from the expenses channel
			if moreData {
				fmt.Println("Expense of : ", expense)
				endOfDayAmount -= expense // subtracts the expense amount from the total end-of-day balance
			} else {
				expenses = nil // if the channel has been closed, marks the channel as nil , disabling this select case
			}
		}
	}
	fmt.Println("End of day profit and loss : ", endOfDayAmount)
}
