// Code generated by mockery v1.0.0. DO NOT EDIT.

package scope

import mock "github.com/stretchr/testify/mock"

// MockObject is an autogenerated mock type for the Object type
type MockObject struct {
	mock.Mock
}

// Entry provides a mock function with given fields:
func (_m *MockObject) Entry() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IncludeAuthor provides a mock function with given fields:
func (_m *MockObject) IncludeAuthor() {
	_m.Called()
}

// Labels provides a mock function with given fields:
func (_m *MockObject) Labels() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}