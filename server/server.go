package main

import (
	"bufio"
	"context"
	"encoding/json"
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
	defer conn.Close()
	cctx, canc := context.WithCancel(ctx)
	go func() {
		if err := s.handleInbound(cctx, conn); err != nil {
			log.Printf("Handling inbound stream: %v", err)
			canc()
		}
	}()
	go func() {
		if err := s.handleOutbound(cctx, conn); err != nil {
			log.Printf("Handling outbout stream: %v", err)
			canc()
		}
	}()
}

func (s *server) handleInbound(ctx context.Context, conn net.Conn) error {
	rd := bufio.NewReader(conn)
	for {
		select {
		default:
		case <-ctx.Done():
			return ctx.Err()
		}
		var m game.Action
		msg, err := rd.ReadBytes('\n')
		if err != nil {
			return err
		}
		if err := json.Unmarshal(msg, &m); err != nil {
			return err
		}
		if err := s.state.Handle(&m); err != nil {
			return err
		}
	}
}

func (s *server) handleOutbound(ctx context.Context, conn net.Conn) error {
	for {
		select {
		case <-s.state.Changed():
			if err := s.state.Dump(conn); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
