package load

import (
	"encoding/json"
	"io"
	"os"
)

// People loads people from person.json.
func People(path string) ([]*Person, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return people(f)
}

func people(r io.Reader) ([]*Person, error) {
	var ppl []*Person
	if err := json.NewDecoder(r).Decode(&ppl); err != nil {
		return nil, err
	}
	return ppl, nil
}

// Traits loads card traits from cards.json
func Traits(path string) (*Cards, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return traits(f)
}

func traits(r io.Reader) (*Cards, error) {
	var ct Cards
	if err := json.NewDecoder(r).Decode(&ct); err != nil {
		return nil, err
	}
	return &ct, nil
}
