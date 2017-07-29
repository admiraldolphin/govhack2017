package main

import (
	"encoding/json"
	//	"io"
	//	"io/ioutil"
	"net"
	"reflect"
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
	p1 := json.NewEncoder(conn1)

	conn2, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 1 to game server: %v", err)
	}
	defer conn2.Close()
	p2 := json.NewEncoder(conn2)

	// Play a game!
	g := []struct {
		conn   net.Conn
		player *json.Encoder
		action *game.Action
		want   *game.State
	}{
		{
			conn:   conn2,
			player: p2,
			action: &game.Action{Act: game.ActStartGame},
			want: &game.State{
				State: game.StateLobby,
				Players: []game.Player{
					{}, {},
				},
			},
		},
		{
			conn:   conn1,
			player: p1,
			action: &game.Action{Act: game.ActPlayCard},
			want: &game.State{
				State: game.StateInGame,
				Players: []game.Player{
					{}, {},
				},
			},
		},
		{
			conn:   conn2,
			player: p2,
			action: &game.Action{Act: game.ActDiscard},
			want: &game.State{
				State: game.StateInGame,
				Players: []game.Player{
					{}, {},
				},
			},
		},
	}

	for i, p := range g {
		if err := p.player.Encode(p.action); err != nil {
			t.Fatalf("Message %d [%v] got error %v", i, p.action, err)
		}
		s.state.RLock()
		if got, want := s.state.State, p.want.State; got != want {
			t.Errorf("After message %d [%v]: state.State = %v, want %v", i, p.action, got, want)
		}
		if got, want := s.state.Players, p.want.Players; !reflect.DeepEqual(got, want) {
			t.Errorf("After message %d [%v]: state.Players = %v, want %v", i, p.action, got, want)
		}
		s.state.RUnlock()
	}
}
