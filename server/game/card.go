package game

import (
	"github.com/admiraldolphin/govhack2017/server/load"
)

// PersonCard models a game card.
type PersonCard struct {
	Name   string       `json:"name"`
	Traits []*Trait     `json:"traits"`
	Source *load.Person `json:"source"`
}

// New returns a fresh state for this card.
func (c *PersonCard) New() *PersonCardState {
	return &PersonCardState{Card: c}
}

// PersonCardState is state of a person card.
type PersonCardState struct {
	Card            *PersonCard `json:"card"`
	Dead            bool        `json:"dead"`
	CompletedTraits []int       `json:"completed_traits"`
	Score           int         `json:"score"`
}

// ActionCard models a game card.
type ActionCard struct {
	Name  string `json:"name"`
	Trait *Trait `json:"trait"`
}

// New returns a fresh state for this card.
func (c *ActionCard) New() *ActionCardState {
	return &ActionCardState{Card: c}
}

// ActionCardState is state of an action card.
type ActionCardState struct {
	Card      *ActionCard `json:"card"`
	Played    bool        `json:"played"`
	Discarded bool        `json:"discarded"`
}

// Trait is a condition which match person cards, and are
// played by action cards.
type Trait struct {
	Key            string  `json:"key"`
	Name           string  `json:"name"`
	Death          bool    `json:"death"`
	PeopleMatching float64 `json:"people_matching"`
}
