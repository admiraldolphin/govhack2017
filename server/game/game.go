package game

import (
	"fmt"
)

// Handle handles an action.
func (s *State) Handle(a *Action) error {
	s.Lock()
	defer s.Unlock()

	if a.Act == ActNoOp {
		s.notify()
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
			// TODO: shuffle deck, deal cards
		default:
			return fmt.Errorf("bad action for StateLobby [%d]", a.Act)
		}
	case StateInGame:
		switch a.Act {
		case ActPlayCard:
			// TODO: check the card is valid play
			// TODO: compute effects
		case ActDiscard:
			// TODO: check the card is valid discard
			// TODO: compute effects / draw new card
		default:
			return fmt.Errorf("bad action for StateInGame [%d]", a.Act)
		}
		// Advance whose-turn to the next player, and game clock
		n := s.nextPlayer(s.WhoseTurn)
		if n < s.WhoseTurn {
			s.Clock++
		}
		s.WhoseTurn = n

	case StateGameOver:
		switch a.Act {
		case ActReturnToLobby:
			// Anyone can return to the lobby when the game is over.
			s.State = StateLobby
		default:
			return fmt.Errorf("bad action for StateGameOver [%d]", a.Act)
		}
	}
	s.notify()
	return nil
}
