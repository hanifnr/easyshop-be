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
	})

	router := mux.NewRouter()
	router.HandleFunc("/", handlerIndex)
	router.HandleFunc("/index", handlerIndex)
	router.HandleFunc("/basictoken", c.BasicTokenController).Methods("GET")
	router.HandleFunc("/cust", c.ListCust).Methods("GET")
	router.HandleFunc("/cust/{id}", c.ViewCust).Methods("GET")
	router.HandleFunc("/cust/create", c.CreateCust).Methods("POST")
	router.HandleFunc("/cust/update", c.UpdateCust).Methods("POST")
	router.HandleFunc("/cust/handle", c.HandleCust).Methods("POST")
	router.HandleFunc("/shop", c.ListShop).Methods("GET")
	router.HandleFunc("/shop/{id}", c.ViewShop).Methods("GET")
	router.HandleFunc("/shopcategory", c.ListShopCategory).Methods("GET")
	router.HandleFunc("/shopcategory/{id}", c.ViewShopCategory).Methods("GET")
	router.HandleFunc("/order", c.ListOrder).Methods("GET")
	router.HandleFunc("/order/{id}", c.ViewOrder).Methods("GET")
	router.HandleFunc("/order/create", c.CreateOrder).Methods("POST")
	router.HandleFunc("/order/update", c.UpdateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", c.ViewOrder).Methods("GET")
	router.HandleFunc("/order/handle", c.HandleOrder).Methods("POST")
	router.HandleFunc("/passport", c.ListPassport).Methods("GET")
	router.HandleFunc("/passport/{id}", c.ViewPassport).Methods("GET")
	router.HandleFunc("/passport/create", c.CreatePassport).Methods("POST")
	router.HandleFunc("/passport/update", c.UpdatePassport).Methods("POST")
	router.HandleFunc("/passport/cust/{id}", c.ViewPassportCust).Methods("GET")

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
