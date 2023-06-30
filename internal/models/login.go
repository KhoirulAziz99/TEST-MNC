package models

//Struct digunakan untuk cetakan atau blueprint, kalau di bahasa lain mirip class


type LoginCustomer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
