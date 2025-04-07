package main

import (
	"fmt"
	"time"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

// Queue class
type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// New method initializes queue
func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}

				if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitPass--
					queue.waitTicket--
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (queue *Queue) StartTicketIssue() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

// EndTicketEndIssue ends the ticket issue
func (queue *Queue) EndTicketEndIssue() {
	queue.message <- messageTicketEnd
}

// ticketIssue starts and ends the ticket issue
func ticketIssue(queue *Queue) {
	for {
		// Sleep up to 10 Millisecond.
		time.Sleep(10 * time.Millisecond)
		queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")

		// Sleep up to 2 Millisecond
		time.Sleep(2 * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		queue.EndTicketEndIssue()
	}
}

// StartPass ends the Pass Queue
func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

// EndPass ends the Pass Queue
func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

// passenger method strats and ends the pass Queue
func passenger(queue *Queue) {
	for {
		// sleep up to 10 Millisecond.
		time.Sleep(10 * time.Millisecond)
		queue.StartPass()
		// sleep up to 2 Millisecond.
		time.Sleep(2 * time.Millisecond)
		fmt.Println("Passenger ends")
		queue.EndPass()
	}
}

func main() {
	queue := Queue{}
	queue.New()
	fmt.Println("queue : ", queue)
	for i := 0; i < 10; i++ {
		go passenger(&queue)
	}

	for j := 0; j < 5; j++ {
		go ticketIssue(&queue)
	}

	select {}
}
