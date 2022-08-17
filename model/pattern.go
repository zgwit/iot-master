package model

type WithRunning[T any] struct {
	T
	Running bool `json:"running"`
}
