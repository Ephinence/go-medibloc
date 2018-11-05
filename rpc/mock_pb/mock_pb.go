// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/medibloc/go-medibloc/rpc/pb (interfaces: ApiServiceClient,ApiService_SubscribeClient,ApiServiceServer,ApiService_SubscribeServer)

// Package mock_pb is a generated GoMock package.
package mock_pb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/medibloc/go-medibloc/rpc/pb"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockApiServiceClient is a mock of ApiServiceClient interface
type MockApiServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceClientMockRecorder
}

// MockApiServiceClientMockRecorder is the mock recorder for MockApiServiceClient
type MockApiServiceClientMockRecorder struct {
	mock *MockApiServiceClient
}

// NewMockApiServiceClient creates a new mock instance
func NewMockApiServiceClient(ctrl *gomock.Controller) *MockApiServiceClient {
	mock := &MockApiServiceClient{ctrl: ctrl}
	mock.recorder = &MockApiServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceClient) EXPECT() *MockApiServiceClientMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockApiServiceClient) GetAccount(arg0 context.Context, arg1 *pb.GetAccountRequest, arg2 ...grpc.CallOption) (*pb.GetAccountResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccount", varargs...)
	ret0, _ := ret[0].(*pb.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockApiServiceClientMockRecorder) GetAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockApiServiceClient)(nil).GetAccount), varargs...)
}

// GetAccountTransactions mocks base method
func (m *MockApiServiceClient) GetAccountTransactions(arg0 context.Context, arg1 *pb.GetAccountTransactionsRequest, arg2 ...grpc.CallOption) (*pb.GetTransactionsResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountTransactions", varargs...)
	ret0, _ := ret[0].(*pb.GetTransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountTransactions indicates an expected call of GetAccountTransactions
func (mr *MockApiServiceClientMockRecorder) GetAccountTransactions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountTransactions", reflect.TypeOf((*MockApiServiceClient)(nil).GetAccountTransactions), varargs...)
}

// GetBlock mocks base method
func (m *MockApiServiceClient) GetBlock(arg0 context.Context, arg1 *pb.GetBlockRequest, arg2 ...grpc.CallOption) (*pb.GetBlockResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlock", varargs...)
	ret0, _ := ret[0].(*pb.GetBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockApiServiceClientMockRecorder) GetBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockApiServiceClient)(nil).GetBlock), varargs...)
}

