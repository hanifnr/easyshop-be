package model

import "time"

type Cust struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CountryCode string    `json:"country_code"`
	PhoneNumber string    `json:"phone_number"`
	Passport    string    `json:"passport"`
	Status      string    `json:"status"`
	Isactive    bool      `gorm:"DEFAULT:TRUE" json:"isactive"`
	CreatedAt   time.Time `json:"created_at" gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"CURRENT_TIMESTAMP"`
}

func (Cust) TableName() string {
	return "usr"
}
