package models

//Struct digunakan untuk cetakan atau blueprint, kalau di bahasa lain mirip class


import "time"

type HistoryLog struct {
	ID         int       `json:"id"`
	CustomerId Customer  `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}
