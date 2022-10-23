package model

type Model interface {
	ID() int64
	TableName() string
	Validate() error
}

type Detail interface {
	SetMasterId(id int64)
}
