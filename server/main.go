package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	gamePort = flag.Int("game_port", 23456, "Port for the game to listen on")
	httpPort = flag.Int("http_port", 23480, "Port the webserver listens on")
)

func main() {
	_, err := net.Listen("tcp", fmt.Sprintf(":%d", *gamePort))
	if err != nil {
		log.Fatalf("Couldn't listen on game port: %v", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil {
		log.Fatalf("Couldn't serve HTTP: %v", err)
	}
}
