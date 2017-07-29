package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/admiraldolphin/govhack2017/server/game"
)

var (
	gamePort = flag.Int("game_port", 23456, "Port for the game to listen on")
	httpPort = flag.Int("http_port", 23480, "Port the webserver listens on")
)

func main() {
	s := server{state: game.New()}

	// Set up HTTP handlers
	http.HandleFunc("/helloz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!\n")
	})
	http.HandleFunc("/statusz", func(w http.ResponseWriter, r *http.Request) {
		s.state.Dump(w)
	})

	// Start listening on game port.
	gl, err := net.Listen("tcp", fmt.Sprintf(":%d", *gamePort))
	if err != nil {
		log.Fatalf("Couldn't listen on game port: %v", err)
	}
	go func() {
		for {
			conn, err := gl.Accept()
			if err != nil {
				log.Printf("Couldn't accept connection: %v", err)
			}
			go s.handleConnection(conn)
		}
	}()

	// Start listening on HTTP port; block.
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil {
		log.Fatalf("Couldn't serve HTTP: %v", err)
	}
}
