package model

import "time"

type Manifest[T any] struct {
	Type  string    `json:"type"`
	Node  string    `json:"node"`
	Time  time.Time `json:"time"`
	Model *T        `json:"model"`
}
