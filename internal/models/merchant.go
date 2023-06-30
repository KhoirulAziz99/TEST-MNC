package models

//Struct digunakan untuk cetakan atau blueprint, kalau di bahasa lain mirip class


import "time"

type Merchant struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	NoTelephon string    `json:"no_telephon"`
	Category   string    `json:"category"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"craeted_at"`
	UpdetedAt  time.Time `json:"updeted_at"`
	Status     bool      `json:"status"`
}
