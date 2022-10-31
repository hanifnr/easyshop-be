package model

import (
	"time"

	"gorm.io/gorm"
)

type Model interface {
	ID() int64
	TableName() string
	Validate() error
}

type Master interface {
	GetTrxno() string
	SetTrxno(trxno string)
}

type Detail interface {
	SetMasterId(id int64)
}

type TimeField interface {
	SetCreatedAt(time time.Time)
	SetUpdatedAt(time time.Time)
}

type ModelExt interface {
	SetValueModelExt(db *gorm.DB)
}
