package game

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

// Statum is some made up latin.
type Statum int

// Broad states (statums).
const (
	StateLobby Statum = iota
	StateInGame
	StateGameOver
)

// State models the entire game state.
type State struct {
	State     Statum          `json:"state"`
	Players   map[int]*Player `json:"players"`
	Clock     int             `json:"clock"`
	WhoseTurn int             `json:"whose_turn"`

	// Fields for coordinating state.
	changedNote chan struct{}
	mu          sync.RWMutex
	nextID      int
}

// New returns a new game state.
func New() *State {
	return &State{
		Players:     make(map[int]*Player),
		changedNote: make(chan struct{}),
	}
}

// AddPlayer adds a player.
func (s *State) AddPlayer() (int, error) {
	s.Lock()
	defer s.Unlock()
	if s.State != StateLobby {
		return -1, fmt.Errorf("game not in lobby state [%d!=%d]", s.State, StateLobby)
	}
	id := s.nextID
	s.Players[id] = &Player{}
	s.nextID++
	s.notify()
	return id, nil
}

// RemovePlayer quits a player.
func (s *State) RemovePlayer(id int) error {
	s.Lock()
	defer s.Unlock()
	if s.Players[id] == nil {
		return fmt.Errorf("id %d not present", id)
	}
	delete(s.Players, id)

	switch len(s.Players) {
	case 1:
		if s.State == StateInGame {
			// If there's one player remaining, they win.
			s.State = StateGameOver
		}
	case 0:
		// If there are no players remaining, go back to lobby.
		s.State = StateLobby
	}
	s.notify()
	return nil
}

func (s *State) nextPlayer(after int) int {
	min, sup := (1<<31)-1, (1<<31)-1
	for id := range s.Players {
		if id < min {
			min = id
		}
		if id > after && id < sup {
			sup = id
		}
	}
	if sup == (1<<31)-1 {
		return min
	}
	return sup
}

// Changed returns a channel closed when the state has changed.
func (s *State) Changed() <-chan struct{} {
	s.RLock()
	defer s.RUnlock()
	return s.changedNote
}

// Dump writes the state to a writer in JSON.
func (s *State) Dump(w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	s.RLock()
	defer s.RUnlock()
	return enc.Encode(s)
}

// MUST GUARD WITH LOCK
func (s *State) notify() {
	close(s.changedNote)
	s.changedNote = make(chan struct{})
}

func (s *State) Lock()    { s.mu.Lock() }
func (s *State) Unlock()  { s.mu.Unlock() }
func (s *State) RLock()   { s.mu.RLock() }
func (s *State) RUnlock() { s.mu.RUnlock() }

// Player is the state relative to a particular player.
type Player struct {
	Name  string `json:"name"`
	Hand  Hand   `json:"hand"`
	Score int    `json:"score"`
}

// Hand is some cards that a player has.
type Hand struct {
	People  []PersonCard `json:"people"`
	Actions []ActionCard `json:"actions"`
}

// PersonCard models a game card.
type PersonCard struct {
	Name string `json:"name"`
	// TODO
}

// ActionCard models a game card.
type ActionCard struct {
	Name string `json:"name"`
	// TODO
}
