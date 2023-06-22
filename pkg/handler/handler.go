package handler

import (
	"BillingGo/pkg/model"
	"BillingGo/pkg/utils"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var NewUser model.User

const (
	CONTENTTYPE     = "ContentType"
	APPLICATIONJSON = "application/json"
)

type UserHandler struct{}

func NewUserHandler() Handler {
	return &UserHandler{}
}

func (*UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	Createuser := &model.User{}
	utils.ParseBody(r, Createuser)
	b := Createuser.Createuser()
	res, _ := json.Marshal(b)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func (*UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	NewUser := model.GetAllUsers()
	res, _ := json.Marshal(NewUser)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (*UserHandler) Getbyid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["UserId"]
	UserDetails, _, _ := model.Getbyid(userId)
	res, _ := json.Marshal(UserDetails)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (*UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["UserId"]
	user := model.Deleteuser(userId)
	res, _ := json.Marshal(user)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func (*UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updateuser = &model.User{}
	utils.ParseBody(r, updateuser)
	vars := mux.Vars(r)
	userId := vars["UserId"]
	userDetails, db, _ := model.Getbyid(userId)
	if updateuser.Username != "" {
		userDetails.Username = updateuser.Username
	}
	if updateuser.Userid != "" {
		userDetails.Userid = updateuser.Userid
	}
	if updateuser.Password != "" {
		userDetails.Password = updateuser.Password
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
