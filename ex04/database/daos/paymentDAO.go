package daos

import (
	"apprit/store/database/dtos"

	"gorm.io/gorm"
)

func GetPayments(db *gorm.DB) ([]dtos.PaymentDTO, error) {
	var payments []dtos.PaymentDTO
	return GetEntities(db, &payments)
}

func GetPaymentById(db *gorm.DB, id uint64) (dtos.PaymentDTO, error) {
	var paymentDTO dtos.PaymentDTO
	return GetEntityById(db, id, &paymentDTO)
}

func DeletePaymentById(db *gorm.DB, id uint64) error {
	var paymentDTO dtos.PaymentDTO
	return DeleteEntityById(db, id, &paymentDTO)
}

func AddPayment(db *gorm.DB, paymentDTO dtos.PaymentDTO) error {
	return AddEntity(db, &paymentDTO)
}

func ReplacePayment(db *gorm.DB, id uint64, paymentDTO dtos.PaymentDTO) error {
	newValues := map[string]interface{}{"total": paymentDTO.Total, "transactiondtoid": paymentDTO.TransactionDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.PaymentDTO{})
}
