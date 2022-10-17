package model

import (
	"easyshop/utils"

	"gorm.io/gorm"
)

type Model interface {
	TableName() string
	Validate() error
	CreateModel() map[string]interface{}
}

func Save(model Model) error {
	db := utils.GetDB()
	tx := db.Begin()
	if err := CreateModel(tx, model); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func CreateModel(db *gorm.DB, model Model) error {
	if err := model.Validate(); err != nil {
		db.Rollback()
		return err
	}
	if err := db.Create(model).Error; err != nil {
		return err
	}
	return nil
}
