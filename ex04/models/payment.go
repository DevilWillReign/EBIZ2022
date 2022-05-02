package models

import "github.com/shopspring/decimal"

type Payment struct {
	ID            uint
	Total         decimal.Decimal `json:"total" validate:"required"`
	TransactionID uint            `json:"transactionid" validate:"required"`
}

func (p *Payment) Equals(o Payment) bool {
	return p.Total == o.Total && p.TransactionID == o.TransactionID
}
