// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/Dinesh/go/src/BillingGo/CustomerMS/repository/bill-repo.go
//
// Generated by this command:
//
//	mockgen -source=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/repository/bill-repo.go -destination=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/mocks/repositoryMocks.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	models "CustomerMS/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBillRespository is a mock of BillRespository interface.
type MockBillRespository struct {
	ctrl     *gomock.Controller
	recorder *MockBillRespositoryMockRecorder
}

// MockBillRespositoryMockRecorder is the mock recorder for MockBillRespository.
type MockBillRespositoryMockRecorder struct {
	mock *MockBillRespository
}

// NewMockBillRespository creates a new mock instance.
func NewMockBillRespository(ctrl *gomock.Controller) *MockBillRespository {
	mock := &MockBillRespository{ctrl: ctrl}
	mock.recorder = &MockBillRespositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBillRespository) EXPECT() *MockBillRespositoryMockRecorder {
	return m.recorder
}

// CreateCutomer mocks base method.
func (m *MockBillRespository) CreateCutomer(model *models.Customer) (*models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCutomer", model)
	ret0, _ := ret[0].(*models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCutomer indicates an expected call of CreateCutomer.
func (mr *MockBillRespositoryMockRecorder) CreateCutomer(model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCutomer", reflect.TypeOf((*MockBillRespository)(nil).CreateCutomer), model)
}

// DeleteCutomer mocks base method.
func (m *MockBillRespository) DeleteCutomer(id string) (*models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCutomer", id)
	ret0, _ := ret[0].(*models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCutomer indicates an expected call of DeleteCutomer.
func (mr *MockBillRespositoryMockRecorder) DeleteCutomer(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCutomer", reflect.TypeOf((*MockBillRespository)(nil).DeleteCutomer), id)
}

// GetAllCutomer mocks base method.
func (m *MockBillRespository) GetAllCutomer() ([]*models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCutomer")
	ret0, _ := ret[0].([]*models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCutomer indicates an expected call of GetAllCutomer.
func (mr *MockBillRespositoryMockRecorder) GetAllCutomer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCutomer", reflect.TypeOf((*MockBillRespository)(nil).GetAllCutomer))
}

// GetCutomerById mocks base method.
func (m *MockBillRespository) GetCutomerById(id string) (models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCutomerById", id)
	ret0, _ := ret[0].(models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCutomerById indicates an expected call of GetCutomerById.
func (mr *MockBillRespositoryMockRecorder) GetCutomerById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCutomerById", reflect.TypeOf((*MockBillRespository)(nil).GetCutomerById), id)
}

// UpdateCutomer mocks base method.
func (m *MockBillRespository) UpdateCutomer(model *models.Customer) (*models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCutomer", model)
	ret0, _ := ret[0].(*models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCutomer indicates an expected call of UpdateCutomer.
func (mr *MockBillRespositoryMockRecorder) UpdateCutomer(model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCutomer", reflect.TypeOf((*MockBillRespository)(nil).UpdateCutomer), model)
}
