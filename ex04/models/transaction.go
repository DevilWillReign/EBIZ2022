package models

import "time"

type Transaction struct {
	ID                uint
	CreatedAt         time.Time
	UserID            uint `json:"userid" validate:"required"`
	Payment           Payment
	QuantifiedProduct []QuantifiedProduct
}
