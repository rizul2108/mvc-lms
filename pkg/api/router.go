package api

import (
	"mvc-go/pkg/controller"
	// "mvc-go/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/home", controller.Home).Methods("GET")
	r.HandleFunc("/signup", controller.SignUp).Methods("GET")
	r.HandleFunc("/signup", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/login", controller.LogIn).Methods("GET")
	r.HandleFunc("/profile", controller.Profile).Methods("GET")
	r.HandleFunc("/adminBooks", controller.AdminBooks).Methods("GET")
	r.HandleFunc("/makeAdmin", controller.MakeAdmin).Methods("GET")
	r.HandleFunc("/admin/add_book", controller.AddNewBook).Methods("POST")
	r.HandleFunc("/client/books", controller.ClientBooks).Methods("GET")
	r.HandleFunc("/admin/addBook", controller.AddBook).Methods("GET")
	r.HandleFunc("/makeAdmin", controller.AddAdmin).Methods("POST")
	r.HandleFunc("/addQty", controller.AddQuantity).Methods("POST")
	r.HandleFunc("/delete", controller.DecreaseQuantity).Methods("POST")
	r.HandleFunc("/issueBook", controller.AddRequest).Methods("POST")
	http.ListenAndServe(":9000", r)
}
