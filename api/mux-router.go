package api

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
	port          = os.Getenv("SERVERPORT")
	srv           = &http.Server{
		Addr:    port,
		Handler: muxDispatcher,
	}
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

// DELETE implements Router.
func (*muxRouter) DELETE(uri string, funcojb func(respoce http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funcojb).Methods("DELETE")
	//panic("unimplemented")
}

// UPDATE implements Router.
func (*muxRouter) UPDATE(uri string, funcojb func(respoce http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funcojb).Methods("UPDATE")
	//panic("unimplemented")
}

// GET implements Router.
func (*muxRouter) GET(uri string, funcojb func(responce http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funcojb).Methods("GET")
	//panic("unimplemented")
}

// POST implements Router.
func (*muxRouter) POST(uri string, funcojb func(responce http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, funcojb).Methods("POST")
	//panic("unimplemented")
}

// SERVE implements Router.
func (*muxRouter) SERVE(port string) {
	logrus.Infof("HTTP Server is Running on PORT %v", port)
	//http.ListenAndServe(port, muxDispatcher)
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("Http Server Failed to start: %v", err)
	}
}

// To handle the GreaceFul Shutdown after interruptions
func (*muxRouter) GraceFulShutDown(ctx context.Context) error {
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Printf("http Server shutdown error : %v", err)
	}
	return nil
}
