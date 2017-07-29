package main

import (
	"encoding/json"
	"net"
	"reflect"
	"testing"

	"github.com/admiraldolphin/govhack2017/server/game"
)

func TestGame(t *testing.T) {
	s := server{state: game.New(testDeck)}
	r := &response{}
	if err := s.listenAndServe("localhost:0"); err != nil {
		t.Fatalf("Couldn't start game server: %v", err)
	}
	defer s.Close()

	conn0, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 0 to game server: %v", err)
	}
	defer conn0.Close()

	send0 := json.NewEncoder(conn0)
	recv0 := json.NewDecoder(conn0)

	// Should get a response immediately
	if err := recv0.Decode(r); err != nil {
		t.Fatalf("Couldn't get state immediately: %v", err)
	}
	if got, want := r.PlayerID, 0; got != want {
		t.Errorf("response.PlayerID = %d, want %d", got, want)
	}
	if got, want := r.State.State, game.StateLobby; got != want {
		t.Errorf("response.State.State = %d, want %d", got, want)
	}

	conn1, err := net.Dial("tcp", s.Addr().String())
	if err != nil {
		t.Fatalf("Couldn't connect player 1 to game server: %v", err)
	}
	defer conn1.Close()

	send1 := json.NewEncoder(conn1)
	recv1 := json.NewDecoder(conn1)

	// Should get a response immediately
	if err := recv1.Decode(r); err != nil {
		t.Fatalf("Couldn't get state immediately: %v", err)
	}
	if got, want := r.PlayerID, 1; got != want {
		t.Errorf("response.PlayerID = %d, want %d", got, want)
	}
	if got, want := r.State.State, game.StateLobby; got != want {
		t.Errorf("response.State.State = %d, want %d", got, want)
	}

	players := map[int]*game.Player{
		0: {
			Name: "Player 0",
			Hand: &game.HandState{
				Actions: []*game.ActionCardState{
					{Card: &game.ActionCard{Name: "action0"}},
					{Card: &game.ActionCard{Name: "action1"}},
					{Card: &game.ActionCard{Name: "action2"}},
					{Card: &game.ActionCard{Name: "action3"}},
					{Card: &game.ActionCard{Name: "action4"}},
					{Card: &game.ActionCard{Name: "action5"}},
					{Card: &game.ActionCard{Name: "action6"}},
				},
				People: []*game.PersonCardState{
					{Card: &game.PersonCard{Name: "person0"}},
					{Card: &game.PersonCard{Name: "person1"}},
					{Card: &game.PersonCard{Name: "person2"}},
					{Card: &game.PersonCard{Name: "person3"}},
					{Card: &game.PersonCard{Name: "person4"}},
				},
			},
		},
		1: {
			Name: "Player 1",
			Hand: &game.HandState{
				Actions: []*game.ActionCardState{
					{Card: &game.ActionCard{Name: "action7"}},
					{Card: &game.ActionCard{Name: "action8"}},
					{Card: &game.ActionCard{Name: "action9"}},
					{Card: &game.ActionCard{Name: "action10"}},
					{Card: &game.ActionCard{Name: "action11"}},
					{Card: &game.ActionCard{Name: "action12"}},
					{Card: &game.ActionCard{Name: "action13"}},
				},
				People: []*game.PersonCardState{
					{Card: &game.PersonCard{Name: "person5"}},
					{Card: &game.PersonCard{Name: "person6"}},
					{Card: &game.PersonCard{Name: "person7"}},
					{Card: &game.PersonCard{Name: "person8"}},
					{Card: &game.PersonCard{Name: "person9"}},
				},
			},
		},
	}

	// Play a game!
	g := []struct {
		send   *json.Encoder
		recv   *json.Decoder
		action *game.Action
		want   *game.State
	}{
		{
			// Player 1 starts the game
			send:   send1,
			recv:   recv1,
			action: &game.Action{Act: game.ActStartGame},
			want: &game.State{
				State:     game.StateInGame,
				Players:   players,
				WhoseTurn: 0,
				Clock:     0,
			},
		},
		{
			// Player 0 plays a card
			send:   send0,
			recv:   recv0,
			action: &game.Action{Act: game.ActPlayCard},
			want: &game.State{
				State:     game.StateInGame,
				Players:   players,
				WhoseTurn: 1,
				Clock:     0,
			},
		},
		{
			// Player 0 does a no-op
			send:   send0,
			recv:   recv0,
			action: &game.Action{Act: game.ActNoOp},
			want: &game.State{
				State:     game.StateInGame,
				Players:   players,
				WhoseTurn: 1,
				Clock:     0,
			},
		},
		{
			// Player 1 discards a card
			send:   send1,
			recv:   recv1,
			action: &game.Action{Act: game.ActDiscard},
			want: &game.State{
				State:     game.StateInGame,
				Players:   players,
				WhoseTurn: 0,
				Clock:     1,
			},
		},
	}

	for i, p := range g {
		if err := p.send.Encode(p.action); err != nil {
			t.Fatalf("Message %d [%v] got error %v", i, p.action, err)
		}
		if err := recv0.Decode(r); err != nil {
			t.Errorf("After message %d [%v]: recv0.Decode = error %v", i, p.action, err)
		}
		if err := recv1.Decode(r); err != nil {
			t.Errorf("After message %d [%v]: recv1.Decode = error %v", i, p.action, err)
		}

		if got, want := r.State.State, p.want.State; got != want {
			t.Errorf("After message %d [%v]: state.State = %#v, want %#v", i, p.action, got, want)
		}
		if got, want := r.State.Players, p.want.Players; !reflect.DeepEqual(got, want) {
			t.Errorf("After message %d [%v]: state.Players = %#v, want %#v", i, p.action, got, want)
		}
		if got, want := r.State.WhoseTurn, p.want.WhoseTurn; got != want {
			t.Errorf("After message %d [%v]: state.WhoseTurn = %#v, want %#v", i, p.action, got, want)
		}
		if got, want := r.State.Clock, p.want.Clock; got != want {
			t.Errorf("After message %d [%v]: state.Clock = %#v, want %#v", i, p.action, got, want)
		}
	}
}
