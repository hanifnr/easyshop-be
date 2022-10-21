package model

type Model interface {
	ID() int64
	TableName() string
	Validate() error
}
