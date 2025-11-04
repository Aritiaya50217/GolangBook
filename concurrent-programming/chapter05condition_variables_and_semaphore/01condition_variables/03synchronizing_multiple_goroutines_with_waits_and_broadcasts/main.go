package main

import (
	"fmt"
	"sync"
)

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int) {
	cond.L.Lock() // locks the mutex on the condition variable to avoid race condition
	fmt.Println(playerId, ": Connected")
	*playersRemaining-- // subtracts 1 from the shared remaining players variable
	if *playersRemaining == 0 {
		cond.Broadcast() // sends a broadcast when all players have connected
	}
	for *playersRemaining > 0 {
		fmt.Println(playerId, " : Waiting for more players")
		cond.Wait() // waits on a condition variable as long as there are more players to connect
	}
	cond.L.Unlock() // unlocks the mutex so that all goroutines can resume executin and start the game.
	fmt.Println("All players connected. Ready player", playerId)
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})            // creates a new condition variable
	playersInGame := 4                             // initializes the total number of players to be 4
	for playerId := 0; playerId <= 4; playerId++ { // starts goroutine sharing a condition variable ,players in game, and player ID
		go playerHandler(cond, &playersInGame, playerId) // Sleeps for a 1-second interval before the next player connects
	}
}
