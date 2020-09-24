// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/meshplus/bitxhub-kit/crypto"
	types "github.com/meshplus/bitxhub-kit/types"
	pb "github.com/meshplus/bitxhub-model/pb"
	rpcx "github.com/meshplus/go-bitxhub-client"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Stop mocks base method
func (m *MockClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop))
}

// SetPrivateKey mocks base method
func (m *MockClient) SetPrivateKey(arg0 crypto.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPrivateKey", arg0)
}

// SetPrivateKey indicates an expected call of SetPrivateKey
func (mr *MockClientMockRecorder) SetPrivateKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrivateKey", reflect.TypeOf((*MockClient)(nil).SetPrivateKey), arg0)
}

// SendView mocks base method
func (m *MockClient) SendView(tx *pb.Transaction) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendView", tx)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendView indicates an expected call of SendView
func (mr *MockClientMockRecorder) SendView(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendView", reflect.TypeOf((*MockClient)(nil).SendView), tx)
}

// SendTransaction mocks base method
func (m *MockClient) SendTransaction(tx *pb.Transaction, opts *rpcx.TransactOpts) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", tx, opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockClientMockRecorder) SendTransaction(tx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockClient)(nil).SendTransaction), tx, opts)
}

// SendTransactionWithReceipt mocks base method
func (m *MockClient) SendTransactionWithReceipt(tx *pb.Transaction, opts *rpcx.TransactOpts) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransactionWithReceipt", tx, opts)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransactionWithReceipt indicates an expected call of SendTransactionWithReceipt
func (mr *MockClientMockRecorder) SendTransactionWithReceipt(tx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransactionWithReceipt", reflect.TypeOf((*MockClient)(nil).SendTransactionWithReceipt), tx, opts)
}

// GetReceipt mocks base method
func (m *MockClient) GetReceipt(hash string) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", hash)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt
func (mr *MockClientMockRecorder) GetReceipt(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockClient)(nil).GetReceipt), hash)
}

// GetTransaction mocks base method
func (m *MockClient) GetTransaction(hash string) (*pb.GetTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", hash)
	ret0, _ := ret[0].(*pb.GetTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockClientMockRecorder) GetTransaction(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockClient)(nil).GetTransaction), hash)
}