// GetBlocks mocks base method
func (m *MockApiServiceClient) GetBlocks(arg0 context.Context, arg1 *pb.GetBlocksRequest, arg2 ...grpc.CallOption) (*pb.GetBlocksResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlocks", varargs...)
	ret0, _ := ret[0].(*pb.GetBlocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocks indicates an expected call of GetBlocks
func (mr *MockApiServiceClientMockRecorder) GetBlocks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocks", reflect.TypeOf((*MockApiServiceClient)(nil).GetBlocks), varargs...)
}

// GetCandidates mocks base method
func (m *MockApiServiceClient) GetCandidates(arg0 context.Context, arg1 *pb.NonParamRequest, arg2 ...grpc.CallOption) (*pb.GetCandidatesResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCandidates", varargs...)
	ret0, _ := ret[0].(*pb.GetCandidatesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidates indicates an expected call of GetCandidates
func (mr *MockApiServiceClientMockRecorder) GetCandidates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidates", reflect.TypeOf((*MockApiServiceClient)(nil).GetCandidates), varargs...)
}

// GetDynasty mocks base method
func (m *MockApiServiceClient) GetDynasty(arg0 context.Context, arg1 *pb.NonParamRequest, arg2 ...grpc.CallOption) (*pb.GetDynastyResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDynasty", varargs...)
	ret0, _ := ret[0].(*pb.GetDynastyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDynasty indicates an expected call of GetDynasty
func (mr *MockApiServiceClientMockRecorder) GetDynasty(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDynasty", reflect.TypeOf((*MockApiServiceClient)(nil).GetDynasty), varargs...)
}

// GetMedState mocks base method
func (m *MockApiServiceClient) GetMedState(arg0 context.Context, arg1 *pb.NonParamRequest, arg2 ...grpc.CallOption) (*pb.GetMedStateResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMedState", varargs...)
	ret0, _ := ret[0].(*pb.GetMedStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMedState indicates an expected call of GetMedState
func (mr *MockApiServiceClientMockRecorder) GetMedState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMedState", reflect.TypeOf((*MockApiServiceClient)(nil).GetMedState), varargs...)
}

// GetPendingTransactions mocks base method
func (m *MockApiServiceClient) GetPendingTransactions(arg0 context.Context, arg1 *pb.NonParamRequest, arg2 ...grpc.CallOption) (*pb.GetTransactionsResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPendingTransactions", varargs...)
	ret0, _ := ret[0].(*pb.GetTransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingTransactions indicates an expected call of GetPendingTransactions
func (mr *MockApiServiceClientMockRecorder) GetPendingTransactions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingTransactions", reflect.TypeOf((*MockApiServiceClient)(nil).GetPendingTransactions), varargs...)
}

// GetTransaction mocks base method
func (m *MockApiServiceClient) GetTransaction(arg0 context.Context, arg1 *pb.GetTransactionRequest, arg2 ...grpc.CallOption) (*pb.GetTransactionResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTransaction", varargs...)
	ret0, _ := ret[0].(*pb.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockApiServiceClientMockRecorder) GetTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockApiServiceClient)(nil).GetTransaction), varargs...)
}

// GetTransactionReceipt mocks base method
func (m *MockApiServiceClient) GetTransactionReceipt(arg0 context.Context, arg1 *pb.GetTransactionRequest, arg2 ...grpc.CallOption) (*pb.GetTransactionReceiptResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTransactionReceipt", varargs...)
	ret0, _ := ret[0].(*pb.GetTransactionReceiptResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionReceipt indicates an expected call of GetTransactionReceipt
func (mr *MockApiServiceClientMockRecorder) GetTransactionReceipt(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceipt", reflect.TypeOf((*MockApiServiceClient)(nil).GetTransactionReceipt), varargs...)
}

// HealthCheck mocks base method
func (m *MockApiServiceClient) HealthCheck(arg0 context.Context, arg1 *pb.NonParamRequest, arg2 ...grpc.CallOption) (*pb.HealthCheckResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HealthCheck", varargs...)
	ret0, _ := ret[0].(*pb.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockApiServiceClientMockRecorder) HealthCheck(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockApiServiceClient)(nil).HealthCheck), varargs...)
}

// SendTransaction mocks base method
func (m *MockApiServiceClient) SendTransaction(arg0 context.Context, arg1 *pb.SendTransactionRequest, arg2 ...grpc.CallOption) (*pb.SendTransactionResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendTransaction", varargs...)
	ret0, _ := ret[0].(*pb.SendTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockApiServiceClientMockRecorder) SendTransaction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockApiServiceClient)(nil).SendTransaction), varargs...)
}

// Subscribe mocks base method
func (m *MockApiServiceClient) Subscribe(arg0 context.Context, arg1 *pb.SubscribeRequest, arg2 ...grpc.CallOption) (pb.ApiService_SubscribeClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(pb.ApiService_SubscribeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockApiServiceClientMockRecorder) Subscribe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockApiServiceClient)(nil).Subscribe), varargs...)
}

// MockApiService_SubscribeClient is a mock of ApiService_SubscribeClient interface
type MockApiService_SubscribeClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiService_SubscribeClientMockRecorder
}

// MockApiService_SubscribeClientMockRecorder is the mock recorder for MockApiService_SubscribeClient
type MockApiService_SubscribeClientMockRecorder struct {
	mock *MockApiService_SubscribeClient
}

// NewMockApiService_SubscribeClient creates a new mock instance
func NewMockApiService_SubscribeClient(ctrl *gomock.Controller) *MockApiService_SubscribeClient {
	mock := &MockApiService_SubscribeClient{ctrl: ctrl}
	mock.recorder = &MockApiService_SubscribeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiService_SubscribeClient) EXPECT() *MockApiService_SubscribeClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockApiService_SubscribeClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockApiService_SubscribeClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockApiService_SubscribeClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockApiService_SubscribeClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).Context))
}

