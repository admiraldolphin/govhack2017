package main

import (
	"encoding/json"
	"log"
	"net"

	"github.com/admiraldolphin/govhack2017/server/game"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	stop := make(chan struct{})
	q := make(chan game.Act, 10)
	go func() {
		d := json.NewDecoder(conn)
		for {
			select {
			default:
			case <-stop:
				return
			}
			var m game.Act
			if err := d.Decode(&m); err != nil {
				log.Printf("Couldn't decode message: %v", err)
				close(stop)
				return
			}
			q <- m
		}
	}()

	for {
		select {
		case m := <-q:
			// TODO: handle incoming message
			log.Printf("Got incoming message: %v", m)
		case <-state.Changed():
			state.RLock()
			defer state.RUnlock()
			if err := json.NewEncoder(conn).Encode(state); err != nil {
				log.Printf("Couldn't encode state: %v", err)
				close(stop)
				return
			}
		case <-stop:
			// Stop the connection
			return
		}
	}
}
