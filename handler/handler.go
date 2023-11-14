package handler

import "net/http"

type BillHandler interface {
	GET(response http.ResponseWriter, req *http.Request) error
	POST(response http.ResponseWriter, req *http.Request) error
	PUT(response http.ResponseWriter, req *http.Request) error
	DELETE(response http.ResponseWriter, req *http.Request) error
}
