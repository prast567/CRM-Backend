package model

type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsContacted bool   `json:"contacted"`
}
