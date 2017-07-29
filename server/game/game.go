package game

// Handle handles an action.
func (s *State) Handle(a *Action) error {
	s.Lock()
	defer s.Unlock()
	st := s.State

	switch st {
	case StateLobby:
		switch a.Act {

		}
	case StateInGame:
		switch a.Act {

		}
	case StateGameOver:
		switch a.Act {

		}
	}
	s.notify()
	return nil
}
