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
			// Player 2 starts the game
			send:   send2,
			recv:   recv2,
			action: &game.Action{Act: game.ActStartGame},
			want: &game.State{
				State: game.StateInGame,
				Players: map[int]*game.Player{
					0: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
					1: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
				},
				WhoseTurn: 0,
				Clock:     0,
			},
		},
		{
			// Player 1 plays a card
			send:   send1,
			recv:   recv1,
			action: &game.Action{Act: game.ActPlayCard},
			want: &game.State{
				State: game.StateInGame,
				Players: map[int]*game.Player{
					0: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
					1: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
				},
				WhoseTurn: 1,
				Clock:     0,
			},
		},
		{
			// Player 1 does a no-op
			send:   send1,
			recv:   recv1,
			action: &game.Action{Act: game.ActNoOp},
			want: &game.State{
				State: game.StateInGame,
				Players: map[int]*game.Player{
					0: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
					1: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
				},
				WhoseTurn: 1,
				Clock:     0,
			},
		},
		{
			// Player 2 discards a card
			send:   send2,
			recv:   recv2,
			action: &game.Action{Act: game.ActDiscard},
			want: &game.State{
				State: game.StateInGame,
				Players: map[int]*game.Player{
					0: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
					1: {
						Hand: game.Hand{
							Actions: make([]game.ActionCard, game.ActionHandSize),
							People:  make([]game.PersonCard, game.PeopleHandSize),
						},
					},
				},
				WhoseTurn: 0,
				Clock:     1,
			},
		},
	}

	for i, p := range g {
		if err := p.send.Encode(p.action); err != nil {
			t.Fatalf("Message %d [%v] got error %v", i, p.action, err)
		}
		if err := recv1.Decode(r); err != nil {
			t.Errorf("After message %d [%v]: recv1.Decode = error %v", i, p.action, err)
		}
		if err := recv2.Decode(r); err != nil {
			t.Errorf("After message %d [%v]: recv2.Decode = error %v", i, p.action, err)
		}

		if got, want := r.State.State, p.want.State; got != want {
			t.Errorf("After message %d [%v]: state.State = %v, want %v", i, p.action, got, want)
		}
		if got, want := r.State.Players, p.want.Players; !reflect.DeepEqual(got, want) {
			t.Errorf("After message %d [%v]: state.Players = %v, want %v", i, p.action, got, want)
		}
		if got, want := r.State.WhoseTurn, p.want.WhoseTurn; got != want {
			t.Errorf("After message %d [%v]: state.WhoseTurn = %v, want %v", i, p.action, got, want)
		}
		if got, want := r.State.Clock, p.want.Clock; got != want {
			t.Errorf("After message %d [%v]: state.Clock = %v, want %v", i, p.action, got, want)
		}
	}
}
