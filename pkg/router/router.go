package api

import "net/http"

type Router interface {
	GET(uri string, funcojb func(responce http.ResponseWriter, request *http.Request))
	POST(uri string, funcojb func(responce http.ResponseWriter, request *http.Request))
	UPDATE(uri string, funcojb func(respoce http.ResponseWriter, request *http.Request))
	DELETE(uri string, funcojb func(respoce http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
