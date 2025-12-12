package implementingthesendfunctioninourchannel

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

func (c *Channel[M]) Send(message M) {
	c.capacitySema.Acquire() // acquires one permit from the capacity semaphore
	c.mutex.Lock()           // adds a message to the buffer queue while protecting against rce conditions by using a mutex
	c.buffer.PushBack(message)
	c.mutex.Unlock()
	c.sizeSema.Release() // release one permit from the buffer size semaphore
}
