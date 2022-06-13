package models

type Entity interface {
	Auth | Category | Payment | Product | QuantifiedProduct | Transaction | User | UserLogin | UserTransaction | PostCategory | PostTransaction
}

type ResponseArrayEntity[E Entity] struct {
	Elements []E `json:"elements"`
}
