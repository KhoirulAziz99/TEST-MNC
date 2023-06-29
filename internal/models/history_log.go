package models

import "time"

type HistoryLog struct {
	ID         int       `json:"id"`
	CustomerId Customer  `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}
