package load

import (
	"encoding/json"
	"io"
	"os"
)

// People loads people from the JSON file.
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
