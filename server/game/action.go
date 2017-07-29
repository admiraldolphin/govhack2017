package game

// Act is different things a player can do.
type Act int

// The acts.
const (
	ActNoOp      Act = iota // Do nothing, just tell everybody the state
	ActStartGame            // Go from lobby state to ingame state
	ActPlayCard
	ActDiscard
	ActReturnToLobby // Go from game over state to lobby state
)

// Action is all info needed from the frontend to act.
type Action struct {
	Act  Act `json:"act"`
	Card int `json:"card"`
}
