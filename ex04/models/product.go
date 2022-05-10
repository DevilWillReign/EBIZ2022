package models

import "github.com/shopspring/decimal"

type Product struct {
	ID         uint            `json:"id"`
	Name       string          `json:"name" validate:"required"`
	Code       string          `json:"code" validate:"required"`
	Price      decimal.Decimal `json:"price" validate:"required"`
	CategoryID uint            `json:"categoryid" validate:"required"`
}

func (p *Product) Equals(o Product) bool {
	return p.Code == o.Code && p.Price == o.Price && p.Name == o.Name && p.CategoryID == o.CategoryID
}
