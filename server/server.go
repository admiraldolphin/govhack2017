package main

import (
	"encoding/json"
	"log"
	"net"
)

type TODO struct{}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	stop := make(chan struct{})
	q := make(chan TODO, 10)
	go func() {
		d := json.NewDecoder(conn)
		for {
			var m TODO
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
			stateMu.Lock()
			defer stateMu.Unlock()
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
