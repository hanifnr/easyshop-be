package model

import (
	"errors"

	"gorm.io/gorm"
)

func Create(model Model, db *gorm.DB) error {
	if err := db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func Load(id int64, model Model, db *gorm.DB) error {
	query := db.Where("id = ?", id).Find(model)
	if err := query.Error; err != nil {
		return err
	} else if rows := query.RowsAffected; rows == 0 {
		err := errors.New("data not found")
		return err
	}
	return nil
}

func Save(model Model, db *gorm.DB) error {
	if err := db.Save(model).Error; err != nil {
		return err
	}
	return nil
}
