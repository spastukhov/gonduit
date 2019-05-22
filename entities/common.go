package entities

import (
	"encoding/json"
	"errors"
)

// Cursor represents the pagination cursor on many responses.
type Cursor struct {
	Limit  uint64 `json:"limit"`
	After  string `json:"after"`
	Before string `json:"before"`
}

// Epoch represents time.
type Epoch int64

// MarshalJSON marshals Epoch instance into JSON.
func (e *Epoch) MarshalJSON() ([]byte, error) {
	if e == nil {
		return nil, errors.New("epoch is nil")
	}
	return json.Marshal(int64(*e))
}
