package dtos

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PaymentDTO struct {
	gorm.Model
	Total            decimal.Decimal `gorm:"not null"`
	TransactionDTOID uint            `gorm:"unique;not null"`
}