// Header mocks base method
func (m *MockApiService_SubscribeClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockApiService_SubscribeClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).Header))
}

// Recv mocks base method
func (m *MockApiService_SubscribeClient) Recv() (*pb.SubscribeResponse, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*pb.SubscribeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockApiService_SubscribeClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockApiService_SubscribeClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockApiService_SubscribeClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockApiService_SubscribeClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockApiService_SubscribeClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockApiService_SubscribeClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockApiService_SubscribeClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockApiService_SubscribeClient)(nil).Trailer))
}

// MockApiServiceServer is a mock of ApiServiceServer interface
type MockApiServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceServerMockRecorder
}

// MockApiServiceServerMockRecorder is the mock recorder for MockApiServiceServer
type MockApiServiceServerMockRecorder struct {
	mock *MockApiServiceServer
}

// NewMockApiServiceServer creates a new mock instance
func NewMockApiServiceServer(ctrl *gomock.Controller) *MockApiServiceServer {
	mock := &MockApiServiceServer{ctrl: ctrl}
	mock.recorder = &MockApiServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceServer) EXPECT() *MockApiServiceServerMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockApiServiceServer) GetAccount(arg0 context.Context, arg1 *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockApiServiceServerMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockApiServiceServer)(nil).GetAccount), arg0, arg1)
}

// GetAccountTransactions mocks base method
func (m *MockApiServiceServer) GetAccountTransactions(arg0 context.Context, arg1 *pb.GetAccountTransactionsRequest) (*pb.GetTransactionsResponse, error) {
	ret := m.ctrl.Call(m, "GetAccountTransactions", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountTransactions indicates an expected call of GetAccountTransactions
func (mr *MockApiServiceServerMockRecorder) GetAccountTransactions(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountTransactions", reflect.TypeOf((*MockApiServiceServer)(nil).GetAccountTransactions), arg0, arg1)
}

// GetBlock mocks base method
func (m *MockApiServiceServer) GetBlock(arg0 context.Context, arg1 *pb.GetBlockRequest) (*pb.GetBlockResponse, error) {
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockApiServiceServerMockRecorder) GetBlock(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockApiServiceServer)(nil).GetBlock), arg0, arg1)
}

