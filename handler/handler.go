package handler

import "net/http"

type BillHandler interface {
	GET(response http.ResponseWriter, req *http.Request)
	POST(response http.ResponseWriter, req *http.Request)
	PUT(response http.ResponseWriter, req *http.Request)
	DELETE(response http.ResponseWriter, req *http.Request)
}
