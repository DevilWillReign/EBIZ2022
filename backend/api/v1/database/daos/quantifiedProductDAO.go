package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetQuantifiedProducts(db *gorm.DB) ([]dtos.QuantifiedProductDTO, error) {
	quantifiedProducts := []dtos.QuantifiedProductDTO{}
	return GetEntities(db, &quantifiedProducts)
}

func GetQuantifiedProductById(db *gorm.DB, id uint64) (dtos.QuantifiedProductDTO, error) {
	quantifiedProductDTO := dtos.QuantifiedProductDTO{}
	return GetEntityById(db, id, &quantifiedProductDTO)
}

func DeleteQuantifiedProductById(db *gorm.DB, id uint64) error {
	quantifiedProductDTO := dtos.QuantifiedProductDTO{}
	return DeleteEntityById(db, id, &quantifiedProductDTO)
}

func AddQuantifiedProduct(db *gorm.DB, quantifiedProductDTO dtos.QuantifiedProductDTO) (dtos.QuantifiedProductDTO, error) {
	return AddEntity(db, &quantifiedProductDTO)
}

func ReplaceQuantifiedProduct(db *gorm.DB, id uint64, quantifiedProductDTO dtos.QuantifiedProductDTO) error {
	if _, err := GetQuantifiedProductById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"product_dto_id": quantifiedProductDTO.ProductDTOID, "quantity": quantifiedProductDTO.Quantity,
		"transaction_dto_id": quantifiedProductDTO.TransactionDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.QuantifiedProductDTO{})
}

func GetQuantifiedProductsByTransactionId(db *gorm.DB, transactionid uint64) ([]dtos.QuantifiedProductDTO, error) {
	quantifiedProducts := []dtos.QuantifiedProductDTO{}
	if err := db.Where("transaction_dto_id=?", transactionid).Find(&quantifiedProducts).Error; err != nil {
		return []dtos.QuantifiedProductDTO{}, err
	}
	return quantifiedProducts, nil
}

func AddAllQuantifiedProduct(db *gorm.DB, quantifiedProductDTOs []dtos.QuantifiedProductDTO) error {
	return AddAllEntities(db, &quantifiedProductDTOs)
}
