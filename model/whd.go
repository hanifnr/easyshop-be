package model

type Whd struct {
	Id      int64   `json:"id"`
	WhId    int64   `json:"wh_id"`
	Sno     int     `json:"dno"`
	PurcId  int64   `json:"purc_id"`
	PurcDno int     `json:"purc_dno"`
	Qtywh   float32 `json:"qtywh" gorm:"DEFAULT:0"`
}

func (Whd) TableName() string {
	return "whd"
}