// GetChainMeta mocks base method
func (m *MockClient) GetChainMeta() (*pb.ChainMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainMeta")
	ret0, _ := ret[0].(*pb.ChainMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainMeta indicates an expected call of GetChainMeta
func (mr *MockClientMockRecorder) GetChainMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainMeta", reflect.TypeOf((*MockClient)(nil).GetChainMeta))
}

// GetBlocks mocks base method
func (m *MockClient) GetBlocks(start, end uint64) (*pb.GetBlocksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocks", start, end)
	ret0, _ := ret[0].(*pb.GetBlocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocks indicates an expected call of GetBlocks
func (mr *MockClientMockRecorder) GetBlocks(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocks", reflect.TypeOf((*MockClient)(nil).GetBlocks), start, end)
}

// GetBlock mocks base method
func (m *MockClient) GetBlock(value string, blockType pb.GetBlockRequest_Type) (*pb.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", value, blockType)
	ret0, _ := ret[0].(*pb.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockClientMockRecorder) GetBlock(value, blockType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockClient)(nil).GetBlock), value, blockType)
}

// GetChainStatus mocks base method
func (m *MockClient) GetChainStatus() (*pb.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainStatus")
	ret0, _ := ret[0].(*pb.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainStatus indicates an expected call of GetChainStatus
func (mr *MockClientMockRecorder) GetChainStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainStatus", reflect.TypeOf((*MockClient)(nil).GetChainStatus))
}

// GetValidators mocks base method
func (m *MockClient) GetValidators() (*pb.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidators")
	ret0, _ := ret[0].(*pb.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidators indicates an expected call of GetValidators
func (mr *MockClientMockRecorder) GetValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidators", reflect.TypeOf((*MockClient)(nil).GetValidators))
}

// GetNetworkMeta mocks base method
func (m *MockClient) GetNetworkMeta() (*pb.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkMeta")
	ret0, _ := ret[0].(*pb.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkMeta indicates an expected call of GetNetworkMeta
func (mr *MockClientMockRecorder) GetNetworkMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkMeta", reflect.TypeOf((*MockClient)(nil).GetNetworkMeta))
}

// GetAccountBalance mocks base method
func (m *MockClient) GetAccountBalance(address string) (*pb.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountBalance", address)
	ret0, _ := ret[0].(*pb.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountBalance indicates an expected call of GetAccountBalance
func (mr *MockClientMockRecorder) GetAccountBalance(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountBalance", reflect.TypeOf((*MockClient)(nil).GetAccountBalance), address)
}

// GetBlockHeader mocks base method
func (m *MockClient) GetBlockHeader(ctx context.Context, begin, end uint64, ch chan<- *pb.BlockHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHeader", ctx, begin, end, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetBlockHeader indicates an expected call of GetBlockHeader
func (mr *MockClientMockRecorder) GetBlockHeader(ctx, begin, end, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHeader", reflect.TypeOf((*MockClient)(nil).GetBlockHeader), ctx, begin, end, ch)
}

// GetInterchainTxWrappers mocks base method
func (m *MockClient) GetInterchainTxWrappers(ctx context.Context, pid string, begin, end uint64, ch chan<- *pb.InterchainTxWrappers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterchainTxWrappers", ctx, pid, begin, end, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetInterchainTxWrappers indicates an expected call of GetInterchainTxWrappers
func (mr *MockClientMockRecorder) GetInterchainTxWrappers(ctx, pid, begin, end, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterchainTxWrappers", reflect.TypeOf((*MockClient)(nil).GetInterchainTxWrappers), ctx, pid, begin, end, ch)
}

// Subscribe mocks base method
func (m *MockClient) Subscribe(arg0 context.Context, arg1 pb.SubscriptionRequest_Type, arg2 []byte) (<-chan interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockClientMockRecorder) Subscribe(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockClient)(nil).Subscribe), arg0, arg1, arg2)
}

// DeployContract mocks base method
func (m *MockClient) DeployContract(contract []byte, opts *rpcx.TransactOpts) (types.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployContract", contract, opts)
	ret0, _ := ret[0].(types.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployContract indicates an expected call of DeployContract
func (mr *MockClientMockRecorder) DeployContract(contract, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployContract", reflect.TypeOf((*MockClient)(nil).DeployContract), contract, opts)
}

// GenerateContractTx mocks base method
func (m *MockClient) GenerateContractTx(vmType pb.TransactionData_VMType, address types.Address, method string, args ...*pb.Arg) (*pb.Transaction, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{vmType, address, method}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GenerateContractTx", varargs...)
	ret0, _ := ret[0].(*pb.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateContractTx indicates an expected call of GenerateContractTx
func (mr *MockClientMockRecorder) GenerateContractTx(vmType, address, method interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{vmType, address, method}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateContractTx", reflect.TypeOf((*MockClient)(nil).GenerateContractTx), varargs...)
}

// InvokeContract mocks base method
func (m *MockClient) InvokeContract(vmType pb.TransactionData_VMType, address types.Address, method string, opts *rpcx.TransactOpts, args ...*pb.Arg) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{vmType, address, method, opts}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeContract", varargs...)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeContract indicates an expected call of InvokeContract
func (mr *MockClientMockRecorder) InvokeContract(vmType, address, method, opts interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{vmType, address, method, opts}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeContract", reflect.TypeOf((*MockClient)(nil).InvokeContract), varargs...)
}

// InvokeBVMContract mocks base method
func (m *MockClient) InvokeBVMContract(address types.Address, method string, opts *rpcx.TransactOpts, args ...*pb.Arg) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{address, method, opts}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeBVMContract", varargs...)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeBVMContract indicates an expected call of InvokeBVMContract
func (mr *MockClientMockRecorder) InvokeBVMContract(address, method, opts interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{address, method, opts}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeBVMContract", reflect.TypeOf((*MockClient)(nil).InvokeBVMContract), varargs...)
}

// InvokeXVMContract mocks base method
func (m *MockClient) InvokeXVMContract(address types.Address, method string, opts *rpcx.TransactOpts, args ...*pb.Arg) (*pb.Receipt, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{address, method, opts}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeXVMContract", varargs...)
	ret0, _ := ret[0].(*pb.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeXVMContract indicates an expected call of InvokeXVMContract
func (mr *MockClientMockRecorder) InvokeXVMContract(address, method, opts interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{address, method, opts}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeXVMContract", reflect.TypeOf((*MockClient)(nil).InvokeXVMContract), varargs...)
}

// GetMultiSigns mocks base method
func (m *MockClient) GetMultiSigns(id string, typ pb.GetMultiSignsRequest_Type) (*pb.SignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultiSigns", id, typ)
	ret0, _ := ret[0].(*pb.SignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiSigns indicates an expected call of GetMultiSigns
func (mr *MockClientMockRecorder) GetMultiSigns(id, typ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiSigns", reflect.TypeOf((*MockClient)(nil).GetMultiSigns), id, typ)
}

// GetTPS mocks base method
func (m *MockClient) GetTPS(begin, end uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTPS", begin, end)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTPS indicates an expected call of GetTPS
func (mr *MockClientMockRecorder) GetTPS(begin, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTPS", reflect.TypeOf((*MockClient)(nil).GetTPS), begin, end)
}

// GetPendingNonceByAccount mocks base method
func (m *MockClient) GetPendingNonceByAccount(account string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingNonceByAccount", account)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingNonceByAccount indicates an expected call of GetPendingNonceByAccount
func (mr *MockClientMockRecorder) GetPendingNonceByAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingNonceByAccount", reflect.TypeOf((*MockClient)(nil).GetPendingNonceByAccount), account)
}
