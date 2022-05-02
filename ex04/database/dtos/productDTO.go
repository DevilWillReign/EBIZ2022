package dtos

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductDTO struct {
	gorm.Model
	Name          string          `gorm:"not null"`
	Code          string          `gorm:"unique;not null"`
	Price         decimal.Decimal `gorm:"not null"`
	CategoryDTOID uint            `gorm:"not null"`
}
