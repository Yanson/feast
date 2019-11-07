// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gojek/feast/protos/generated/go/feast/core (interfaces: CoreServiceClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	core "github.com/gojek/feast/protos/generated/go/feast/core"
	specs "github.com/gojek/feast/protos/generated/go/feast/specs"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCoreServiceClient is a mock of CoreServiceClient interface
type MockCoreServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCoreServiceClientMockRecorder
}

// MockCoreServiceClientMockRecorder is the mock recorder for MockCoreServiceClient
type MockCoreServiceClientMockRecorder struct {
	mock *MockCoreServiceClient
}

// NewMockCoreServiceClient creates a new mock instance
func NewMockCoreServiceClient(ctrl *gomock.Controller) *MockCoreServiceClient {
	mock := &MockCoreServiceClient{ctrl: ctrl}
	mock.recorder = &MockCoreServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreServiceClient) EXPECT() *MockCoreServiceClientMockRecorder {
	return m.recorder
}

// ApplyEntity mocks base method
func (m *MockCoreServiceClient) ApplyEntity(arg0 context.Context, arg1 *specs.EntitySpec, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_ApplyEntityResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyEntity", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_ApplyEntityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyEntity indicates an expected call of ApplyEntity
func (mr *MockCoreServiceClientMockRecorder) ApplyEntity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyEntity", reflect.TypeOf((*MockCoreServiceClient)(nil).ApplyEntity), varargs...)
}

// ApplyFeature mocks base method
func (m *MockCoreServiceClient) ApplyFeature(arg0 context.Context, arg1 *specs.FeatureSpec, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_ApplyFeatureResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyFeature", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_ApplyFeatureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyFeature indicates an expected call of ApplyFeature
func (mr *MockCoreServiceClientMockRecorder) ApplyFeature(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFeature", reflect.TypeOf((*MockCoreServiceClient)(nil).ApplyFeature), varargs...)
}

// ApplyFeatureGroup mocks base method
func (m *MockCoreServiceClient) ApplyFeatureGroup(arg0 context.Context, arg1 *specs.FeatureGroupSpec, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_ApplyFeatureGroupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyFeatureGroup", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_ApplyFeatureGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyFeatureGroup indicates an expected call of ApplyFeatureGroup
func (mr *MockCoreServiceClientMockRecorder) ApplyFeatureGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyFeatureGroup", reflect.TypeOf((*MockCoreServiceClient)(nil).ApplyFeatureGroup), varargs...)
}

// GetEntities mocks base method
func (m *MockCoreServiceClient) GetEntities(arg0 context.Context, arg1 *core.CoreServiceTypes_GetEntitiesRequest, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_GetEntitiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEntities", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_GetEntitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntities indicates an expected call of GetEntities
func (mr *MockCoreServiceClientMockRecorder) GetEntities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntities", reflect.TypeOf((*MockCoreServiceClient)(nil).GetEntities), varargs...)
}

// GetFeatures mocks base method
func (m *MockCoreServiceClient) GetFeatures(arg0 context.Context, arg1 *core.CoreServiceTypes_GetFeaturesRequest, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_GetFeaturesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFeatures", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_GetFeaturesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatures indicates an expected call of GetFeatures
func (mr *MockCoreServiceClientMockRecorder) GetFeatures(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatures", reflect.TypeOf((*MockCoreServiceClient)(nil).GetFeatures), varargs...)
}

// GetStorage mocks base method
func (m *MockCoreServiceClient) GetStorage(arg0 context.Context, arg1 *core.CoreServiceTypes_GetStorageRequest, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_GetStorageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStorage", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_GetStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorage indicates an expected call of GetStorage
func (mr *MockCoreServiceClientMockRecorder) GetStorage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorage", reflect.TypeOf((*MockCoreServiceClient)(nil).GetStorage), varargs...)
}

// GetUploadUrl mocks base method
func (m *MockCoreServiceClient) GetUploadUrl(arg0 context.Context, arg1 *core.CoreServiceTypes_GetUploadUrlRequest, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_GetUploadUrlResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUploadUrl", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_GetUploadUrlResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploadUrl indicates an expected call of GetUploadUrl
func (mr *MockCoreServiceClientMockRecorder) GetUploadUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploadUrl", reflect.TypeOf((*MockCoreServiceClient)(nil).GetUploadUrl), varargs...)
}

// ListEntities mocks base method
func (m *MockCoreServiceClient) ListEntities(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_ListEntitiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEntities", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_ListEntitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities
func (mr *MockCoreServiceClientMockRecorder) ListEntities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockCoreServiceClient)(nil).ListEntities), varargs...)
}

// ListFeatures mocks base method
func (m *MockCoreServiceClient) ListFeatures(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_ListFeaturesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFeatures", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_ListFeaturesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFeatures indicates an expected call of ListFeatures
func (mr *MockCoreServiceClientMockRecorder) ListFeatures(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeatures", reflect.TypeOf((*MockCoreServiceClient)(nil).ListFeatures), varargs...)
}

// ListStorage mocks base method
func (m *MockCoreServiceClient) ListStorage(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*core.CoreServiceTypes_ListStorageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStorage", varargs...)
	ret0, _ := ret[0].(*core.CoreServiceTypes_ListStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorage indicates an expected call of ListStorage
func (mr *MockCoreServiceClientMockRecorder) ListStorage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorage", reflect.TypeOf((*MockCoreServiceClient)(nil).ListStorage), varargs...)
}