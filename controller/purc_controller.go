package controllers

import (
	"easyshop/functions"
	"easyshop/model"
	"easyshop/utils"
	"net/http"

	"gorm.io/gorm"
)

var CreatePurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	CreateTransAction(purcController, w, r)
}

var UpdatePurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	UpdateTransAction(purcController, w, r)
}

var DeletePurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	DeleteTransAction(purcController, w, r)
}

var ViewPurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	ViewTransAction(purcController, w, r)
}

var ListPurc = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	ListTransAction(purcController, w, r)
}

var ListPurcd = func(w http.ResponseWriter, r *http.Request) {
	param := utils.ProcessParam(r)
	purcController := &PurcController{}
	resp := purcController.ListDetail(param)
	utils.Respond(w, resp)
}

var ListPurcShop = func(w http.ResponseWriter, r *http.Request) {
	purcController := &PurcController{}
	resp := purcController.ListShop()
	utils.Respond(w, resp)
}

type PurcController struct {
	Purc    model.Purc    `json:"purc"`
	Purcd   []model.Purcd `json:"purcd"`
	Details []model.Model `json:"-"`
}

func (purcController *PurcController) MasterField() string {
	return "purc_id"
}

func (purcController *PurcController) MasterModel() model.Model {
	return &purcController.Purc
}

func (purcController *PurcController) DetailsModel() []model.Model {
	return purcController.Details
}

func (purcController *PurcController) CreateTrans() map[string]interface{} {
	if retval := CreateTrans(purcController, func(db *gorm.DB) error {
		for i := range purcController.Purcd {
			purcd := &purcController.Purcd[i]
			type OrderProduct struct {
				ProductId string
				Name      string
			}
			orderProduct := &OrderProduct{}
			if err := db.Select("product_id,name").Table("orderd").Where("order_id=? AND dno=?", purcd.OrderId, purcd.Dno).Scan(&orderProduct).Error; err != nil {
				return err
			}
			purcd.ProductId = orderProduct.ProductId
			purcd.Name = orderProduct.Name
			purcController.Details = append(purcController.Details, &purcController.Purcd[i])
		}
		return nil
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, purcController)
}

func (purcController *PurcController) ViewTrans(id int64) map[string]interface{} {
	if retval := ViewTrans(id, purcController, func(db *gorm.DB) error {
		err := db.Where("purc_id = ?", id).Find(&purcController.Purcd).Error
		for i := range purcController.Purcd {
			purcController.Details = append(purcController.Details, &purcController.Purcd[i])
		}
		return err
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, purcController)
}

func (purcController *PurcController) ListTrans(param *utils.Param) map[string]interface{} {
	return ListTrans("purc", "id ASC", &purcController.Purc, make([]*model.Purc, 0), param)
}

func (purcController *PurcController) UpdateTrans() map[string]interface{} {
	retval := UpdateTrans(purcController, &model.Purc{}, &model.Purcd{}, func(modelSrc, modelTemp model.Model, db *gorm.DB) error {
		purcSrc := modelSrc.(*model.Purc)
		purcTemp := modelTemp.(*model.Purc)
		purcSrc.Trxno = purcTemp.Trxno
		purcSrc.Date = purcTemp.Date
		purcSrc.ShopId = purcTemp.ShopId
		purcSrc.Total = purcTemp.Total
		purcSrc.Refno = purcTemp.Refno
		for i := range purcController.Purcd {
			purcd := &purcController.Purcd[i]
			type OrderProduct struct {
				ProductId string
				Name      string
			}
			orderProduct := &OrderProduct{}
			if err := db.Select("product_id,name").Table("orderd").Where("order_id=? AND dno=?", purcd.OrderId, purcd.Dno).Scan(&orderProduct).Error; err != nil {
				return err
			}
			purcd.ProductId = orderProduct.ProductId
			purcd.Name = orderProduct.Name
			purcController.Details = append(purcController.Details, purcd)
		}
		return nil
	})
	if retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.MessageData(true, purcController)

}

func (purcController *PurcController) DeleteTrans(id int64) map[string]interface{} {
	if retval := DeleteTrans(id, purcController, func() utils.StatusReturn {
		return utils.StatusReturnOK()
	}); retval.ErrCode != 0 {
		return utils.MessageErr(false, retval.ErrCode, retval.Message)
	}
	return utils.Message(true)
}

func (purcController *PurcController) FNew() functions.SQLFunction {
	return &functions.FPurcNew{}
}

func (purcController *PurcController) FDelete() functions.SQLFunction {
	return &functions.FPurcDelete{}
}

func (purcController *PurcController) ListDetail(param *utils.Param) map[string]interface{} {
	imported := false
	param.Imported = &imported
	param.SetDefaultDelete()
	return ListJoinModel("purcd", "purc_id DESC,dno ASC", make([]*model.Purcd, 0), param, func(query *gorm.DB) {
		query.Joins("JOIN purc ON purc_id = purc.id")
	}, func(query *gorm.DB) {})
}

func (purcController *PurcController) ListShop() map[string]interface{} {
	db := utils.GetDB()

	type ShopOrder struct {
		Id   int64   `json:"id"`
		Name string  `json:"name"`
		Qty  float64 `json:"qty"`
	}

	list := make([]*ShopOrder, 0)
	db.Select("shop.id, shop.name, SUM(orderd.qty-orderd.qtypurc) AS qty").
		Table("orderd").
		Joins("JOIN \"order\" ON \"order\".id = orderd.order_id").
		Joins("JOIN shop ON shop.id = orderd.shop_id").
		Where("status_code IN ('PA','IP') AND imported = FALSE").
		Group("shop.id, shop.name, shop.idx").Order("shop.idx").
		Scan(&list)

	return utils.MessageData(true, list)
}
