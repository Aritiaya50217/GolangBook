package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Printf("Error upgrading to WebSocket: %v ", err)
			return
		}

		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("Error reading from WebSocket: %v", err)
					return
				}

				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					log.Printf("Error writing to WebSocket: %v", err)
					break
				}
			}
			// Outside the loop, close the connection
			log.Println("Closing WebSocket connection")
			conn.Close()
		}()
	}))
}
