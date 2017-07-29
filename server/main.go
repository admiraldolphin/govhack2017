package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/admiraldolphin/govhack2017/server/game"
	"github.com/admiraldolphin/govhack2017/server/load"
)

var (
	gamePort = flag.Int("game_port", 23456, "Port for the game to listen on")
	httpPort = flag.Int("http_port", 23480, "Port the webserver listens on")

	peopleJSON = flag.String("people", "/Users/josh/govhack2017/output.json", "File to load people from")
)

func main() {
	ppl, err := load.People(*peopleJSON)
	if err != nil {
		log.Printf("Couldn't load people: %v", err)
	}
	log.Printf("Loaded %d people", len(ppl))

	s := server{state: game.New(testDeck)}

	// Set up HTTP handlers
	http.HandleFunc("/helloz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, GovHack 2017!\n")
	})
	http.HandleFunc("/statusz", func(w http.ResponseWriter, r *http.Request) {
		s.state.Dump(w)
	})

	// Start listening on game port; don't block.
	if err := s.listenAndServe(fmt.Sprintf(":%d", *gamePort)); err != nil {
		log.Fatalf("Coudn't serve game: %v", err)
	}

	// Start listening on HTTP port; block.
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil {
		log.Fatalf("Couldn't serve HTTP: %v", err)
	}
}
