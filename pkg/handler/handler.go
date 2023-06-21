package handler

import (
	"BillingGo/pkg/model"
	"BillingGo/pkg/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewUser model.User

const (
	CONTENTTYPE     = "ContentType"
	APPLICATIONJSON = "application/json"
)

func Getuser(w http.ResponseWriter, r *http.Request) {
	NewUser := model.GetAllUsers()
	res, _ := json.Marshal(NewUser)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func Getuserbyid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Print("ERROR WHILE PARSING ")
	}
	UserDetails, _ := model.Getuserbyid(ID)
	res, _ := json.Marshal(UserDetails)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Createuser(w http.ResponseWriter, r *http.Request) {
	Createuser := &model.User{}
	utils.ParseBody(r, Createuser)
	b := Createuser.Createuser()
	res, _ := json.Marshal(b)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Print("ERROR WHILE PARSING ")
	}
	user := model.Deleteuser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set(CONTENTTYPE, APPLICATIONJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func Updateuser(w http.ResponseWriter, r *http.Request) {
	var updateuser = &model.User{}
	//utils.parseBody(r, updateuser)
	utils.ParseBody(r, updateuser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Print("ERROR WHILR PARSING ")
	}
	userDetails, db := model.Getuserbyid(ID)
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
