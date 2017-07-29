package types

// Broad states.
const (
	StateNoGame = iota
	StateLobby
	StateInGame
	StateGameOver
)

// State models the entire game state.
type State struct {
	State int `json:"state"`
	Lobby struct {
	} `json:"lobby"`
	Game struct {
	} `json:"game"`
	GameOver struct {
	} `json:"game_over"`
}
