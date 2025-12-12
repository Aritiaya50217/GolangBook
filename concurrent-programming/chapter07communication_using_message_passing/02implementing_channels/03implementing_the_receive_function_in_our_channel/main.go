package implementingthereceivefunctioninourchannel

import (
	channelStruct "concurrent-programming/chapter07communication_using_message_passing/02implementing_channels/01creating_a_channel_with_semaphores"
	"container/list"
	"sync"
)

type Channel[M any] struct {
	capacitySema *channelStruct.Semaphore // capacity semaphore to block sender when the buffer is full
	sizeSema     *channelStruct.Semaphore // buffer size semaphore to block the receiver when the buffer is empty
	mutex        sync.Mutex               // mutex protecting our shared list data structure
	buffer       *list.List               // linked list to be used as a queue data structure
}

func (c *Channel[M]) Receive() M {
	c.capacitySema.Release() // Releases one permit from the capacity semaphore
	c.sizeSema.Acquire()     // Acquires one permit from the buffer size semaphore
	c.mutex.Lock()
	v := c.buffer.Remove(c.buffer.Front()).(M) // Removes one message from the buffer while protecting against race conditions using the mutex
	c.mutex.Unlock()
	return v // return the message's value
}
