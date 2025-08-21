package main

import (
	"log"
	"time"
)

func process(doneChannel chan bool) {
	time.Sleep(time.Second * 3)
	doneChannel <- true
}

func main() {
	// channels are nil util initialized with make
	doneChannel := make(chan bool)
	go process(doneChannel)

	tempBool := <-doneChannel
	log.Println(tempBool)

	// or to simply ignore the value but still wait
	// <- doneChannel

	// Start another process thread to run in background
	// and signal when done
	go process(doneChannel)

	// make channel non-blocking with select statement
	// this gives you ability to continue executing
	// even if no message is waiting in the channel
	var readerToExit bool
	for !readerToExit {
		select {
		case done := <-doneChannel: // ถ้ามีค่าเข้ามาใน channel
			log.Panicln("Done message received.", done)
			readerToExit = true // เปลี่ยน flag ให้หลุดจาก loop
		default:
			log.Println("No done signal yet. Waiting.")
			time.Sleep(time.Millisecond * 500) // พักครึ่งวินาทีแล้ววนใหม่
		}
	}

}
