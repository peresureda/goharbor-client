// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	io "io"

	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	systeminfo "github.com/peresureda/goharbor-client/v5/apiv2/internal/api/client/systeminfo"
)

// MockSysteminfoClientService is an autogenerated mock type for the ClientService type
type MockSysteminfoClientService struct {
	mock.Mock
}

// GetCert provides a mock function with given fields: params, authInfo, writer
func (_m *MockSysteminfoClientService) GetCert(params *systeminfo.GetCertParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*systeminfo.GetCertOK, error) {
	ret := _m.Called(params, authInfo, writer)

	var r0 *systeminfo.GetCertOK
	if rf, ok := ret.Get(0).(func(*systeminfo.GetCertParams, runtime.ClientAuthInfoWriter, io.Writer) *systeminfo.GetCertOK); ok {
		r0 = rf(params, authInfo, writer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*systeminfo.GetCertOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*systeminfo.GetCertParams, runtime.ClientAuthInfoWriter, io.Writer) error); ok {
		r1 = rf(params, authInfo, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSystemInfo provides a mock function with given fields: params, authInfo
func (_m *MockSysteminfoClientService) GetSystemInfo(params *systeminfo.GetSystemInfoParams, authInfo runtime.ClientAuthInfoWriter) (*systeminfo.GetSystemInfoOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *systeminfo.GetSystemInfoOK
	if rf, ok := ret.Get(0).(func(*systeminfo.GetSystemInfoParams, runtime.ClientAuthInfoWriter) *systeminfo.GetSystemInfoOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*systeminfo.GetSystemInfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*systeminfo.GetSystemInfoParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolumes provides a mock function with given fields: params, authInfo
func (_m *MockSysteminfoClientService) GetVolumes(params *systeminfo.GetVolumesParams, authInfo runtime.ClientAuthInfoWriter) (*systeminfo.GetVolumesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *systeminfo.GetVolumesOK
	if rf, ok := ret.Get(0).(func(*systeminfo.GetVolumesParams, runtime.ClientAuthInfoWriter) *systeminfo.GetVolumesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*systeminfo.GetVolumesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*systeminfo.GetVolumesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockSysteminfoClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockSysteminfoClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSysteminfoClientService creates a new instance of MockSysteminfoClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSysteminfoClientService(t mockConstructorTestingTNewMockSysteminfoClientService) *MockSysteminfoClientService {
	mock := &MockSysteminfoClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
