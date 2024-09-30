package models

import "time"

type CPU struct {
	ID          int       `json:"id" db:"id"`
	Model       string    `json:"model" db:"model"`
	Cores       int       `json:"cores" db:"cores"`
	Threads     int       `json:"threads" db:"threads"`
	Cache       int       `json:"cache" db:"cache"`
	BaseClock   float64   `json:"base_clock" db:"base_clock"`
	MaxClock    float64   `json:"max_clock" db:"max_clock"`
	Rank        int       `json:"rank" db:"rank"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	InsertedAt  time.Time `json:"inserted_at" db:"inserted_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
