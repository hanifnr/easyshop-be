package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/scrape"
	"easyshop/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"gorm.io/gorm"
)

const (
	ORDER_NOTIFICATION = iota
	APPROVED_NOTIFICATION
	CANCELED_NOTIFICATION
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	CreateTransAction(orderController, w, r)
}

var UpdateOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	UpdateTransAction(orderController, w, r)
}

var ViewOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	ViewTransAction(orderController, w, r)
}

var DeleteOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	DeleteTransAction(orderController, w, r)
}

var ListOrder = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	ListTransAction(orderController, w, r)
}

var HandleOrder = func(w http.ResponseWriter, r *http.Request) {
	type Order struct {
		Id    int64
		Value string
		Note  string
	}
	order := &Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		data := utils.MessageErr(false, http.StatusBadRequest, err.Error())
		utils.RespondError(w, data, http.StatusBadRequest)
		return
	}
	orderController := &OrderController{}
	resp := orderController.HandleOrder(order.Id, order.Value, order.Note)
	utils.Respond(w, resp)
}

var TrackingNumber = func(w http.ResponseWriter, r *http.Request) {
	scuModel := &model.SingleStringColumnUpdate{}
	model.GetSingleColumnUpdate(w, r, scuModel, func() map[string]interface{} {
		orderController := &OrderController{}
		return orderController.TrackingNumber(scuModel.Id, scuModel.Value)
	})
}

var ShippingCost = func(w http.ResponseWriter, r *http.Request) {
	scuModel := &model.SingleNumericColumnUpdate{}
	model.GetSingleColumnUpdate(w, r, scuModel, func() map[string]interface{} {
		orderController := &OrderController{}
		return orderController.ShippingCost(scuModel.Id, scuModel.Value)
	})
}

var TaxDuty = func(w http.ResponseWriter, r *http.Request) {
	scuModel := &model.SingleStringColumnUpdate{}
	model.GetSingleColumnUpdate(w, r, scuModel, func() map[string]interface{} {
		orderController := &OrderController{}
		return orderController.TaxDuty(scuModel.Id, scuModel.Value)
	})
}

var ListOrderd = func(w http.ResponseWriter, r *http.Request) {
	param := utils.ProcessParam(r)
	orderController := &OrderController{}
	resp := orderController.ListDetail(param)
	utils.Respond(w, resp)
}

var UploadOrderProof = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	orderController.UploadOrderProof(w, r)
}

var LoadOrderProof = func(w http.ResponseWriter, r *http.Request) {
	orderController := &OrderController{}
	orderController.LoadOrderProof(w, r)
}

var ExportOrderXls = func(w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	orderController := &OrderController{}
	orderController.ExportOrderXls(id, w)
}

type OrderController struct {
	Order   model.Order    `json:"order"`
	Orderd  []model.Orderd `json:"orderd"`
	Details []model.Model  `json:"-"`
}

func (orderController *OrderController) MasterField() string {
	return "order_id"
}

func (orderController *OrderController) MasterModel() model.Model {
	return &orderController.Order
}

func (orderController *OrderController) DetailsModel() []model.Model {
	return orderController.Details
}

