package game

import (
	"testing"
)

func TestNextPlayer(t *testing.T) {
	s := &State{
		Players: map[int]*Player{
			2:  {},
			3:  {},
			5:  {},
			7:  {},
			11: {},
		},
	}
	tests := []struct {
		after, want int
	}{
		{after: -1, want: 2},
		{after: 0, want: 2},
		{after: 1, want: 2},
		{after: 2, want: 3},
		{after: 3, want: 5},
		{after: 4, want: 5},
		{after: 5, want: 7},
		{after: 6, want: 7},
		{after: 7, want: 11},
		{after: 8, want: 11},
		{after: 9, want: 11},
		{after: 10, want: 11},
		{after: 11, want: 2},
	}

	for _, test := range tests {
		if got, want := s.nextPlayer(test.after), test.want; got != want {
			t.Errorf("nextPlayer(%d) = %d, want %d", test.after, got, want)
		}
	}
}
