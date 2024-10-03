package models

import "time"

type Laptop struct {
	Id            int       `json:"id" db:"id"`
	Brand         string    `json:"brand" db:"brand"`
	Model         string    `json:"model" db:"model"`
	Processor     CPU       `json:"processor" db:"processor"`
	GPU           GPU       `json:"gpu" db:"gpu"`
	RAM           int       `json:"ram" db:"ram"`
	SSD           int       `json:"ssd" db:"ssd"`
	HDD           int       `json:"hdd" db:"hdd"`
	UsbC          string    `json:"usb_c" db:"usb_c"`
	UsbA          string    `json:"usb_a" db:"usb_a"`
	HDMI          string    `json:"hdmi" db:"hdmi"`
	Ethernet      string    `json:"ethernet" db:"ethernet"`
	HeadphoneJack string    `json:"headphone_jack" db:"headphone_jack"`
	Weight        float64   `json:"weight" db:"weight"`
	PriceInRupees float64   `json:"price" db:"price"`
	Rank          int       `json:"rank" db:"rank"` // by decision
	ReleaseDate   time.Time `json:"release_date" db:"release_date"`
	InsertedAt    time.Time `json:"inserted_at" db:"inserted_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
