package main

import (
	creatingachannelwithsemaphores "concurrent-programming/chapter07communication_using_message_passing/02implementing_channels/01creating_a_channel_with_semaphores"
	implementingthereceivefunctioninourchannel "concurrent-programming/chapter07communication_using_message_passing/02implementing_channels/03implementing_the_receive_function_in_our_channel"
	"container/list"
	"fmt"
	"sync"
	"time"
)

func receiver(messages *implementingthereceivefunctioninourchannel.Channel[int], wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = messages.Receive()
		fmt.Println("Received:", msg)
	}
	wGroup.Done()
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capacitySema: NewSemaphore(capacity), // creates a new semaphore with the number of permits equal to the input capacity
		sizeSema:     NewSemaphore(0),        // creates a new semaphore with the number of permits equal to 0
		buffer:       list.New(),             // creates a new , empty linked list
	}
}

func main() {
	channel := creatingachannelwithsemaphores.NewChannel[int](10)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go receiver(channel, &wGroup)
}

181