func (orderController *OrderController) CreateTrans() map[string]interface{} {
	if retval := CreateTrans(orderController, func(db *gorm.DB) error {
		order := &orderController.Order
		passport, err := utils.QueryStringValue(db.Select("number").Table("passport").Where("cust_id=?", order.CustId))
		if err != nil {
			return err
		}
		order.Trxno = GetOrderTrxno(order, db)
		order.StatusCode = "W"
		order.Passport = passport
		order.GrandTotal = order.Total - order.DiscAmount
		for i := range orderController.Orderd {
			orderd := &orderController.Orderd[i]
			orderd.Dno = i + 1
			orderController.Details = append(orderController.Details, orderd)
		}
		if retval := ProcessVoucher(order, db); retval.ErrCode != 0 {
			return errors.New(retval.Message)
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	} else {
		cust, notifOrder := getDataNotifOrder(orderController)
		SendEmailNotification(ORDER_NOTIFICATION, cust, *notifOrder)
		SendNewOrderPushNotification(&orderController.Order)
		return utils.MessageData(true, orderController)
	}
}

func (orderController *OrderController) ViewTrans(id int64) map[string]interface{} {
	if retval := ViewTrans(id, orderController, func(db *gorm.DB) error {
		err := db.Where("order_id = ?", id).Find(&orderController.Orderd).Error
		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
		}
		return err
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, orderController)
}

func (orderController *OrderController) ListTrans(param *utils.Param) map[string]interface{} {
	param.SetDefaultDelete()
	return ListJoinModel("\"order\"", "id ASC", make([]*model.Order, 0), param, func(query *gorm.DB) {
		query.Joins("JOIN status ON status_code = status.code")
	}, func(query *gorm.DB) {})
}

func (orderController *OrderController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(orderController, &model.Order{}, &model.Orderd{}, func(modelSrc, modelTemp model.Model, db *gorm.DB) error {
		orderSrc := modelSrc.(*model.Order)
		orderTemp := modelTemp.(*model.Order)

		orderSrc.CustId = orderTemp.CustId
		orderSrc.Date = orderTemp.Date
		orderSrc.PickDate = orderTemp.PickDate
		orderSrc.Total = orderTemp.Total
		orderSrc.VoucherId = orderTemp.VoucherId
		orderSrc.Disc = orderTemp.Disc
		orderSrc.DiscAmount = orderTemp.DiscAmount
		orderSrc.GrandTotal = orderTemp.Total + orderSrc.ShippingCost - orderTemp.DiscAmount
		orderSrc.Trxno = orderTemp.Trxno
		orderSrc.AddrId = orderTemp.AddrId
		orderSrc.ArrivalDate = orderTemp.ArrivalDate
		orderSrc.TaxAmount = orderTemp.TaxAmount

		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
		}
		return nil
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, orderController)

}

func (orderController *OrderController) DeleteTrans(id int64) map[string]interface{} {
	if retval := DeleteTrans(id, orderController, func() utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (orderController *OrderController) FNew() functions.SQLFunction {
	return &functions.FOrderNew{}
}

func (orderController *OrderController) FDelete() functions.SQLFunction {
	return &functions.FOrderDelete{}
}

func (orderController *OrderController) HandleOrder(id int64, status, note string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.StatusCode = strings.ToUpper(status)
		flsoNew := &functions.FOrderLogNew{
			Note: note,
		}
		if retval := flsoNew.Run(order, db); retval.ErrCode != 0 {
			return retval
		}
		cust, notifOrder := getDataNotifOrder(orderController)
		if status == utils.ORDER_STATUS_ACCEPTED {
			if validateAcceptOrder(order, db) {
				SendEmailNotification(APPROVED_NOTIFICATION, cust, *notifOrder)
			} else {
				return utils.StatusReturn{ErrCode: utils.ErrValidate, Message: "Please input shipping cost for order outside Japan!"}
			}
		} else if status == utils.ORDER_STATUS_CANCELED {
			SendEmailNotification(CANCELED_NOTIFICATION, cust, *notifOrder)
		}
		SendStatusOrderPushNotification(status)
		return utils.StatusReturnOK()
	})
}

func (orderController *OrderController) TrackingNumber(id int64, trackingNumber string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.TrackingNumber = strings.ToUpper(trackingNumber)

		orderController.HandleOrder(id, "IS", "")
		return utils.StatusReturnOK()
	})
}

func (orderController *OrderController) ShippingCost(id int64, cost float64) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.ExchangeRate = scrape.GetRupiah()
		order.ShippingCost = cost
		order.GrandTotal = order.Total + cost - order.DiscAmount
		return utils.StatusReturnOK()
	})
}

func (orderController *OrderController) TaxDuty(id int64, taxAmount string) map[string]interface{} {
	return UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.TaxAmount = taxAmount

		return utils.StatusReturnOK()
	})
}

func (orderController *OrderController) ListDetail(param *utils.Param) map[string]interface{} {
	imported := false
	param.Imported = &imported
	param.SetDefaultDelete()
	return ListJoinModel("orderd", "order_id DESC,dno ASC", make([]*model.Orderd, 0), param, func(query *gorm.DB) {
		query.Joins("JOIN \"order\" ON order_id = \"order\".id")
	}, func(query *gorm.DB) {
		query.Where("status_code IN ('PA','IP')")
	})
}

