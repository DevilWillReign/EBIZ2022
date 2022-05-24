package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetTransactions(db *gorm.DB) ([]dtos.TransactionDTO, error) {
	var transactions []dtos.TransactionDTO
	return GetEntities(db, &transactions)
}

func GetTransactionById(db *gorm.DB, id uint64) (dtos.TransactionDTO, error) {
	var transactionDTO dtos.TransactionDTO
	return GetEntityById(db, id, &transactionDTO)
}

func DeleteTransactionById(db *gorm.DB, id uint64) error {
	var transactionDTO dtos.TransactionDTO
	return DeleteEntityById(db, id, &transactionDTO)
}

func AddTransaction(db *gorm.DB, transactionDTO dtos.TransactionDTO) error {
	return AddEntity(db, &transactionDTO)
}

func ReplaceTransaction(db *gorm.DB, id uint64, transactionDTO dtos.TransactionDTO) error {
	if _, err := GetTransactionById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"userdtoid": transactionDTO.UserDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.TransactionDTO{})
}
