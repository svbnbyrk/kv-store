package handlers

import (
	"encoding/json"
	"io"
)
type SetModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//decoding json
func (p *SetModel) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//encoding json 
func ToJSON(w io.Writer, key string) error {
	e := json.NewEncoder(w)
	return e.Encode(key)
}