package main

import (
	"encoding/json"
	"net"
	"reflect"
	"testing"

	"github.com/admiraldolphin/govhack2017/server/game"
)

func TestGame(t *testing.T) {
	s := server{state: game.New()}
	r := &response{}
	if err := s.listenAndServe("localhost:0"); err != nil {
		t.Fatalf("Couldn't start game server: %v", err)
	}
	defer s.Close()

	conn1, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 0 to game server: %v", err)
	}
	defer conn1.Close()

	send1 := json.NewEncoder(conn1)
	recv1 := json.NewDecoder(conn1)

	// Should get a response immediately
	if err := recv1.Decode(r); err != nil {
		t.Fatalf("Couldn't get state immediately: %v", err)
	}
	if got, want := r.PlayerID, 0; got != want {
		t.Errorf("response.PlayerID = %d, want %d", got, want)
	}
	if got, want := r.State.State, game.StateLobby; got != want {
		t.Errorf("response.State.State = %d, want %d", got, want)
	}

	conn2, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 1 to game server: %v", err)
	}
	defer conn2.Close()

	send2 := json.NewEncoder(conn2)
	recv2 := json.NewDecoder(conn2)

	// Should get a response immediately
	if err := recv2.Decode(r); err != nil {
		t.Fatalf("Couldn't get state immediately: %v", err)
	}
	if got, want := r.PlayerID, 1; got != want {
		t.Errorf("response.PlayerID = %d, want %d", got, want)
	}
	if got, want := r.State.State, game.StateLobby; got != want {
		t.Errorf("response.State.State = %d, want %d", got, want)
	}

	// Play a game!
	g := []struct {
		send   *json.Encoder
		recv   *json.Decoder
		action *game.Action
		want   *game.State
	}{
		{
			send:   send2,
			recv:   recv2,
			action: &game.Action{Act: game.ActStartGame},
			want: &game.State{
				State: game.StateInGame,
				Players: []game.Player{
					{}, {},
				},
			},
		},
		{
			send:   send1,
			recv:   recv1,
			action: &game.Action{Act: game.ActPlayCard},
			want: &game.State{
				State: game.StateInGame,
				Players: []game.Player{
					{}, {},
				},
			},
		},
		{
			send:   send2,
			recv:   recv2,
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
		if err := p.send.Encode(p.action); err != nil {
			t.Fatalf("Message %d [%v] got error %v", i, p.action, err)
		}
		if err := p.recv.Decode(r); err != nil {
			t.Errorf("After message %d [%v]: got error %v", i, p.action, err)
		}

		if got, want := r.State.State, p.want.State; got != want {
			t.Errorf("After message %d [%v]: state.State = %v, want %v", i, p.action, got, want)
		}
		if got, want := r.State.Players, p.want.Players; !reflect.DeepEqual(got, want) {
			t.Errorf("After message %d [%v]: state.Players = %v, want %v", i, p.action, got, want)
		}
	}
}
