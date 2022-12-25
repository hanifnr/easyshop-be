package main

import (
	"fmt"
	"net/http"
	"os"

	c "easyshop/controller"
	u "easyshop/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

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
	router.HandleFunc("/cust/email/register", c.RegisterEmail).Methods("POST")
	router.HandleFunc("/cust/email/verifyregister", c.VerifyRegisterEmail).Methods("POST")
	router.HandleFunc("/cust/email/auth", c.AuthEmail).Methods("POST")
	router.HandleFunc("/cust/email/verifyauth", c.VerifyAuthEmail).Methods("POST")
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
	router.HandleFunc("/order/details", c.ListOrderd).Methods("GET")
	router.HandleFunc("/order/log", c.ListOrderLog).Methods("GET")
	router.HandleFunc("/order/proof/create/{id}", c.UploadOrderProof).Methods("POST")
	router.HandleFunc("/order/proof/view/{id}", c.LoadOrderProof).Methods("GET")
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
	router.HandleFunc("/status", c.ListStatus).Methods("GET")
	router.HandleFunc("/status/view/{id}", c.ViewStatus).Methods("GET")

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
