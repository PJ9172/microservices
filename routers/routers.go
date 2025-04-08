package routers

import (
	"Microservices/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "templates/index.html") })
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	r.HandleFunc("/products", handlers.ShowProducts).Methods("GET")
	r.HandleFunc("/purchase", handlers.PurchaseProduct).Methods("POST")
	r.HandleFunc("/payment", handlers.MakePayment).Methods("GET")

	return r
}
