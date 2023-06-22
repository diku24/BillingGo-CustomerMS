package handler

import "net/http"

type Handler interface {
	Create(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Getbyid(response http.ResponseWriter, request *http.Request)
}
