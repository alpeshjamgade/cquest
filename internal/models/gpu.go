package models

import "time"

type GPU struct {
	ID           int       `json:"id" db:"id"`
	Model        string    `json:"model" db:"model"`
	Memory       int       `json:"memory" db:"memory"`
	ClockSpeed   float64   `json:"clock_speed" db:"clock_speed"`
	MemoryType   string    `json:"memory_type" db:"memory_type"`
	TDP          float64   `json:"tdp" db:"tdp"`
	Architecture string    `json:"architecture" db:"architecture"`
	Rank         int       `json:"rank" db:"rank"`
	ReleaseDate  time.Time `json:"release_date" db:"release_date"`
	InsertedAt   time.Time `json:"inserted_at" db:"inserted_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
