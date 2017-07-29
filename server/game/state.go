package game

import (
	"encoding/json"
	"io"
	"sync"
)

// Statum is some made up latin.
type Statum int

// Broad states (statums).
const (
	StateNoGame Statum = iota
	StateLobby
	StateInGame
	StateGameOver
)

// State models the entire game state.
type State struct {
	State   Statum   `json:"state"`
	Players []Player `json:"players"`

	// Non-JSON fields for coordinating state.
	changedNote chan struct{}
	mu          sync.RWMutex
}

// New returns a new game state.
func New() *State {
	return &State{
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
	s.RLock()
	defer s.RUnlock()
	return json.NewEncoder(w).Encode(s)
}

// Notify notifies all listeners (on the channel return from Changed) that the state has changed.
func (s *State) Notify() {
	s.Lock()
	close(s.changedNote)
	s.changedNote = make(chan struct{})
	s.Unlock()
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
