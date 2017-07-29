package game

// Broad states.
const (
	StateNoGame = iota
	StateLobby
	StateInGame
	StateGameOver
)

// State models the entire game state.
type State struct {
	State   int      `json:"state"`
	Players []Player `json:"players"`
}

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
}

// ActionCard models a game card.
type ActionCard struct {
	Name string `json:"name"`
}
