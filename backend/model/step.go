package model

type Step struct {
	Position  `json:",omitempty"`
	Timestamp int64 `json:"ts,omitempty"`
}
