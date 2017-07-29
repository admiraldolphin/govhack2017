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
	dec := json.NewDecoder(r)
	for dec.More() {
		p := new(Person)
		if err := dec.Decode(p); err != nil {
			return nil, err
		}
		ppl = append(ppl, p)
	}
	return ppl, nil
}
