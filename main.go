package main

import (
	"fmt"
	"net/http"
	"os"

	c "easyshop/controller"
	"easyshop/service"
	u "easyshop/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	service.InitFirebase()
	// service.InitRedis()
}

func main() {
	u.SetAuthSecret("1GN1T3CH")
	u.SetNoAuth([]string{
		"/basictoken",
		"/cust",
		"/shop",
		"/order",
		"/passport",
		"/addr",
		"/purc",
		"/wh",
		"/status",
		"/status",
		"/product",
		"/properties",
		"/firebase",
		"/notification",
		"/partnership",
		"/voucher",
		"/email",
		"/req",
	})
	service.InitScheduler([]func(){
		c.CleanProduct,
		c.CleanEmail,
	})
	router := mux.NewRouter()
	router.HandleFunc("/", handlerIndex)
	router.HandleFunc("/index", handlerIndex)
	router.HandleFunc("/basictoken", c.BasicTokenController).Methods("GET")
	router.HandleFunc("/cust", c.ListCust).Methods("GET")
	router.HandleFunc("/cust/view/{id}", c.ViewCust).Methods("GET")
	router.HandleFunc("/cust/create", c.CreateCust).Methods("POST")
	router.HandleFunc("/cust/update", c.UpdateCust).Methods("PUT")
	router.HandleFunc("/cust/delete/{id}", c.DeleteCust).Methods("DELETE")
	router.HandleFunc("/cust/combo", c.ListComboCust).Methods("GET")
	router.HandleFunc("/email/register", c.RegisterEmail).Methods("POST")
	router.HandleFunc("/email/verifyregister", c.VerifyRegisterEmail).Methods("POST")
	router.HandleFunc("/email/auth", c.AuthEmail).Methods("POST")
	router.HandleFunc("/email/verifyauth", c.VerifyAuthEmail).Methods("POST")
	router.HandleFunc("/shop", c.ListShop).Methods("GET")
	router.HandleFunc("/shop/view/{id}", c.ViewShop).Methods("GET")
	router.HandleFunc("/shop/combo", c.ListComboShop).Methods("GET")
	router.HandleFunc("/shopcategory", c.ListShopCategory).Methods("GET")
	router.HandleFunc("/shopcategory/view/{id}", c.ViewShopCategory).Methods("GET")
	router.HandleFunc("/order", c.ListOrder).Methods("GET")
	router.HandleFunc("/order/view/{id}", c.ViewOrder).Methods("GET")
	router.HandleFunc("/order/delete/{id}", c.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/order/create", c.CreateOrder).Methods("POST")
	router.HandleFunc("/order/update", c.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/handle", c.HandleOrder).Methods("PUT")
	router.HandleFunc("/order/track", c.TrackingNumber).Methods("PUT")
	router.HandleFunc("/order/shipping", c.ShippingCost).Methods("PUT")
	router.HandleFunc("/order/taxduty", c.TaxDuty).Methods("PUT")
	router.HandleFunc("/order/details", c.ListOrderd).Methods("GET")
	router.HandleFunc("/order/log", c.ListOrderLog).Methods("GET")
	router.HandleFunc("/order/proof/create/{id}", c.UploadOrderProof).Methods("POST")
	router.HandleFunc("/order/proof/view/{id}", c.LoadOrderProof).Methods("GET")
	router.HandleFunc("/order/tax/{id}", c.GetTaxOffice).Methods("GET")
	router.HandleFunc("/passport", c.ListPassport).Methods("GET")
	router.HandleFunc("/passport/view/{id}", c.ViewPassport).Methods("GET")
	router.HandleFunc("/passport/create", c.CreatePassport).Methods("POST")
	router.HandleFunc("/passport/update", c.UpdatePassport).Methods("PUT")
	router.HandleFunc("/passport/cust/{id}", c.ViewPassportCust).Methods("GET")
	router.HandleFunc("/addr", c.ListAddr).Methods("GET")
	router.HandleFunc("/addr/view/{id}", c.ViewAddr).Methods("GET")
	router.HandleFunc("/addr/create", c.CreateAddr).Methods("POST")
	router.HandleFunc("/addr/update", c.UpdateAddr).Methods("PUT")
	router.HandleFunc("/addr/combo", c.ListComboAddr).Methods("GET")
	router.HandleFunc("/addr/delete/{id}", c.DeleteAddr).Methods("DELETE")
	router.HandleFunc("/purc", c.ListPurc).Methods("GET")
	router.HandleFunc("/purc/view/{id}", c.ViewPurc).Methods("GET")
	router.HandleFunc("/purc/create", c.CreatePurc).Methods("POST")
	router.HandleFunc("/purc/update", c.UpdatePurc).Methods("PUT")
	router.HandleFunc("/purc/delete/{id}", c.DeletePurc).Methods("DELETE")
	router.HandleFunc("/purc/details", c.ListPurcd).Methods("GET")
	router.HandleFunc("/purc/shop", c.ListPurcShop).Methods("GET")
	router.HandleFunc("/wh", c.ListWh).Methods("GET")
	router.HandleFunc("/wh/view/{id}", c.ViewWh).Methods("GET")
	router.HandleFunc("/wh/delete/{id}", c.DeleteWh).Methods("DELETE")
	router.HandleFunc("/wh/create", c.CreateWh).Methods("POST")
	router.HandleFunc("/wh/update", c.UpdateWh).Methods("PUT")
	router.HandleFunc("/wh/handle", c.HandleWh).Methods("PUT")
	router.HandleFunc("/status", c.ListStatus).Methods("GET")
	router.HandleFunc("/status/view/{id}", c.ViewStatus).Methods("GET")
	router.HandleFunc("/product/list", c.ListProduct).Methods("GET")
	router.HandleFunc("/product", c.ViewProduct).Methods("GET")
	router.HandleFunc("/product/top", c.GetTopProduct).Methods("GET")
	router.HandleFunc("/product/easyshop", c.GetEasyShopProduct).Methods("GET")
	router.HandleFunc("/product/trending", c.ListTrendingProduct).Methods("GET")
	router.HandleFunc("/product/trending/view/{id}", c.ViewTrendingProduct).Methods("GET")
	router.HandleFunc("/product/trending/create", c.CreateTrendingProduct).Methods("POST")
	router.HandleFunc("/product/trending/update", c.UpdateTrendingProduct).Methods("PUT")
	router.HandleFunc("/product/trending/delete/{id}", c.DeleteTrendingProduct).Methods("DELETE")
	// router.HandleFunc("/product/trending/clean", c.CleanProduct).Methods("DELETE")
	router.HandleFunc("/firebase/token", c.ListFirebaseToken).Methods("GET")
	router.HandleFunc("/firebase/token/view/{id}", c.ViewFirebaseToken).Methods("GET")
	router.HandleFunc("/firebase/token/create", c.CreateFirebaseToken).Methods("POST")
	router.HandleFunc("/firebase/token/update", c.UpdateFirebaseToken).Methods("PUT")
	router.HandleFunc("/partnership", c.ListPartnership).Methods("GET")
	router.HandleFunc("/partnership/view/{id}", c.ViewPartnership).Methods("GET")
	router.HandleFunc("/partnership/create", c.CreatePartnership).Methods("POST")
	router.HandleFunc("/partnership/update", c.UpdatePartnership).Methods("PUT")
	router.HandleFunc("/partnership/delete/{id}", c.DeletePartnership).Methods("DELETE")
	router.HandleFunc("/partnership/combo", c.ListComboPartnership).Methods("GET")
	router.HandleFunc("/partnership/type/combo", c.ListComboPartnershipType).Methods("GET")
	router.HandleFunc("/partnership/approve", c.ApprovePartnership).Methods("PUT")
	router.HandleFunc("/voucher", c.ListVoucher).Methods("GET")
	router.HandleFunc("/voucher/view/{id}", c.ViewVoucher).Methods("GET")
	router.HandleFunc("/voucher/create", c.CreateVoucher).Methods("POST")
	router.HandleFunc("/voucher/update", c.UpdateVoucher).Methods("PUT")
	router.HandleFunc("/voucher/delete/{id}", c.DeleteVoucher).Methods("DELETE")
	router.HandleFunc("/voucher/check", c.CheckVoucher).Methods("GET")
	router.HandleFunc("/voucher/log", c.ListVoucherLog).Methods("GET")
	router.HandleFunc("/req", c.ListReqOrder).Methods("GET")
	router.HandleFunc("/req/view/{id}", c.ViewReqOrder).Methods("GET")
	router.HandleFunc("/req/create", c.CreateReqOrder).Methods("POST")
	router.HandleFunc("/req/update", c.UpdateReqOrder).Methods("PUT")
	router.HandleFunc("/req/delete/{id}", c.DeleteReqOrder).Methods("DELETE")
	router.HandleFunc("/req/handle", c.HandleReqOrder).Methods("PUT")
	router.HandleFunc("/req/approve", c.ApproveReqOrder).Methods("PUT")
	router.HandleFunc("/req/count", c.CountWaitingReqOrder).Methods("GET")
	router.HandleFunc("/notification", c.ListNotification).Methods("GET")
	router.HandleFunc("/properties", c.GetProps).Methods("GET")

	router.Use(u.JwtAuthentication)

	port := os.Getenv("APP_PORT")

	optionsCode := handlers.OptionStatusCode(204)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-CSRF-Token", "Content-Length", "Accept-Encoding", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT", "OPTIONS"})

	fmt.Printf("listening to port %s", port)

	err := http.ListenAndServe(":"+port, handlers.CORS(optionsCode, originsOk, headersOk, methodsOk)(router))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcomee"
	w.Write([]byte(message))
}
