// Code generated by MockGen. DO NOT EDIT.
// Source: pass_manager_grpc.pb.go

// Package proto is a generated GoMock package.
package proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockPassManagerClient is a mock of PassManagerClient interface.
type MockPassManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockPassManagerClientMockRecorder
}

// MockPassManagerClientMockRecorder is the mock recorder for MockPassManagerClient.
type MockPassManagerClientMockRecorder struct {
	mock *MockPassManagerClient
}

// NewMockPassManagerClient creates a new mock instance.
func 	NewMockPassManagerClient(ctrl *gomock.Controller) *MockPassManagerClient {
	mock := &MockPassManagerClient{ctrl: ctrl}
	mock.recorder = &MockPassManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPassManagerClient) EXPECT() *MockPassManagerClientMockRecorder {
	return m.recorder
}

// AddBytes mocks base method.
func (m *MockPassManagerClient) AddBytes(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddBytes", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBytes indicates an expected call of AddBytes.
func (mr *MockPassManagerClientMockRecorder) AddBytes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytes", reflect.TypeOf((*MockPassManagerClient)(nil).AddBytes), varargs...)
}

// AddCard mocks base method.
func (m *MockPassManagerClient) AddCard(ctx context.Context, in *Card, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCard", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCard indicates an expected call of AddCard.
func (mr *MockPassManagerClientMockRecorder) AddCard(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockPassManagerClient)(nil).AddCard), varargs...)
}

// AddNote mocks base method.
func (m *MockPassManagerClient) AddNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddNote", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNote indicates an expected call of AddNote.
func (mr *MockPassManagerClientMockRecorder) AddNote(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockPassManagerClient)(nil).AddNote), varargs...)
}

// AddPassword mocks base method.
func (m *MockPassManagerClient) AddPassword(ctx context.Context, in *Password, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPassword", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPassword indicates an expected call of AddPassword.
func (mr *MockPassManagerClientMockRecorder) AddPassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPassword", reflect.TypeOf((*MockPassManagerClient)(nil).AddPassword), varargs...)
}

// GetBytes mocks base method.
func (m *MockPassManagerClient) GetBytes(ctx context.Context, in *GetBytesRequest, opts ...grpc.CallOption) (*Bytes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBytes", varargs...)
	ret0, _ := ret[0].(*Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytes indicates an expected call of GetBytes.
func (mr *MockPassManagerClientMockRecorder) GetBytes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockPassManagerClient)(nil).GetBytes), varargs...)
}

// GetCard mocks base method.
func (m *MockPassManagerClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCard", varargs...)
	ret0, _ := ret[0].(*Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard.
func (mr *MockPassManagerClientMockRecorder) GetCard(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockPassManagerClient)(nil).GetCard), varargs...)
}

// GetNote mocks base method.
func (m *MockPassManagerClient) GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*Note, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNote", varargs...)
	ret0, _ := ret[0].(*Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockPassManagerClientMockRecorder) GetNote(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockPassManagerClient)(nil).GetNote), varargs...)
}

// GetPassword mocks base method.
func (m *MockPassManagerClient) GetPassword(ctx context.Context, in *GetPasswordRequest, opts ...grpc.CallOption) (*Password, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPassword", varargs...)
	ret0, _ := ret[0].(*Password)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockPassManagerClientMockRecorder) GetPassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockPassManagerClient)(nil).GetPassword), varargs...)
}

// Login mocks base method.
func (m *MockPassManagerClient) Login(ctx context.Context, in *Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockPassManagerClientMockRecorder) Login(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockPassManagerClient)(nil).Login), varargs...)
}

// Register mocks base method.
func (m *MockPassManagerClient) Register(ctx context.Context, in *Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockPassManagerClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPassManagerClient)(nil).Register), varargs...)
}

// MockPassManagerServer is a mock of PassManagerServer interface.
type MockPassManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockPassManagerServerMockRecorder
}

// MockPassManagerServerMockRecorder is the mock recorder for MockPassManagerServer.
type MockPassManagerServerMockRecorder struct {
	mock *MockPassManagerServer
}

// NewMockPassManagerServer creates a new mock instance.
func NewMockPassManagerServer(ctrl *gomock.Controller) *MockPassManagerServer {
	mock := &MockPassManagerServer{ctrl: ctrl}
	mock.recorder = &MockPassManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPassManagerServer) EXPECT() *MockPassManagerServerMockRecorder {
	return m.recorder
}

// AddBytes mocks base method.
func (m *MockPassManagerServer) AddBytes(arg0 context.Context, arg1 *Bytes) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBytes", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBytes indicates an expected call of AddBytes.
func (mr *MockPassManagerServerMockRecorder) AddBytes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytes", reflect.TypeOf((*MockPassManagerServer)(nil).AddBytes), arg0, arg1)
}

