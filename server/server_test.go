package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"testing"

	"github.com/admiraldolphin/govhack2017/server/game"
)

func TestGame(t *testing.T) {
	s := server{state: game.New()}
	if err := s.listenAndServe("localhost:0"); err != nil {
		t.Fatalf("Couldn't start game server: %v", err)
	}
	defer s.Close()

	conn1, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 0 to game server: %v", err)
	}
	defer conn1.Close()
	go io.Copy(ioutil.Discard, conn1)
	p1 := json.NewEncoder(conn1)

	conn2, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 1 to game server: %v", err)
	}
	defer conn2.Close()
	go io.Copy(ioutil.Discard, conn2)
	p2 := json.NewEncoder(conn2)

	// Play a game!
	g := []struct {
		player *json.Encoder
		action *game.Action
		want   *game.State
	}{
		{p2, &game.Action{Act: game.ActStartGame}, &game.State{State: game.StateLobby}},
		{p1, &game.Action{Act: game.ActPlayCard}, &game.State{State: game.StateInGame}},
		{p2, &game.Action{Act: game.ActDiscard}, &game.State{State: game.StateInGame}},
	}

	for i, p := range g {
		if err := p.player.Encode(p.action); err != nil {
			t.Errorf("Message %d [%v] got error %v", i, p.action, err)
		}
		// TODO: if got, want := s.state, p.want; ...
	}
}
