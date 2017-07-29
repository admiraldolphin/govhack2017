package game

import (
	"math/rand"
)

// A RiggedDeck is a deck which can't be shuffled.
type RiggedDeck struct{ Hand }

// Instance makes a shallow copy of the deck.
func (r *RiggedDeck) Instance() Deck {
	r2 := *r
	return &r2
}

// Shuffle doesn't, in this case.
func (r *RiggedDeck) Shuffle() {}

// Hand is some cards that a player has, or that the "game" has (the "deck").
type Hand struct {
	People  []*PersonCard `json:"people"`
	Actions []*ActionCard `json:"actions"`
}

// HandState is like a hand, but tracks state of cards.
type HandState struct {
	People  []*PersonCardState `json:"people"`
	Actions []*ActionCardState `json:"actions"`
}

// Deck is what a deck can do.
type Deck interface {
	Instance() Deck
	Shuffle()
	DrawPeople(int) []*PersonCardState
	DrawActions(int) []*ActionCardState
}

// Instance makes a copy of the deck, so that dealing cards
// (destructive to slices) can be repeated.
func (h *Hand) Instance() Deck {
	h2 := *h
	return &h2
}

// Shuffle reorders the cards.
func (h *Hand) Shuffle() {
	for i, j := range rand.Perm(len(h.People)) {
		h.People[i], h.People[j] = h.People[j], h.People[i]
	}
	for i, j := range rand.Perm(len(h.Actions)) {
		h.Actions[i], h.Actions[j] = h.Actions[j], h.Actions[i]
	}
}

// DrawPeople removes the top card from the person deck and returns it,
// unless there are not enough, in which case it returns only the
// remaining card or nil.
func (h *Hand) DrawPeople(n int) (s []*PersonCardState) {
	var cs []*PersonCard
	defer func() {
		for _, c := range cs {
			s = append(s, c.New())
		}
	}()
	if len(h.People) <= n {
		cs, h.People = h.People, nil
		return
	}
	cs, h.People = h.People[:n], h.People[n:]
	return
}

// DrawActions does the same thing as DrawPerson but for action cards.
func (h *Hand) DrawActions(n int) (s []*ActionCardState) {
	var cs []*ActionCard
	defer func() {
		for _, c := range cs {
			s = append(s, c.New())
		}
	}()
	if len(h.People) <= n {
		cs, h.Actions = h.Actions, nil
		return
	}
	cs, h.Actions = h.Actions[:n], h.Actions[n:]
	return
}
