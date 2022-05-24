package dtos

type Entity interface {
	AuthDTO | CategoryDTO | PaymentDTO | ProductDTO | UserDTO | QuantifiedProductDTO | TransactionDTO
}
