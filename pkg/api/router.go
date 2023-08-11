package api

import (
	"mvc-go/pkg/controller"
	"mvc-go/pkg/models"
	"mvc-go/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()

	r.Use(models.TokenMiddleware)
	//common routes
	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/home", controller.Home).Methods("GET")
	r.HandleFunc("/signup", controller.SignUp).Methods("GET")
	r.HandleFunc("/signup", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/login", controller.LogIn).Methods("GET")

	//client Routes
	r.HandleFunc("/client/profile", controller.Profile).Methods("GET")
	r.HandleFunc("/client/books", controller.ClientBooks).Methods("GET")
	r.HandleFunc("/client/issueBook", controller.AddRequest).Methods("POST")
	r.HandleFunc("/client/deleteRequest", controller.DeleteRequest).Methods("POST")
	r.HandleFunc("/client/returnBook", controller.ReturnBook).Methods("POST")
	r.HandleFunc("/client/serverError", controller.ClientServerError).Methods("GET")

	//admin Routes
	r.HandleFunc("/admin/addQuantity", controller.AddQuantity).Methods("POST")
	r.HandleFunc("/admin/deleteBooks", controller.DecreaseQuantity).Methods("POST")
	r.HandleFunc("/admin/acceptRequest", controller.AcceptRequest).Methods("POST")
	r.HandleFunc("/admin/declineRequest", controller.DeclineRequest).Methods("POST")
	r.HandleFunc("/admin/addBook", controller.AddBook).Methods("GET")
	r.HandleFunc("/admin/requests", controller.AdminRequests).Methods("GET")
	r.HandleFunc("/admin/adminRequests", controller.GetAdminRequests).Methods("GET")
	r.HandleFunc("/admin/makeAdmin", controller.AddAdmin).Methods("POST")
	r.HandleFunc("/admin/addBook", controller.AddNewBook).Methods("POST")
	r.HandleFunc("/admin/deleteBook", controller.DeleteBook).Methods("POST")
	r.HandleFunc("/admin/books", controller.AdminBooks).Methods("GET")
	r.HandleFunc("/admin/makeAdmin", controller.MakeAdmin).Methods("GET")
	r.HandleFunc("/admin/declineAdmin", controller.DeclineAdmin).Methods("POST")
	r.HandleFunc("/admin/acceptAdmin", controller.AcceptAdmin).Methods("POST")
	r.HandleFunc("/admin/serverError", controller.AdminServerError).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)

	r.HandleFunc("/logout", controller.Logout).Methods("GET")
	pathCSS, err := utils.GetCurrentDirPath()
	if err == nil {
		s := http.StripPrefix("/static/", http.FileServer(http.Dir(pathCSS+"/templates/static/")))
		r.PathPrefix("/static/").Handler(s)
	}

	http.ListenAndServe(":9000", r)
}
