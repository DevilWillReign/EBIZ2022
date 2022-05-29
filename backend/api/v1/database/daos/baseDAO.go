package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetEntities[E dtos.Entity](db *gorm.DB, etities *[]E) ([]E, error) {
	if err := db.Find(etities).Error; err != nil {
		return *etities, err
	}
	return *etities, nil
}

func GetEntityById[E dtos.Entity](db *gorm.DB, id uint64, entity *E) (E, error) {
	if err := db.First(entity, id).Error; err != nil {
		return *entity, err
	}
	return *entity, nil
}

func DeleteEntityById[E dtos.Entity](db *gorm.DB, id uint64, entity *E) error {
	if err := db.Delete(entity, id).Error; err != nil {
		return err
	}
	return nil
}

func AddEntity[E dtos.Entity](db *gorm.DB, entity *E) (E, error) {
	if err := db.Create(entity).Error; err != nil {
		return *entity, err
	}
	return *entity, nil
}

func ReplaceEntity[E dtos.Entity](db *gorm.DB, id uint64, newValues map[string]interface{}, e *E) error {
	if err := db.Model(e).Select("*").Where("ID = ?", id).Updates(newValues).Error; err != nil {
		return err
	}
	return nil
}
