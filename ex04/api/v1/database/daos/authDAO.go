package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetAuths(db *gorm.DB) ([]dtos.AuthDTO, error) {
	var auths []dtos.AuthDTO
	return GetEntities(db, &auths)
}

func GetAuthById(db *gorm.DB, id uint64) (dtos.AuthDTO, error) {
	var authDTO dtos.AuthDTO
	return GetEntityById(db, id, &authDTO)
}

func DeleteAuthById(db *gorm.DB, id uint64) error {
	var authDTO dtos.AuthDTO
	return DeleteEntityById(db, id, &authDTO)
}

func AddAuth(db *gorm.DB, authDTO dtos.AuthDTO) error {
	return AddEntity(db, &authDTO)
}

func ReplaceAuth(db *gorm.DB, id uint64, authDTO dtos.AuthDTO) error {
	newValues := map[string]interface{}{"authtype": authDTO.Authtype, "userdtoid": authDTO.UserDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.AuthDTO{})
}
