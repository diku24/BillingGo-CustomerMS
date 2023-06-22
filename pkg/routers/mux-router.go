package routers

import (
	"BillingGo/pkg/handler"

	"github.com/gorilla/mux"
)

type muxRouter struct {
}

const (
	users        = "/users/"
	usersWithId  = "/users/{UserId}"
	userWithName = "/users/{Username}"
)

var RegisterUser = func(router *mux.Router) {
	router.HandleFunc(users, handler.NewUserHandler().Create).Methods("POST")
	router.HandleFunc(users, handler.NewUserHandler().Get).Methods("GET")
	router.HandleFunc(usersWithId, handler.NewUserHandler().Getbyid).Methods("GET")
	//router.HandleFunc(userWithName, handler.Getuser).Methods("GET")

	router.HandleFunc(usersWithId, handler.NewUserHandler().Update).Methods("PUT")
	//router.HandleFunc(userWithName, handler.Updateuser).Methods("PUT")

	router.HandleFunc(usersWithId, handler.NewUserHandler().Delete).Methods("DELETE")
	//router.HandleFunc(userWithName, handler.Deleteuser).Methods("DELETE")
}

// var RegisterCoustomer = func(router *mux.router) {
// 	router.HandleFunc("/customer/", handler.Createclient).Methods("POST")
// 	router.HandleFunc("/customer/", handler.getclient).Methods("GET")
// 	router.HandleFunc("/customer/{CutomerId}", handler.getclientbyid).Methods("GET")
// 	router.HandleFunc("/customer/{CustomerName}", handler.getclientbyname).Methods("GET")

// 	router.HandleFunc("/customer/{CutomerId}", handler.Updateclientbyid).Methods("PUT")
// 	router.HandleFunc("/customer/{CustomerName}", handler.Updateclientbyname).Methods("PUT")

// 	router.HandleFunc("/customer/{CutomerId}", handler.Deleteclientbyid).Methods("DELETE")
// 	router.HandleFunc("/customer/{CustomerName}", handler.Deleteclientbyname).Methods("DELETE")
// }

// var Registerproduc = func(router *mux.router) {
// 	router.HandleFunc("/product/", handler.Createuser).Methods("POST")
// 	router.HandleFunc("/product/", handler.getuser).Methods("GET")
// 	router.HandleFunc("/product/{pId}", handler.getuserbyid).Methods("GET")
// 	router.HandleFunc("/product/{pname}", handler.getuserbyid).Methods("GET")

// 	router.HandleFunc("/product/{pId}", handler.Updateuserbyid).Methods("PUT")
// 	router.HandleFunc("/product/{pname}", handler.Updateuserbyname).Methods("PUT")

// 	router.HandleFunc("/product/{pId}", handler.Deleteuserbyid).Methods("DELETE")
// 	router.HandleFunc("/product/{pname}", handler.Deleteuserbyname).Methods("DELETE")
// }

// var Registerperches = func(router *mux.router) {
// 	router.HandleFunc("/purchases/", handler.Createuser).Methods("POST")
// 	router.HandleFunc("/purchases/", handler.getuser).Methods("GET")
// 	router.HandleFunc("/purchases/{rate}", handler.getuserbyid).Methods("GET")
// 	router.HandleFunc("/purchases/{quantity}", handler.getuserbyid).Methods("GET")

// 	router.HandleFunc("/purchases/{rate}", handler.Updateuserbyid).Methods("PUT")
// 	router.HandleFunc("/purchases/{quantity}", handler.Updateuserbyname).Methods("PUT")

// 	router.HandleFunc("/purchases/{rate}", handler.Deleteuserbyid).Methods("DELETE")
// 	router.HandleFunc("/purchases/{quantity}", handler.Deleteuserbyname).Methods("DELETE")
// }

// var Registersals = func(router *mux.router) {
// 	router.HandleFunc("/sales/", handler.Createuser).Methods("POST")
// 	router.HandleFunc("/sales/", handler.getuser).Methods("GET")
// 	router.HandleFunc("/sales/{rate}", handler.getuserbyid).Methods("GET")
// 	router.HandleFunc("/sales/{quantity}", handler.getuserbyid).Methods("GET")

// 	router.HandleFunc("/sales/{rate}", handler.Updateuserbyid).Methods("PUT")
// 	router.HandleFunc("/sales/{quantity}", handler.Updateuserbyname).Methods("PUT")

// 	router.HandleFunc("/sales/{rate}", handler.Deleteuserbyid).Methods("DELETE")
// 	router.HandleFunc("/sales/{quantity}", handler.Deleteuserbyname).Methods("DELETE")
// }
