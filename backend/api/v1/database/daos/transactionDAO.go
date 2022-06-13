package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetTransactions(db *gorm.DB) ([]dtos.TransactionDTO, error) {
	transactions := []dtos.TransactionDTO{}
	return GetEntities(db, &transactions)
}

func GetTransactionById(db *gorm.DB, id uint64) (dtos.TransactionDTO, error) {
	transactionDTO := dtos.TransactionDTO{}
	return GetEntityById(db, id, &transactionDTO)
}

func DeleteTransactionById(db *gorm.DB, id uint64) error {
	transactionDTO := dtos.TransactionDTO{}
	return DeleteEntityById(db, id, &transactionDTO)
}

func AddTransaction(db *gorm.DB, transactionDTO dtos.TransactionDTO) (dtos.TransactionDTO, error) {
	return AddEntity(db, &transactionDTO)
}

func ReplaceTransaction(db *gorm.DB, id uint64, transactionDTO dtos.TransactionDTO) error {
	if _, err := GetTransactionById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"total": transactionDTO.Total, "user_dto_id": transactionDTO.UserDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.TransactionDTO{})
}

func GetTransactionsByUserId(db *gorm.DB, userdtoid uint64) ([]dtos.TransactionDTO, error) {
	transactions := []dtos.TransactionDTO{}
	if err := db.Where("user_dto_id = ?", userdtoid).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return []dtos.TransactionDTO{}, nil
	}
	return transactions, nil
}

func GetUserTransactionById(db *gorm.DB, userdtoid uint64, id uint64) (dtos.TransactionDTO, error) {
	transactionDTO := dtos.TransactionDTO{}
	if err := db.Where("id = ? AND user_dto_id = ?", id, userdtoid).First(&transactionDTO).Error; err != nil {
		return dtos.TransactionDTO{}, nil
	}
	return transactionDTO, nil
}
