// Code generated by MockGen. DO NOT EDIT.
// Source: restaurant.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/momeemt/2000s/domain/model"
)

// MockRestaurant is a mock of Restaurant interface.
type MockRestaurant struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantMockRecorder
}

// MockRestaurantMockRecorder is the mock recorder for MockRestaurant.
type MockRestaurantMockRecorder struct {
	mock *MockRestaurant
}

// NewMockRestaurant creates a new mock instance.
func NewMockRestaurant(ctrl *gomock.Controller) *MockRestaurant {
	mock := &MockRestaurant{ctrl: ctrl}
	mock.recorder = &MockRestaurantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurant) EXPECT() *MockRestaurantMockRecorder {
	return m.recorder
}

// GetNearbyRestaurants mocks base method.
func (m *MockRestaurant) GetNearbyRestaurants(arg0 model.Location, arg1 time.Time, arg2 bool) ([]model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearbyRestaurants", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearbyRestaurants indicates an expected call of GetNearbyRestaurants.
func (mr *MockRestaurantMockRecorder) GetNearbyRestaurants(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearbyRestaurants", reflect.TypeOf((*MockRestaurant)(nil).GetNearbyRestaurants), arg0, arg1, arg2)
}

// GetNextCloseTime mocks base method.
func (m *MockRestaurant) GetNextCloseTime(arg0 model.Restaurant, arg1 time.Time) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextCloseTime", arg0, arg1)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextCloseTime indicates an expected call of GetNextCloseTime.
func (mr *MockRestaurantMockRecorder) GetNextCloseTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextCloseTime", reflect.TypeOf((*MockRestaurant)(nil).GetNextCloseTime), arg0, arg1)
}

// GetRestaurantDetail mocks base method.
func (m *MockRestaurant) GetRestaurantDetail(placeId string) (model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantDetail", placeId)
	ret0, _ := ret[0].(model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantDetail indicates an expected call of GetRestaurantDetail.
func (mr *MockRestaurantMockRecorder) GetRestaurantDetail(placeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantDetail", reflect.TypeOf((*MockRestaurant)(nil).GetRestaurantDetail), placeId)
}
