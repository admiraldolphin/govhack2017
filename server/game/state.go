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
