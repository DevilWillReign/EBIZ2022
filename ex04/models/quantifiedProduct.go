package models

import "github.com/shopspring/decimal"

type QuantifiedProduct struct {
	ID            uint
	Name          string          `json:"name" validate:"required"`
	Code          string          `json:"code" validate:"required"`
	Price         decimal.Decimal `json:"price" validate:"required"`
	CategoryID    uint            `json:"categoryid" validate:"required"`
	Quantity      uint            `json:"quantity" validate:"required"`
	TransactionID uint            `json:"transactioid" validate:"required"`
}