func (orderController *OrderController) UploadOrderProof(w http.ResponseWriter, r *http.Request) {
	mode := os.Getenv("MODE")
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		utils.RespondError(w, utils.MessageErr(false, utils.ErrRequest, err.Error()), http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.RespondError(w, utils.MessageErr(false, utils.ErrRequest, err.Error()), http.StatusBadRequest)
		return
	}

	var fileName string
	if mode != "DEV" {
		file, retval := utils.GetImageFile(w, r)
		if retval.ErrCode != 0 {
			utils.RespondError(w, utils.MessageErr(false, retval.ErrCode, retval.Message), http.StatusUnsupportedMediaType)
			return
		}
		imageWritter := utils.GetImageWritter()
		currentTime := time.Now()

		retval = ViewModel(id, &orderController.Order)
		if retval.ErrCode != 0 {
			utils.Respond(w, utils.MessageErr(false, retval.ErrCode, retval.Message))
			return
		} else {
			proofLink := orderController.Order.ProofLink
			if proofLink != "" {
				if err := imageWritter.DeleteFile(proofLink); err != nil {
					utils.RespondError(w, utils.MessageErr(false, utils.ErrIO, err.Error()), http.StatusNotFound)
					return
				}
			}
		}

		fileName = "payment-" + strconv.Itoa(int(id)) + "-" + currentTime.Format("20060102150405")
		if err := imageWritter.UploadFile(file, fileName); err != nil {
			utils.RespondError(w, utils.MessageErr(false, utils.ErrIO, err.Error()), http.StatusUnsupportedMediaType)
			return
		}
	}
	resp := UpdateFieldMaster(id, orderController, func(m model.Model, db *gorm.DB) utils.StatusReturn {
		order := m.(*model.Order)
		order.ProofLink = fileName

		return utils.StatusReturnOK()
	})
	utils.Respond(w, resp)
}

func (orderController *OrderController) LoadOrderProof(w http.ResponseWriter, r *http.Request) {
	id, err := GetInt64Param("id", w, r)
	if err != nil {
		return
	}
	order := &orderController.Order
	if retval := ViewModel(id, order); retval.ErrCode != 0 {
		utils.Respond(w, utils.MessageErr(false, retval.ErrCode, retval.Message))
		return
	}
	url, retval := utils.GenerateSignedUrl(order.ProofLink)
	if retval.ErrCode != 0 {
		utils.Respond(w, utils.MessageErr(false, retval.ErrCode, retval.Message))
		return
	}
	resp := make(map[string]interface{})
	resp["url"] = url
	utils.Respond(w, resp)
}

type NotifOrder struct {
	Custname string
	Trxno    string
	Trxdate  string
	Addr     string
	Country  string
	Phone    string
	Email    string
	Total    string
	Shipping string
	Grand    string
	Rate     string
	Voucher  string
	Orderd   []DetailNotifOrder
}

func SendEmailNotification(mode int, cust *model.Cust, notifOrder NotifOrder) {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	runtime.GOMAXPROCS(1)
	switch mode {
	case ORDER_NOTIFICATION:
		go utils.SendEmailNotifOrder(adminEmail, cust.Email, notifOrder, notifOrder.Trxno)
	case APPROVED_NOTIFICATION:
		go utils.SendEmailNotifApprove(adminEmail, cust.Email, notifOrder, notifOrder.Trxno)
	case CANCELED_NOTIFICATION:
		go utils.SendEmailNotifCanceled(adminEmail, cust.Email, notifOrder, notifOrder.Trxno)
	}

}

type DetailNotifOrder struct {
	Subtotal string
	model.Orderd
}

func getDataNotifOrder(orderController *OrderController) (*model.Cust, *NotifOrder) {
	db := utils.GetDB()

	order := orderController.Order

	if len(orderController.Orderd) == 0 {
		db.Where("order_id = ?", order.Id).Find(&orderController.Orderd)
		for i := range orderController.Orderd {
			orderController.Details = append(orderController.Details, &orderController.Orderd[i])
		}
	}

	var orderd []DetailNotifOrder
	for _, data := range orderController.Orderd {
		subtotal := humanize.Comma(int64(data.Subtotal))
		detail := &DetailNotifOrder{
			Subtotal: subtotal,
			Orderd:   data,
		}
		orderd = append(orderd, *detail)
	}

	addr := &model.Addr{}
	db.Where("id = ?", order.AddrId).Find(addr)

	cust := &model.Cust{}
	db.Where("id = ?", order.CustId).Find(&cust)

	// voucher :=

	fullAddr := fmt.Sprintf("%s,\n%s, %s, %s - %s", addr.FullAddress, addr.City, addr.Province, addr.CountryCode, addr.ZipCode)
	var grandTotal string
	if order.ExchangeRate > 0 {
		grandTotal = fmt.Sprintf("    (Rp.%s)", humanize.Comma(int64(order.GrandTotal*order.ExchangeRate)))
	}

	notifOrder := &NotifOrder{
		Custname: order.CustName,
		Trxno:    order.Trxno,
		Trxdate:  utils.FormatTimeToDate(order.Date),
		Addr:     fullAddr,
		Country:  addr.CountryCode,
		Phone:    cust.PhoneNumber,
		Email:    cust.Email,
		Total:    humanize.Comma(int64(order.Total)),
		Shipping: humanize.Comma(int64(order.ShippingCost)),
		Grand:    humanize.Comma(int64(order.GrandTotal)) + grandTotal,
		Rate:     humanize.Comma(int64(order.ExchangeRate)),
		Voucher:  GetDataVoucher(&order, db),
		Orderd:   orderd,
	}

	return cust, notifOrder
}

