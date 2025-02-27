// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Client) Close() {
	_m.Called()
}

// Client_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Client_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Client_Expecter) Close() *Client_Close_Call {
	return &Client_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Client_Close_Call) Run(run func()) *Client_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Close_Call) Return() *Client_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *Client_Close_Call) RunAndReturn(run func()) *Client_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields:
func (_m *Client) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Client_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *Client_Expecter) Init() *Client_Init_Call {
	return &Client_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *Client_Init_Call) Run(run func()) *Client_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Init_Call) Return(_a0 error) *Client_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Init_Call) RunAndReturn(run func() error) *Client_Init_Call {
	_c.Call.Return(run)
	return _c
}

// StreamExists provides a mock function with given fields:
func (_m *Client) StreamExists() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_StreamExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamExists'
type Client_StreamExists_Call struct {
	*mock.Call
}

// StreamExists is a helper method to define mock.On call
func (_e *Client_Expecter) StreamExists() *Client_StreamExists_Call {
	return &Client_StreamExists_Call{Call: _e.mock.On("StreamExists")}
}

func (_c *Client_StreamExists_Call) Run(run func()) *Client_StreamExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_StreamExists_Call) Return(_a0 bool, _a1 error) *Client_StreamExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_StreamExists_Call) RunAndReturn(run func() (bool, error)) *Client_StreamExists_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
