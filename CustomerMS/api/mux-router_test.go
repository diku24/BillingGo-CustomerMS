package api

import (
	"BillingGo/mocks"
	"context"
	"net/http"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func generateMockHandler(t *testing.T) *mocks.MockBillHandler {
	control := gomock.NewController(t)
	defer control.Finish()
	mockHandler := mocks.NewMockBillHandler(control)

	return mockHandler

}
func TestNewMuxRouter(t *testing.T) {

	tests := []struct {
		name string
		want Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMuxRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMuxRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_muxRouter_DELETE(t *testing.T) {
	type args struct {
		uri     string
		funcojb func(respoce http.ResponseWriter, request *http.Request)
	}
	tests := []struct {
		name string
		m    *muxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.DELETE(tt.args.uri, tt.args.funcojb)
		})
	}
}

func Test_muxRouter_UPDATE(t *testing.T) {
	type args struct {
		uri     string
		funcojb func(respoce http.ResponseWriter, request *http.Request)
	}
	tests := []struct {
		name string
		m    *muxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.UPDATE(tt.args.uri, tt.args.funcojb)
		})
	}
}

func Test_muxRouter_GET(t *testing.T) {
	type args struct {
		uri     string
		funcojb func(responce http.ResponseWriter, request *http.Request)
	}
	tests := []struct {
		name string
		m    *muxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.GET(tt.args.uri, tt.args.funcojb)
		})
	}
}

func Test_muxRouter_POST(t *testing.T) {
	type args struct {
		uri     string
		funcojb func(responce http.ResponseWriter, request *http.Request)
	}
	tests := []struct {
		name string
		m    *muxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.POST(tt.args.uri, tt.args.funcojb)
		})
	}
}

func Test_muxRouter_SERVE(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name string
		m    *muxRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SERVE(tt.args.port)
		})
	}
}

func Test_muxRouter_GraceFulShutDown(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		m       *muxRouter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.GraceFulShutDown(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("muxRouter.GraceFulShutDown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMakeHTTPHandlerFunction(t *testing.T) {
	type args struct {
		funcy apiFunction
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeHTTPHandlerFunction(tt.args.funcy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeHTTPHandlerFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
