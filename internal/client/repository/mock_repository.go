// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddBytes mocks base method.
func (m *MockRepository) AddBytes(ctx context.Context, binaryData BinaryData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBytes", ctx, binaryData)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBytes indicates an expected call of AddBytes.
func (mr *MockRepositoryMockRecorder) AddBytes(ctx, binaryData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytes", reflect.TypeOf((*MockRepository)(nil).AddBytes), ctx, binaryData)
}

// AddCard mocks base method.
func (m *MockRepository) AddCard(ctx context.Context, card Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", ctx, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCard indicates an expected call of AddCard.
func (mr *MockRepositoryMockRecorder) AddCard(ctx, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockRepository)(nil).AddCard), ctx, card)
}

// AddNote mocks base method.
func (m *MockRepository) AddNote(ctx context.Context, note Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", ctx, note)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockRepositoryMockRecorder) AddNote(ctx, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockRepository)(nil).AddNote), ctx, note)
}

// AddPassword mocks base method.
func (m *MockRepository) AddPassword(ctx context.Context, pass Password) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPassword", ctx, pass)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPassword indicates an expected call of AddPassword.
func (mr *MockRepositoryMockRecorder) AddPassword(ctx, pass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPassword", reflect.TypeOf((*MockRepository)(nil).AddPassword), ctx, pass)
}

// Close mocks base method.
func (m *MockRepository) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepository)(nil).Close))
}

// GetBytes mocks base method.
func (m *MockRepository) GetBytes(ctx context.Context, bytesName string) (BinaryData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", ctx, bytesName)
	ret0, _ := ret[0].(BinaryData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytes indicates an expected call of GetBytes.
func (mr *MockRepositoryMockRecorder) GetBytes(ctx, bytesName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockRepository)(nil).GetBytes), ctx, bytesName)
}

// GetCard mocks base method.
func (m *MockRepository) GetCard(ctx context.Context, cardName string) (Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", ctx, cardName)
	ret0, _ := ret[0].(Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard.
func (mr *MockRepositoryMockRecorder) GetCard(ctx, cardName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockRepository)(nil).GetCard), ctx, cardName)
}

// GetNote mocks base method.
func (m *MockRepository) GetNote(ctx context.Context, noteName string) (Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", ctx, noteName)
	ret0, _ := ret[0].(Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockRepositoryMockRecorder) GetNote(ctx, noteName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockRepository)(nil).GetNote), ctx, noteName)
}

// GetPassword mocks base method.
func (m *MockRepository) GetPassword(ctx context.Context, name string) (Password, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", ctx, name)
	ret0, _ := ret[0].(Password)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockRepositoryMockRecorder) GetPassword(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockRepository)(nil).GetPassword), ctx, name)
}

// Login mocks base method.
func (m *MockRepository) Login(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockRepositoryMockRecorder) Login(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRepository)(nil).Login), ctx, login, password)
}

// Register mocks base method.
func (m *MockRepository) Register(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockRepositoryMockRecorder) Register(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRepository)(nil).Register), ctx, login, password)
}
