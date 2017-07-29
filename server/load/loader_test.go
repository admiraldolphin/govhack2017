package load

import (
	"strings"
	"testing"
)

func TestPeople(t *testing.T) {
	s := strings.NewReader("{}{}{}")
	ppl, err := people(s)
	if err != nil {
		t.Errorf("people = error %v", err)
	}
	if got, want := len(ppl), 3; got != want {
		t.Errorf("len(ppl) = %d, want %d", got, want)
	}
}
