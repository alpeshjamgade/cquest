package models

import "time"

type CPU struct {
	ID          string    `json:"id"`
	Model       string    `json:"model"`
	Cores       int       `json:"cores"`
	Threads     int       `json:"threads"`
	Cache       int       `json:"cache"`
	BaseClock   float64   `json:"base_clock"`
	MaxClock    float64   `json:"max_clock"`
	Rank        int       `json:"rank"`
	ReleaseDate time.Time `json:"release_date"`
}
