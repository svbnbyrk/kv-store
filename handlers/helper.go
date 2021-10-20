package handlers

import (
	"encoding/json"
	"io"
)

type SetModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p *SetModel) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *SetModel) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
