package model

type NCount struct {
	Code   string
	Name   string
	Number int
	Length int
}

func (NCount) TableName() string {
	return "ncount"
}
