package main

import (
	"bufio"
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
	go func() {
		for {
			conn, err := gl.Accept()
			if err != nil {
				log.Printf("Couldn't accept connection: %v", err)
			}
			go s.handleConnection(conn)
		}
	}()
	return nil
}

func (s *server) handleConnection(conn net.Conn) {
	defer conn.Close()
	stop := make(chan struct{})
	go func() {
		rd := bufio.NewReader(conn)
		for {
			select {
			default:
			case <-stop:
				return
			}
			var m game.Action
			msg, err := rd.ReadBytes('\n')
			if err != nil {
				log.Printf("Couldn't read a message: %v", err)
				close(stop)
				return
			}
			if err := json.Unmarshal(msg, &m); err != nil {
				log.Printf("Couldn't decode message: %v", err)
				close(stop)
				return
			}
			if err := s.state.Handle(&m); err != nil {
				log.Printf("Couldn't handle message: %v", err)
				close(stop)
				return
			}
		}
	}()

	for {
		select {
		case <-s.state.Changed():
			if err := s.state.Dump(conn); err != nil {
				log.Printf("Couldn't encode state: %v", err)
				close(stop)
				return
			}
		case <-stop:
			// Stop the connection
			return
		}
	}
}
