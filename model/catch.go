package model

import "time"

type Catch struct {
	ID        int       `json:"id"`
	FishType  string    `json:"fish_type"`
	Weight    float64   `json:"weight"`
	Length    float64   `json:"length"`
	Lure      string    `json:"lure"`
	CatchTime time.Time `json:"catch_time"`
}
