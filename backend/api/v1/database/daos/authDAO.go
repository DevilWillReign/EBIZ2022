package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetAuths(db *gorm.DB) ([]dtos.AuthDTO, error) {
	auths := []dtos.AuthDTO{}
	return GetEntities(db, &auths)
}

func GetAuthById(db *gorm.DB, id uint64) (dtos.AuthDTO, error) {
	authDTO := dtos.AuthDTO{}
	return GetEntityById(db, id, &authDTO)
}

func DeleteAuthById(db *gorm.DB, id uint64) error {
	authDTO := dtos.AuthDTO{}
	return DeleteEntityById(db, id, &authDTO)
}

func AddAuth(db *gorm.DB, authDTO dtos.AuthDTO) (dtos.AuthDTO, error) {
	return AddEntity(db, &authDTO)
}

func ReplaceAuth(db *gorm.DB, id uint64, authDTO dtos.AuthDTO) error {
	if _, err := GetAuthById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"auth_type": authDTO.Authtype, "user_dto_id": authDTO.UserDTOID}
	return ReplaceEntity(db, id, newValues, &dtos.AuthDTO{})
}