// GetBlocks mocks base method
func (m *MockApiServiceServer) GetBlocks(arg0 context.Context, arg1 *pb.GetBlocksRequest) (*pb.GetBlocksResponse, error) {
	ret := m.ctrl.Call(m, "GetBlocks", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetBlocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocks indicates an expected call of GetBlocks
func (mr *MockApiServiceServerMockRecorder) GetBlocks(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocks", reflect.TypeOf((*MockApiServiceServer)(nil).GetBlocks), arg0, arg1)
}

// GetCandidates mocks base method
func (m *MockApiServiceServer) GetCandidates(arg0 context.Context, arg1 *pb.NonParamRequest) (*pb.GetCandidatesResponse, error) {
	ret := m.ctrl.Call(m, "GetCandidates", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetCandidatesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidates indicates an expected call of GetCandidates
func (mr *MockApiServiceServerMockRecorder) GetCandidates(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidates", reflect.TypeOf((*MockApiServiceServer)(nil).GetCandidates), arg0, arg1)
}

// GetDynasty mocks base method
func (m *MockApiServiceServer) GetDynasty(arg0 context.Context, arg1 *pb.NonParamRequest) (*pb.GetDynastyResponse, error) {
	ret := m.ctrl.Call(m, "GetDynasty", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetDynastyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDynasty indicates an expected call of GetDynasty
func (mr *MockApiServiceServerMockRecorder) GetDynasty(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDynasty", reflect.TypeOf((*MockApiServiceServer)(nil).GetDynasty), arg0, arg1)
}

// GetMedState mocks base method
func (m *MockApiServiceServer) GetMedState(arg0 context.Context, arg1 *pb.NonParamRequest) (*pb.GetMedStateResponse, error) {
	ret := m.ctrl.Call(m, "GetMedState", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetMedStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMedState indicates an expected call of GetMedState
func (mr *MockApiServiceServerMockRecorder) GetMedState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMedState", reflect.TypeOf((*MockApiServiceServer)(nil).GetMedState), arg0, arg1)
}

// GetPendingTransactions mocks base method
func (m *MockApiServiceServer) GetPendingTransactions(arg0 context.Context, arg1 *pb.NonParamRequest) (*pb.GetTransactionsResponse, error) {
	ret := m.ctrl.Call(m, "GetPendingTransactions", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTransactionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingTransactions indicates an expected call of GetPendingTransactions
func (mr *MockApiServiceServerMockRecorder) GetPendingTransactions(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingTransactions", reflect.TypeOf((*MockApiServiceServer)(nil).GetPendingTransactions), arg0, arg1)
}

// GetTransaction mocks base method
func (m *MockApiServiceServer) GetTransaction(arg0 context.Context, arg1 *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	ret := m.ctrl.Call(m, "GetTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockApiServiceServerMockRecorder) GetTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockApiServiceServer)(nil).GetTransaction), arg0, arg1)
}

// GetTransactionReceipt mocks base method
func (m *MockApiServiceServer) GetTransactionReceipt(arg0 context.Context, arg1 *pb.GetTransactionRequest) (*pb.GetTransactionReceiptResponse, error) {
	ret := m.ctrl.Call(m, "GetTransactionReceipt", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTransactionReceiptResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionReceipt indicates an expected call of GetTransactionReceipt
func (mr *MockApiServiceServerMockRecorder) GetTransactionReceipt(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceipt", reflect.TypeOf((*MockApiServiceServer)(nil).GetTransactionReceipt), arg0, arg1)
}

// HealthCheck mocks base method
func (m *MockApiServiceServer) HealthCheck(arg0 context.Context, arg1 *pb.NonParamRequest) (*pb.HealthCheckResponse, error) {
	ret := m.ctrl.Call(m, "HealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*pb.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockApiServiceServerMockRecorder) HealthCheck(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockApiServiceServer)(nil).HealthCheck), arg0, arg1)
}

// SendTransaction mocks base method
func (m *MockApiServiceServer) SendTransaction(arg0 context.Context, arg1 *pb.SendTransactionRequest) (*pb.SendTransactionResponse, error) {
	ret := m.ctrl.Call(m, "SendTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.SendTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockApiServiceServerMockRecorder) SendTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockApiServiceServer)(nil).SendTransaction), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockApiServiceServer) Subscribe(arg0 *pb.SubscribeRequest, arg1 pb.ApiService_SubscribeServer) error {
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockApiServiceServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockApiServiceServer)(nil).Subscribe), arg0, arg1)
}

// MockApiService_SubscribeServer is a mock of ApiService_SubscribeServer interface
type MockApiService_SubscribeServer struct {
	ctrl     *gomock.Controller
	recorder *MockApiService_SubscribeServerMockRecorder
}

// MockApiService_SubscribeServerMockRecorder is the mock recorder for MockApiService_SubscribeServer
type MockApiService_SubscribeServerMockRecorder struct {
	mock *MockApiService_SubscribeServer
}

// NewMockApiService_SubscribeServer creates a new mock instance
func NewMockApiService_SubscribeServer(ctrl *gomock.Controller) *MockApiService_SubscribeServer {
	mock := &MockApiService_SubscribeServer{ctrl: ctrl}
	mock.recorder = &MockApiService_SubscribeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiService_SubscribeServer) EXPECT() *MockApiService_SubscribeServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockApiService_SubscribeServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockApiService_SubscribeServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockApiService_SubscribeServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockApiService_SubscribeServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockApiService_SubscribeServer) Send(arg0 *pb.SubscribeResponse) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockApiService_SubscribeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockApiService_SubscribeServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockApiService_SubscribeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockApiService_SubscribeServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockApiService_SubscribeServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockApiService_SubscribeServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockApiService_SubscribeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockApiService_SubscribeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockApiService_SubscribeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockApiService_SubscribeServer)(nil).SetTrailer), arg0)
}
