package model

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
