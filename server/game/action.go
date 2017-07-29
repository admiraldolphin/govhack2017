package game

// Act is different things a player can do.
type Act int

// The acts.
const (
	ActPlay Act = iota
	ActDiscard
)

// Action is all info needed from the frontend to act.
type Action struct {
	Player int `json:"player"`
	Act    Act `json:"act"`
	Card   int `json:"card"`
}
