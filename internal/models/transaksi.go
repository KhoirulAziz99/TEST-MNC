package models

//Struct digunakan untuk cetakan atau blueprint, kalau di bahasa lain mirip class


import "time"

type Transaction struct {
	CustomerEmail string `json:"email"`
	MerchantName  string `json:"merchant"`
	Paid          int64  `json:"paid"`
}

type TransactionHistory struct {
	ID         int       `json:"id"`
	CustomerId Customer  `json:"customer_id"`
	MerchantId Merchant  `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	Amount     int64     `json:"amount"`
}
