package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// CustomerInput defines input data for inputing and updating customer
type CustomerInput struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsContacted bool   `json:"contacted"`
}

// Validate validates customer input
func (c CustomerInput) Validate() error {
	return validation.ValidateStruct(&c,
		// Name is required and must have 5-127 characters
		validation.Field(&c.Name, validation.Required, validation.Length(5, 127)),
		// Role is required and must have 1-15 characters
		validation.Field(&c.Role, validation.Required, validation.Length(1, 15)),
		// Email is required and must be a valid email address
		validation.Field(&c.Email, validation.Required, is.Email),
		// Phone is required and must follow e164 standard
		validation.Field(&c.Phone, validation.Required, is.E164))
}
