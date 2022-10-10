package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	c "easyshop/controller"
	u "easyshop/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func main() {
	port := os.Getenv("APP_PORT")
	u.SetAuthSecret("1GN1T3CH")
	u.SetNoAuth([]string{"/basictoken"})

	router := mux.NewRouter()
	router.HandleFunc("/", handlerIndex)
	router.HandleFunc("/index", handlerIndex)
	router.HandleFunc("/basictoken", c.BasicTokenController).Methods("GET")

	router.Use(u.JwtAuthentication)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("listening to port %s", port)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcomee"
	w.Write([]byte(message))
}
