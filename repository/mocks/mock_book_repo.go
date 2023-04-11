// Code generated by MockGen. DO NOT EDIT.
// Source: challenge10/repository (interfaces: BookRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "challenge10/model/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBookRepository) Create(arg0 entity.Book) (entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepository)(nil).Create), arg0)
}

// DeleteOne mocks base method.
func (m *MockBookRepository) DeleteOne(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockBookRepositoryMockRecorder) DeleteOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockBookRepository)(nil).DeleteOne), arg0)
}

// GetAll mocks base method.
func (m *MockBookRepository) GetAll() ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBookRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBookRepository)(nil).GetAll))
}

// GetOne mocks base method.
func (m *MockBookRepository) GetOne(arg0 int) (entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0)
	ret0, _ := ret[0].(entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockBookRepositoryMockRecorder) GetOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockBookRepository)(nil).GetOne), arg0)
}

// UpdateOne mocks base method.
func (m *MockBookRepository) UpdateOne(arg0 int, arg1 entity.Book) (entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", arg0, arg1)
	ret0, _ := ret[0].(entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockBookRepositoryMockRecorder) UpdateOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockBookRepository)(nil).UpdateOne), arg0, arg1)
}
