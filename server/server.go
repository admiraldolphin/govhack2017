package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/admiraldolphin/govhack2017/server/game"
)

type server struct {
	state *game.State

	net.Listener
}

func (s *server) listenAndServe(addr string) error {
	// Start listening on game port.
	gl, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.Listener = gl
	ctx := context.Background()
	go func() {
		for {
			conn, err := gl.Accept()
			if err != nil {
				log.Printf("Couldn't accept connection: %v", err)
			}
			s.handleConnection(ctx, conn)
		}
	}()
	return nil
}

func (s *server) handleConnection(ctx context.Context, conn net.Conn) {
	// Assign the player a number.
	id, err := s.state.AddPlayer()
	if err != nil {
		log.Printf("Can't add player: %v", err)
		conn.Close()
		return
	}
	cctx, canc := context.WithCancel(ctx)

	// Immediately send down the current state.
	if err := s.respond(cctx, conn, id); err != nil {
		log.Printf("Can't respond to new player: %v", err)
		canc()
		conn.Close()
		return
	}

	go func() {
		if err := s.handleInbound(cctx, conn, id); err != nil {
			log.Printf("Handling inbound stream: %v", err)
			canc()
		}
	}()
	go func() {
		if err := s.handleOutbound(cctx, conn, id); err != nil {
			log.Printf("Handling outbound stream: %v", err)
			canc()
		}
		conn.Close()
		s.state.RemovePlayer(id)
	}()
}

func (s *server) handleInbound(ctx context.Context, conn net.Conn, playerID int) error {
	rd := bufio.NewReader(conn)
	for {
		select {
		default:
		case <-ctx.Done():
			return ctx.Err()
		}
		m := new(game.Action)
		msg, err := rd.ReadBytes('\n')
		if err != nil {
			return err
		}
		if err := json.Unmarshal(msg, m); err != nil {
			log.Printf("Couldn't unmarshal incoming JSON but continuing: %v", err)
		}
		if err := s.state.Handle(m, playerID); err != nil {
			log.Printf("Couldn't handle action but continuing: %v", err)
		}
	}
}

func (s *server) handleOutbound(ctx context.Context, conn net.Conn, playerID int) error {
	for {
		select {
		case <-s.state.Changed():
			if err := s.respond(ctx, conn, playerID); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

type response struct {
	PlayerID int         `json:"you"`
	State    *game.State `json:"state"`
}

func (s *server) respond(ctx context.Context, w io.Writer, playerID int) error {
	s.state.RLock()
	defer s.state.RUnlock()
	resp := response{
		PlayerID: playerID,
		State:    s.state,
	}
	return json.NewEncoder(w).Encode(resp)
}

// TODO: make this unit test data only
var testDeck = &game.RiggedDeck{}

func init() {
	for i := 0; i < 50; i++ {
		testDeck.People = append(testDeck.People, &game.PersonCard{
			Name: fmt.Sprintf("person%d", i),
		})
	}
	for i := 0; i < 50; i++ {
		testDeck.Actions = append(testDeck.Actions, &game.ActionCard{
			Name: fmt.Sprintf("action%d", i),
		})
	}
}
