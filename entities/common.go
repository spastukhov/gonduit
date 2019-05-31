package entities

// Cursor represents the pagination cursor on many responses.
type Cursor struct {
	Limit  uint64 `json:"limit"`
	After  string `json:"after"`
	Before string `json:"before"`
}