func GetOrderTrxno(order *model.Order, db *gorm.DB) string {
	country := model.GetAddrCountry(order.AddrId, db)
	if strings.ToUpper(country) == "JAPAN" {
		return functions.FGetNewNo("JP", db)
	}
	return functions.FGetNewNo("LN", db)
}

func (orderController *OrderController) ExportOrderXls(orderId int64, w http.ResponseWriter) {
	orderController.ViewTrans(orderId)
	order := orderController.Order
	details := orderController.Orderd

	filename := "order_" + order.Trxno + "_" + utils.FormatTimeToDate(order.Date)

	var header = []string{
		"date",
		"no",
		"product_Id",
		"qty",
		"total",
	}
	data := make([]map[string]interface{}, 0)
	for _, detail := range details {
		mapDetail := make(map[string]interface{})
		mapDetail["date"] = utils.FormatTimeToDate(order.Date)
		mapDetail["no"] = order.Trxno
		mapDetail["product_id"] = detail.ProductId
		mapDetail["qty"] = detail.Qty
		mapDetail["total"] = detail.Subtotal
		data = append(data, mapDetail)
	}
	xls := utils.ConvertToXls(data, header)
	utils.SetXlsHeader(w, filename)

	xls.Write(w)
}

func validateAcceptOrder(order *model.Order, db *gorm.DB) bool {
	country := model.GetAddrCountry(order.AddrId, db)
	if country == "JAPAN" {
		order.ExchangeRate = scrape.GetRupiah()
	}
	return !(strings.ToUpper(country) != "JAPAN" && order.ShippingCost == 0)
}

func SendNewOrderPushNotification(order *model.Order) {
	model.SendPushNotification("New order received", order.Trxno, true)
}

func SendStatusOrderPushNotification(status string) {
	message := GetStatusOrderMessage(status)
	model.SendPushNotification("Easy Shop Order Notification", message, false)
}

func GetStatusOrderMessage(status string) string {
	switch status {
	case utils.ORDER_STATUS_ACCEPTED:
		return "Your order has been accepted. Please continue to proceed your payment!"
	case utils.ORDER_STATUS_PAYMENT_ACCEPTED:
		return "Your payment has been verified. We will process your order!"
	case utils.ORDER_STATUS_PAYMENT_REJECTED:
		return "Your payment has been rejected. Please check again your payment!"
	case utils.ORDER_STATUS_READY:
		return "Your order is ready. We will soon deliver your order!"
	case utils.ORDER_STATUS_SHIPPED:
		return "Your order is on the way to you. Thanks for using our service!"
	}
	return ""
}

func ProcessVoucher(order *model.Order, db *gorm.DB) utils.StatusReturn {
	if order.VoucherId != nil {
		retval, _ := CheckAvailability(order.VoucherId, nil, order.CustId, db)
		if retval.ErrCode == 0 {
			voucherLog := &model.VoucherLog{
				VoucherId: *order.VoucherId,
				CustId:    order.CustId,
				RedeemAt:  time.Now(),
			}
			db.Save(&voucherLog)
		}
		return retval
	}
	return utils.StatusReturnOK()
}

func GetDataVoucher(order *model.Order, db *gorm.DB) string {
	if order.VoucherId != nil {
		voucher := &model.Voucher{}
		db.Where("id = ?", order.VoucherId).Find(&voucher)
		return fmt.Sprintf("%s (%s)", humanize.Comma(int64(order.DiscAmount)), voucher.Code)
	}
	return ""
}
