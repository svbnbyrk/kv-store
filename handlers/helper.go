package handlers

import (
	"encoding/json"
	"io"
)

type SetModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//FromJSON is decoding data
func (p *SetModel) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//ToJSON is encoding data
func (p *SetModel) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
