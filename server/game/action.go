package game

// Act is different things a player can do.
type Act int

// The acts.
const (
	ActStartGame Act = iota
	ActPlayCard
	ActDiscard
	ActReturnToLobby
)

// Action is all info needed from the frontend to act.
type Action struct {
	Act    Act `json:"act"`
	Player int `json:"player"`
	Card   int `json:"card"`
}
