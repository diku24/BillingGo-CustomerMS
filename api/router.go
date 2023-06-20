package api

import "net/http"

type Router interface {
	GET(uri string, funcojb func(responce http.ResponseWriter, request *http.Request))
	POST(uri string, funcojb func(responce http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
