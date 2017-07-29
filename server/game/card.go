package game

// PersonCard models a game card.
type PersonCard struct {
	Name   string   `json:"name"`
	Traits []*Trait `json:"traits"`
}

// ActionCard models a game card.
type ActionCard struct {
	Name   string `json:"name"`
	Trait  *Trait `json:"trait"`
	Played bool   `json:"played"`
}

// Trait is a condition which match person cards, and are
// played by action cards.
type Trait struct {
	Name  string `json:"name"`
	Death bool   `json:"death"`
}
