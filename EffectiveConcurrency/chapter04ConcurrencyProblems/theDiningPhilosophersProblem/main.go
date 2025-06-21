package main

import (
	"fmt"
	"math/rand"
	"time"
)

func philosopher(index int, leftFork, rightFork chan bool) {
	for {
		// think for some time
		fmt.Printf("Philosopher %d is thinking\n", index)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))

		select {
		case <-leftFork:
			select {
			case <-rightFork:
				fmt.Printf("Philosopher %d is eating\n", index)
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
				rightFork <- true
			default:
			}
			leftFork <- true
		}

	}
}

func main() {

	var forks [5]chan bool
	for i := range forks {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}
	go philosopher(0, forks[4], forks[0])
	go philosopher(1, forks[0], forks[1])
	go philosopher(2, forks[1], forks[2])
	go philosopher(3, forks[2], forks[3])
	go philosopher(4, forks[3], forks[4])
	select {}
}
