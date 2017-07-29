package game

import (
	"fmt"
)

// Game parameters
const (
	endGameAtRound = 5

	ActionHandSize = 7
	PeopleHandSize = 5
)

// Handle handles an action.
func (s *State) Handle(a *Action, playerID int) error {
	s.Lock()
	defer s.Unlock()
	defer s.notify()

	// Everyone can always do nothing.
	if a.Act == ActNoOp {
		return nil
	}

	switch s.State {
	case StateLobby:
		switch a.Act {
		case ActStartGame:
			// Anyone can start the game if there are 2 or more players.
			if len(s.Players) < 2 {
				return fmt.Errorf("too few players for game [%d<2]", len(s.Players))
			}
			s.State = StateInGame
			s.startGame()

		default:
			return fmt.Errorf("bad action for StateLobby [%d]", a.Act)
		}
	case StateInGame:
		switch a.Act {
		case ActPlayCard, ActDiscard:
			if playerID != s.WhoseTurn {
				return fmt.Errorf("not your turn [%d!=%d]", playerID, s.WhoseTurn)
			}
			s.playOrDiscard(s.Players[playerID], a)

		default:
			return fmt.Errorf("bad action for StateInGame [%d]", a.Act)
		}
		s.advance()

	case StateGameOver:
		switch a.Act {
		case ActReturnToLobby:
			// Anyone can return to the lobby when the game is over.
			s.State = StateLobby
		default:
			return fmt.Errorf("bad action for StateGameOver [%d]", a.Act)
		}
	}
	return nil
}

// MUST GUARD WITH LOCK
func (s *State) playOrDiscard(p *Player, a *Action) error {
	if lim := len(p.Hand.Actions); a.Card < 0 || a.Card >= lim {
		return fmt.Errorf("card %d out of bounds [0, %d)", a.Card, lim)
	}
	cs := p.Hand.Actions[a.Card]

	switch a.Act {
	case ActPlayCard:
		cs.Played = true
		p.Played = append(p.Played, cs)

		s.tallyEffects(cs.Card)
	case ActDiscard:
		cs.Discarded = true
		p.Discarded = append(p.Discarded, cs)
	}

	nc := s.deck.DrawActions(1)
	if len(nc) == 0 {
		// Cover up the gap.
		copy(p.Hand.Actions[a.Card:], p.Hand.Actions[a.Card+1:])
		p.Hand.Actions = p.Hand.Actions[:len(p.Hand.Actions)-1]
	} else {
		// Replace card.
		p.Hand.Actions[a.Card] = nc[0]
	}
	return nil
}

// MUST GUARD WITH LOCK
func (s *State) tallyEffects(ac *ActionCard) {
	for _, p := range s.Players {
		for _, pc := range p.Hand.People {
			for ti, t := range pc.Card.Traits {
				if ac.Trait.Key != t.Key {
					continue
				}
				if pc.Dead {
					// Rule: Once dead, people don't accumulate points
					continue
				}
				// Record effects
				pc.CompletedTraits = append(pc.CompletedTraits, ti)
				pc.Score++ // Score attributed to this card
				pc.Dead = ac.Trait.Death
				if pc.Dead {
					// The person was just killed; add points to player.
					p.Score += pc.Score
				}
			}
		}
	}
}

// advance advances whose-turn to the next player, and game clock
// MUST GUARD WITH LOCK
func (s *State) advance() {
	n := s.nextPlayer(s.WhoseTurn)
	if n < s.WhoseTurn {
		s.Clock++
	}
	s.WhoseTurn = n
	if s.Clock == endGameAtRound {
		s.State = StateGameOver
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
	s.Players[id] = &Player{
		Name: fmt.Sprintf("Player %d", id),
	}
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

	default:
		// Go to the next player
		if s.WhoseTurn == id {
			s.advance()
		}
	}
	s.notify()
	return nil
}

// MUST GUARD WITH LOCK
func (s *State) nextPlayer(after int) int {
	const bigint = (1 << 31) - 1
	min, sup := bigint, bigint
	// It's gotta be linear in Players to find the next one when wrapping around.
	for id := range s.Players {
		if id < min {
			min = id
		}
		if id > after && id < sup {
			sup = id
		}
	}
	if sup == bigint {
		return min
	}
	return sup
}

// MUST GUARD WITH LOCK
func (s *State) startGame() {
	s.Clock = 0
	s.WhoseTurn = -1
	s.advance()

	s.deck = s.baseDeck.Instance()
	s.deck.Shuffle()
	for _, p := range s.Players {
		p.Hand = &HandState{
			Actions: s.deck.DrawActions(ActionHandSize),
			People:  s.deck.DrawPeople(PeopleHandSize),
		}
	}
}
