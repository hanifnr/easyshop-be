package model

import (
	"easyshop/utils"
	"encoding/json"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type TaxOffice struct {
	SenderId        string             `json:"senderId"`
	SenderIdType    string             `json:"senderIdType"`
	SendNo          string             `json:"sendNo"`
	ProceduresId    string             `json:"proceduresId"`
	Version         string             `json:"version"`
	Name            string             `json:"name"`
	Nation          string             `json:"nation"`
	Birth           string             `json:"birth"`
	Status          string             `json:"status"`
	LandDate        string             `json:"landDate"`
	DocType         string             `json:"docType"`
	PassportNo      string             `json:"passportNo"`
	LandingPermitNo string             `json:"landingPermitNo"`
	PortType        string             `json:"portType"`
	DepartDate      string             `json:"departDate"`
	Port            string             `json:"port"`
	Vehicle         string             `json:"vehicle"`
	ShopId          string             `json:"shopId"`
	ShopType        string             `json:"shopType"`
	ShopName        string             `json:"shopName"`
	ShopPlace       string             `json:"shopPlace"`
	BizName         string             `json:"bizName"`
	BizPlace        string             `json:"bizPlace"`
	SellDate        string             `json:"sellDate"`
	SellTime        string             `json:"sellTime"`
	ReceiptNo       string             `json:"receiptNo"`
	TransOrNot      string             `json:"transOrNot"`
	GeneralTotal    string             `json:"generalTotal"`
	ConsumTotal     string             `json:"consumTotal"`
	LqExemptOrNot   string             `json:"lqExemptOrNot"`
	LqTotal         string             `json:"lqTotal"`
	LqTotalNum      string             `json:"lqTotalNum"`
	Details         []*TaxOfficeDetail `json:"details"`
}

type TaxOfficeDetail struct {
	Serial       string `json:"serial"`
	GoodsType    string `json:"goodsType"`
	GoodsName    string `json:"goodsName"`
	JanCode      string `json:"janCode"`
	Number       string `json:"number"`
	Unit         string `json:"unit"`
	UnitPrice    string `json:"unitPrice"`
	Price        string `json:"price"`
	Reduced      string `json:"reduced"`
	LqIndividual string `json:"lqIndividual"`
}

func GetTaxOffice(id int64, generalTotal, consumTotal, lqTotal, lqTotalNum string, db *gorm.DB) (map[string]interface{}, utils.StatusReturn) {
	streetNumber := "238"
	place := "156-0054 Tokyo-To Setagaya-ku Sakuragaoka 2-13-8"

	order := &Order{}
	if err := Load(id, order, db); err != nil {
		return nil, utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	listOrderd := make([]*Orderd, 0)
	db.Where("order_id = ?", id).Order("dno").Find(&listOrderd)

	cust := &Cust{}
	if err := Load(order.CustId, cust, db); err != nil {
		return nil, utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	passport := &Passport{}
	if err := db.Where("cust_id = ?", cust.Id).Find(&passport).Error; err != nil {
		return nil, utils.StatusReturn{ErrCode: utils.ErrSQLLoad, Message: err.Error()}
	}
	var sellTime time.Time
	db.Select("date").Table("order_log").Where("order_id=? AND status_code = 'IC'", order.Id).Scan(&sellTime)
	listTaxOfficeDetail := make([]*TaxOfficeDetail, 0)

	var goodsType, reduced, lqIndividual int
	lqExempt := lqTotal != "0" || lqTotalNum != "0"
	if consumTotal != "0" {
		goodsType = 2
	} else if generalTotal != "0" {
		goodsType = 1
	}

	if consumTotal != "0" {
		reduced = 1
	} else {
		reduced = 0
	}

	if lqExempt {
		lqIndividual = 1
	} else {
		lqIndividual = 0
	}

	taxOffice := &TaxOffice{
		SenderId:     "101090300570501100001",
		SenderIdType: "0",
		SendNo:       utils.FormatTimeDetail(order.CreatedAt) + streetNumber,
		ProceduresId: "A",
		Version:      "1",
		Name:         cust.Name,
		Nation:       passport.Nationality,
		Birth:        utils.FormatTimeToyyyymmdd(&passport.BirthDate),
		Status:       passport.StatusResidence,
		LandDate:     utils.FormatTimeToyyyymmdd(order.ArrivalDate),
		DocType:      "1", //passport
		PassportNo:   passport.Number,
		// LandingPermitNo: "",
		PortType: "1", //Airport
		// DepartDate:      "",  //Unavailable
		// Port:            "",  //Departure point code
		// Vehicle:         "",  //flight number/ship code
		ShopId:        "101090300570501100001",
		ShopType:      "0",
		ShopName:      "Easy Shop Tokyo",
		ShopPlace:     place,
		BizName:       "Easy Shop LLC",
		BizPlace:      place,
		SellDate:      utils.FormatTimeToyyyymmdd(order.PickDate),
		SellTime:      utils.FormatTimeToyyyymmdd(&sellTime), //payment time
		ReceiptNo:     order.Trxno,
		TransOrNot:    "0",
		GeneralTotal:  generalTotal,
		ConsumTotal:   consumTotal,                  //consumable item?
		LqExemptOrNot: strconv.FormatBool(lqExempt), //liquor?
	}
	if lqExempt {
		taxOffice.LqTotal = lqTotal
		taxOffice.LqTotalNum = lqTotalNum
	}
	for _, orderd := range listOrderd {
		taxOfficeDetail := &TaxOfficeDetail{
			Serial:    strconv.Itoa(orderd.Dno),
			GoodsType: strconv.Itoa(goodsType), // general goods/consumable
			GoodsName: orderd.Name,
			JanCode:   orderd.ProductId,
			Number:    utils.Float64ToString(float64(orderd.Qty)),
			// Unit:         "",                                  //unit
			UnitPrice:    utils.Float64ToString(orderd.Price), //price per unit
			Price:        utils.Float64ToString(orderd.Price),
			Reduced:      strconv.Itoa(reduced),
			LqIndividual: strconv.Itoa(lqIndividual),
		}
		listTaxOfficeDetail = append(listTaxOfficeDetail, taxOfficeDetail)
	}
	taxOffice.Details = listTaxOfficeDetail
	var resultMap map[string]interface{}
	data, _ := json.Marshal(taxOffice)
	json.Unmarshal(data, &resultMap)

	return resultMap, utils.StatusReturnOK()
}
