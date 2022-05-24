package models

type Entity interface {
	Auth | Category | Payment | Product | QuantifiedProduct | Transaction | User
}
