package dtos

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionDTO struct {
	gorm.Model
	UserDTOID         uint                   `gorm:"not null"`
	Total             decimal.Decimal        `gorm:"not null"`
	PaymentDTO        PaymentDTO             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	QuantifiedProduct []QuantifiedProductDTO `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
