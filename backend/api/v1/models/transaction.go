package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID                 uint                `json:"id"`
	CreatedAt          time.Time           `json:"createdat"`
	UserID             uint                `json:"userid" validate:"required"`
	Total              decimal.Decimal     `json:"total" validate:"required"`
	Payment            Payment             `json:"payments"`
	QuantifiedProducts []QuantifiedProduct `json:"quantifiedproducts"`
}
