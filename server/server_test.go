package main

import (
	"fmt"
	"net"
	"testing"

	"github.com/admiraldolphin/govhack2017/server/game"
)

func TestBasicQuery(t *testing.T) {
	s := server{state: game.New()}
	if err := s.listenAndServe("localhost:0"); err != nil {
		t.Fatalf("Couldn't start game server: %v", err)
	}
	defer s.Close()

	conn, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect to game server: %v", err)
	}

	// Send 10 empty messages and quit.
	for i := 0; i < 10; i++ {
		if _, err := fmt.Fprint(conn, "{}\n"); err != nil {
			t.Errorf("Message %d failed: %v", i, err)
		}
	}
}
