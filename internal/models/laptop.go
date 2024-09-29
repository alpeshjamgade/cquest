package models

import "time"

type Laptop struct {
	Brand         string
	Model         string
	Processor     CPU
	GPU           GPU
	RAM           int
	SSD           int
	HDD           int
	UsbC          string
	UsbA          string
	HDMI          string
	Ethernet      string
	HeadphoneJack string
	Weight        float64
	PriceInRupees float64
	Rank          int // by decision
	ReleaseDate   time.Time
}
