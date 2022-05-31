package models

import (
	"apprit/store/api/v1/database/dtos"
)

type Payment struct {
	ID            uint             `json:"id"`
	PaymentType   dtos.PaymentType `json:"paymenttype" validate:"required"`
	TransactionID uint             `json:"transactionid" validate:"required"`
}

func (p *Payment) Equals(o Payment) bool {
	return p.TransactionID == o.TransactionID
}
