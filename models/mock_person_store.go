// Code generated by MockGen. DO NOT EDIT.
// Source: person_store.go

// Package models is a generated GoMock package.
package models

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockPersonStoreInterface is a mock of PersonStoreInterface interface.
type MockPersonStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersonStoreInterfaceMockRecorder
}

// MockPersonStoreInterfaceMockRecorder is the mock recorder for MockPersonStoreInterface.
type MockPersonStoreInterfaceMockRecorder struct {
	mock *MockPersonStoreInterface
}

// NewMockPersonStoreInterface creates a new mock instance.
func NewMockPersonStoreInterface(ctrl *gomock.Controller) *MockPersonStoreInterface {
	mock := &MockPersonStoreInterface{ctrl: ctrl}
	mock.recorder = &MockPersonStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonStoreInterface) EXPECT() *MockPersonStoreInterfaceMockRecorder {
	return m.recorder
}

// FindMany mocks base method.
func (m *MockPersonStoreInterface) FindMany(query bson.M, options options.FindOptions) []PersonShort {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMany", query, options)
	ret0, _ := ret[0].([]PersonShort)
	return ret0
}

// FindMany indicates an expected call of FindMany.
func (mr *MockPersonStoreInterfaceMockRecorder) FindMany(query, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMany", reflect.TypeOf((*MockPersonStoreInterface)(nil).FindMany), query, options)
}

// FindOne mocks base method.
func (m *MockPersonStoreInterface) FindOne(query bson.M) PersonInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", query)
	ret0, _ := ret[0].(PersonInterface)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPersonStoreInterfaceMockRecorder) FindOne(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPersonStoreInterface)(nil).FindOne), query)
}

// FindOneAndUpdate mocks base method.
func (m *MockPersonStoreInterface) FindOneAndUpdate(query, update bson.M) PersonInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneAndUpdate", query, update)
	ret0, _ := ret[0].(PersonInterface)
	return ret0
}

// FindOneAndUpdate indicates an expected call of FindOneAndUpdate.
func (mr *MockPersonStoreInterfaceMockRecorder) FindOneAndUpdate(query, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndUpdate", reflect.TypeOf((*MockPersonStoreInterface)(nil).FindOneAndUpdate), query, update)
}

// Insert mocks base method.
func (m *MockPersonStoreInterface) Insert(information bson.M) *mongo.InsertOneResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", information)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockPersonStoreInterfaceMockRecorder) Insert(information interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPersonStoreInterface)(nil).Insert), information)
}