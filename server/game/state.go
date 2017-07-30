package game

import (
	"bytes"
	"encoding/json"
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

	// Fields for managing the game.
	nextID   int
	nextNum  int
	baseDeck Deck
	deck     Deck

	// Fields for coordinating state.
	changedNote chan struct{}
	mu          sync.RWMutex
}

// New returns a new game state.
func New(deck Deck) *State {
	return &State{
		Players:     make(map[int]*Player),
		changedNote: make(chan struct{}),
		baseDeck:    deck,
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

// String - for debugging.
func (s *State) String() string {
	b := new(bytes.Buffer)
	s.Dump(b)
	return b.String()
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
	Name      string             `json:"name"`
	Hand      *HandState         `json:"hand"`
	Played    []*ActionCardState `json:"played"`
	Discarded []*ActionCardState `json:"discarded"`
	Score     int                `json:"score"`
}