// AddCard mocks base method.
func (m *MockPassManagerServer) AddCard(arg0 context.Context, arg1 *Card) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCard", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCard indicates an expected call of AddCard.
func (mr *MockPassManagerServerMockRecorder) AddCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCard", reflect.TypeOf((*MockPassManagerServer)(nil).AddCard), arg0, arg1)
}

// AddNote mocks base method.
func (m *MockPassManagerServer) AddNote(arg0 context.Context, arg1 *Note) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNote indicates an expected call of AddNote.
func (mr *MockPassManagerServerMockRecorder) AddNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockPassManagerServer)(nil).AddNote), arg0, arg1)
}

// AddPassword mocks base method.
func (m *MockPassManagerServer) AddPassword(arg0 context.Context, arg1 *Password) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPassword", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPassword indicates an expected call of AddPassword.
func (mr *MockPassManagerServerMockRecorder) AddPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPassword", reflect.TypeOf((*MockPassManagerServer)(nil).AddPassword), arg0, arg1)
}

// GetBytes mocks base method.
func (m *MockPassManagerServer) GetBytes(arg0 context.Context, arg1 *GetBytesRequest) (*Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", arg0, arg1)
	ret0, _ := ret[0].(*Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytes indicates an expected call of GetBytes.
func (mr *MockPassManagerServerMockRecorder) GetBytes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockPassManagerServer)(nil).GetBytes), arg0, arg1)
}

// GetCard mocks base method.
func (m *MockPassManagerServer) GetCard(arg0 context.Context, arg1 *GetCardRequest) (*Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", arg0, arg1)
	ret0, _ := ret[0].(*Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard.
func (mr *MockPassManagerServerMockRecorder) GetCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockPassManagerServer)(nil).GetCard), arg0, arg1)
}

// GetNote mocks base method.
func (m *MockPassManagerServer) GetNote(arg0 context.Context, arg1 *GetNoteRequest) (*Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", arg0, arg1)
	ret0, _ := ret[0].(*Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockPassManagerServerMockRecorder) GetNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockPassManagerServer)(nil).GetNote), arg0, arg1)
}

// GetPassword mocks base method.
func (m *MockPassManagerServer) GetPassword(arg0 context.Context, arg1 *GetPasswordRequest) (*Password, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", arg0, arg1)
	ret0, _ := ret[0].(*Password)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockPassManagerServerMockRecorder) GetPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockPassManagerServer)(nil).GetPassword), arg0, arg1)
}

// Login mocks base method.
func (m *MockPassManagerServer) Login(arg0 context.Context, arg1 *Creds) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockPassManagerServerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockPassManagerServer)(nil).Login), arg0, arg1)
}

// Register mocks base method.
func (m *MockPassManagerServer) Register(arg0 context.Context, arg1 *Creds) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockPassManagerServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPassManagerServer)(nil).Register), arg0, arg1)
}

// mustEmbedUnimplementedPassManagerServer mocks base method.
func (m *MockPassManagerServer) mustEmbedUnimplementedPassManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPassManagerServer")
}

// mustEmbedUnimplementedPassManagerServer indicates an expected call of mustEmbedUnimplementedPassManagerServer.
func (mr *MockPassManagerServerMockRecorder) mustEmbedUnimplementedPassManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPassManagerServer", reflect.TypeOf((*MockPassManagerServer)(nil).mustEmbedUnimplementedPassManagerServer))
}

// MockUnsafePassManagerServer is a mock of UnsafePassManagerServer interface.
type MockUnsafePassManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePassManagerServerMockRecorder
}

// MockUnsafePassManagerServerMockRecorder is the mock recorder for MockUnsafePassManagerServer.
type MockUnsafePassManagerServerMockRecorder struct {
	mock *MockUnsafePassManagerServer
}

// NewMockUnsafePassManagerServer creates a new mock instance.
func NewMockUnsafePassManagerServer(ctrl *gomock.Controller) *MockUnsafePassManagerServer {
	mock := &MockUnsafePassManagerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePassManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePassManagerServer) EXPECT() *MockUnsafePassManagerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPassManagerServer mocks base method.
func (m *MockUnsafePassManagerServer) mustEmbedUnimplementedPassManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPassManagerServer")
}

// mustEmbedUnimplementedPassManagerServer indicates an expected call of mustEmbedUnimplementedPassManagerServer.
func (mr *MockUnsafePassManagerServerMockRecorder) mustEmbedUnimplementedPassManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPassManagerServer", reflect.TypeOf((*MockUnsafePassManagerServer)(nil).mustEmbedUnimplementedPassManagerServer))
}
