package main

import (
	"BillingGo/pkg/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterUser(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8800", r))
}
