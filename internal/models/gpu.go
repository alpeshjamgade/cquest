package models

import "time"

type GPU struct {
	Model        string
	Memory       int
	ClockSpeed   float64
	MemoryType   string
	TDP          float64
	Architecture string
	Rank         int // by decision
	ReleaseDate  time.Time
}
