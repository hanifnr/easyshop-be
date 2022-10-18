package model

type Model interface {
	TableName() string
	Validate() error
}
