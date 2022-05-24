package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetQuantifiedProducts(db *gorm.DB) ([]dtos.QuantifiedProductDTO, error) {
	var quantifiedProducts []dtos.QuantifiedProductDTO
	return GetEntities(db, &quantifiedProducts)
}

func GetQuantifiedProductById(db *gorm.DB, id uint64) (dtos.QuantifiedProductDTO, error) {
	var quantifiedProductDTO dtos.QuantifiedProductDTO
	return GetEntityById(db, id, &quantifiedProductDTO)
}

func DeleteQuantifiedProductById(db *gorm.DB, id uint64) error {
	var quantifiedProductDTO dtos.QuantifiedProductDTO
	return DeleteEntityById(db, id, &quantifiedProductDTO)
}

func AddQuantifiedProduct(db *gorm.DB, quantifiedProductDTO dtos.QuantifiedProductDTO) error {
	return AddEntity(db, &quantifiedProductDTO)
}

func ReplaceQuantifiedProduct(db *gorm.DB, id uint64, quantifiedProductDTO dtos.QuantifiedProductDTO) error {
	if _, err := GetQuantifiedProductById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"productdtoid": quantifiedProductDTO.ProductDTOID, "quantity": quantifiedProductDTO.Quantity,
		"transactiondtoid": quantifiedProductDTO.TransactionDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.QuantifiedProductDTO{})
}
