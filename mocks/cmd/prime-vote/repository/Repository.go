// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	model "github.com/atomgunlk/prime-vote/cmd/prime-vote/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// ClearVoteItem provides a mock function with given fields: req
func (_m *Repository) ClearVoteItem(req *model.ClearVoteItemRequest) (*model.ClearVoteItemResponse, error) {
	ret := _m.Called(req)

	var r0 *model.ClearVoteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ClearVoteItemRequest) (*model.ClearVoteItemResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.ClearVoteItemRequest) *model.ClearVoteItemResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClearVoteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.ClearVoteItemRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_ClearVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearVoteItem'
type Repository_ClearVoteItem_Call struct {
	*mock.Call
}

// ClearVoteItem is a helper method to define mock.On call
//   - req *model.ClearVoteItemRequest
func (_e *Repository_Expecter) ClearVoteItem(req interface{}) *Repository_ClearVoteItem_Call {
	return &Repository_ClearVoteItem_Call{Call: _e.mock.On("ClearVoteItem", req)}
}

func (_c *Repository_ClearVoteItem_Call) Run(run func(req *model.ClearVoteItemRequest)) *Repository_ClearVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.ClearVoteItemRequest))
	})
	return _c
}

func (_c *Repository_ClearVoteItem_Call) Return(_a0 *model.ClearVoteItemResponse, _a1 error) *Repository_ClearVoteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_ClearVoteItem_Call) RunAndReturn(run func(*model.ClearVoteItemRequest) (*model.ClearVoteItemResponse, error)) *Repository_ClearVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *Repository) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Repository_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Repository_Expecter) Close() *Repository_Close_Call {
	return &Repository_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Repository_Close_Call) Run(run func()) *Repository_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_Close_Call) Return(_a0 error) *Repository_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Close_Call) RunAndReturn(run func() error) *Repository_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CreateVoteItem provides a mock function with given fields: _a0
func (_m *Repository) CreateVoteItem(_a0 *model.CreateVoteItemRequest) (*model.CreateVoteItemResponse, error) {
	ret := _m.Called(_a0)

	var r0 *model.CreateVoteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.CreateVoteItemRequest) (*model.CreateVoteItemResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.CreateVoteItemRequest) *model.CreateVoteItemResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CreateVoteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.CreateVoteItemRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_CreateVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateVoteItem'
type Repository_CreateVoteItem_Call struct {
	*mock.Call
}

// CreateVoteItem is a helper method to define mock.On call
//   - _a0 *model.CreateVoteItemRequest
func (_e *Repository_Expecter) CreateVoteItem(_a0 interface{}) *Repository_CreateVoteItem_Call {
	return &Repository_CreateVoteItem_Call{Call: _e.mock.On("CreateVoteItem", _a0)}
}

func (_c *Repository_CreateVoteItem_Call) Run(run func(_a0 *model.CreateVoteItemRequest)) *Repository_CreateVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.CreateVoteItemRequest))
	})
	return _c
}

func (_c *Repository_CreateVoteItem_Call) Return(_a0 *model.CreateVoteItemResponse, _a1 error) *Repository_CreateVoteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_CreateVoteItem_Call) RunAndReturn(run func(*model.CreateVoteItemRequest) (*model.CreateVoteItemResponse, error)) *Repository_CreateVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteVoteItem provides a mock function with given fields: req
func (_m *Repository) DeleteVoteItem(req *model.DeleteVoteItemRequest) (*model.DeleteVoteItemResponse, error) {
	ret := _m.Called(req)

	var r0 *model.DeleteVoteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.DeleteVoteItemRequest) (*model.DeleteVoteItemResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.DeleteVoteItemRequest) *model.DeleteVoteItemResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeleteVoteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.DeleteVoteItemRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_DeleteVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteVoteItem'
type Repository_DeleteVoteItem_Call struct {
	*mock.Call
}

// DeleteVoteItem is a helper method to define mock.On call
//   - req *model.DeleteVoteItemRequest
func (_e *Repository_Expecter) DeleteVoteItem(req interface{}) *Repository_DeleteVoteItem_Call {
	return &Repository_DeleteVoteItem_Call{Call: _e.mock.On("DeleteVoteItem", req)}
}

func (_c *Repository_DeleteVoteItem_Call) Run(run func(req *model.DeleteVoteItemRequest)) *Repository_DeleteVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DeleteVoteItemRequest))
	})
	return _c
}

func (_c *Repository_DeleteVoteItem_Call) Return(_a0 *model.DeleteVoteItemResponse, _a1 error) *Repository_DeleteVoteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_DeleteVoteItem_Call) RunAndReturn(run func(*model.DeleteVoteItemRequest) (*model.DeleteVoteItemResponse, error)) *Repository_DeleteVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: id
func (_m *Repository) GetUserByID(id uint64) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*model.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type Repository_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - id uint64
func (_e *Repository_Expecter) GetUserByID(id interface{}) *Repository_GetUserByID_Call {
	return &Repository_GetUserByID_Call{Call: _e.mock.On("GetUserByID", id)}
}

func (_c *Repository_GetUserByID_Call) Run(run func(id uint64)) *Repository_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *Repository_GetUserByID_Call) Return(_a0 *model.User, _a1 error) *Repository_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetUserByID_Call) RunAndReturn(run func(uint64) (*model.User, error)) *Repository_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *Repository) GetUserByUsername(username string) (*model.User, error) {
	ret := _m.Called(username)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type Repository_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - username string
func (_e *Repository_Expecter) GetUserByUsername(username interface{}) *Repository_GetUserByUsername_Call {
	return &Repository_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", username)}
}

func (_c *Repository_GetUserByUsername_Call) Run(run func(username string)) *Repository_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Repository_GetUserByUsername_Call) Return(_a0 *model.User, _a1 error) *Repository_GetUserByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetUserByUsername_Call) RunAndReturn(run func(string) (*model.User, error)) *Repository_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetVoteItem provides a mock function with given fields: req
func (_m *Repository) GetVoteItem(req *model.GetVoteItemRequest) (*model.GetVoteItemResponse, error) {
	ret := _m.Called(req)

	var r0 *model.GetVoteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.GetVoteItemRequest) (*model.GetVoteItemResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.GetVoteItemRequest) *model.GetVoteItemResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GetVoteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.GetVoteItemRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVoteItem'
type Repository_GetVoteItem_Call struct {
	*mock.Call
}

// GetVoteItem is a helper method to define mock.On call
//   - req *model.GetVoteItemRequest
func (_e *Repository_Expecter) GetVoteItem(req interface{}) *Repository_GetVoteItem_Call {
	return &Repository_GetVoteItem_Call{Call: _e.mock.On("GetVoteItem", req)}
}

func (_c *Repository_GetVoteItem_Call) Run(run func(req *model.GetVoteItemRequest)) *Repository_GetVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.GetVoteItemRequest))
	})
	return _c
}

func (_c *Repository_GetVoteItem_Call) Return(_a0 *model.GetVoteItemResponse, _a1 error) *Repository_GetVoteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetVoteItem_Call) RunAndReturn(run func(*model.GetVoteItemRequest) (*model.GetVoteItemResponse, error)) *Repository_GetVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetVoteResult provides a mock function with given fields: req
func (_m *Repository) GetVoteResult(req *model.GetVoteResultRequest) (*model.GetVoteResultResponse, error) {
	ret := _m.Called(req)

	var r0 *model.GetVoteResultResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.GetVoteResultRequest) (*model.GetVoteResultResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.GetVoteResultRequest) *model.GetVoteResultResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GetVoteResultResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.GetVoteResultRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetVoteResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVoteResult'
type Repository_GetVoteResult_Call struct {
	*mock.Call
}

// GetVoteResult is a helper method to define mock.On call
//   - req *model.GetVoteResultRequest
func (_e *Repository_Expecter) GetVoteResult(req interface{}) *Repository_GetVoteResult_Call {
	return &Repository_GetVoteResult_Call{Call: _e.mock.On("GetVoteResult", req)}
}

func (_c *Repository_GetVoteResult_Call) Run(run func(req *model.GetVoteResultRequest)) *Repository_GetVoteResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.GetVoteResultRequest))
	})
	return _c
}

func (_c *Repository_GetVoteResult_Call) Return(_a0 *model.GetVoteResultResponse, _a1 error) *Repository_GetVoteResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetVoteResult_Call) RunAndReturn(run func(*model.GetVoteResultRequest) (*model.GetVoteResultResponse, error)) *Repository_GetVoteResult_Call {
	_c.Call.Return(run)
	return _c
}

// ListVoteItem provides a mock function with given fields: req
func (_m *Repository) ListVoteItem(req *model.ListVoteItemRequest) (*model.ListVoteItemResponse, error) {
	ret := _m.Called(req)

	var r0 *model.ListVoteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ListVoteItemRequest) (*model.ListVoteItemResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.ListVoteItemRequest) *model.ListVoteItemResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ListVoteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.ListVoteItemRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_ListVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListVoteItem'
type Repository_ListVoteItem_Call struct {
	*mock.Call
}

// ListVoteItem is a helper method to define mock.On call
//   - req *model.ListVoteItemRequest
func (_e *Repository_Expecter) ListVoteItem(req interface{}) *Repository_ListVoteItem_Call {
	return &Repository_ListVoteItem_Call{Call: _e.mock.On("ListVoteItem", req)}
}

func (_c *Repository_ListVoteItem_Call) Run(run func(req *model.ListVoteItemRequest)) *Repository_ListVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.ListVoteItemRequest))
	})
	return _c
}

func (_c *Repository_ListVoteItem_Call) Return(_a0 *model.ListVoteItemResponse, _a1 error) *Repository_ListVoteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_ListVoteItem_Call) RunAndReturn(run func(*model.ListVoteItemRequest) (*model.ListVoteItemResponse, error)) *Repository_ListVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// PrepareUserData provides a mock function with given fields:
func (_m *Repository) PrepareUserData() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_PrepareUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrepareUserData'
type Repository_PrepareUserData_Call struct {
	*mock.Call
}

// PrepareUserData is a helper method to define mock.On call
func (_e *Repository_Expecter) PrepareUserData() *Repository_PrepareUserData_Call {
	return &Repository_PrepareUserData_Call{Call: _e.mock.On("PrepareUserData")}
}

func (_c *Repository_PrepareUserData_Call) Run(run func()) *Repository_PrepareUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_PrepareUserData_Call) Return(_a0 error) *Repository_PrepareUserData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_PrepareUserData_Call) RunAndReturn(run func() error) *Repository_PrepareUserData_Call {
	_c.Call.Return(run)
	return _c
}

// PrepareVoteItem provides a mock function with given fields:
func (_m *Repository) PrepareVoteItem() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_PrepareVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrepareVoteItem'
type Repository_PrepareVoteItem_Call struct {
	*mock.Call
}

// PrepareVoteItem is a helper method to define mock.On call
func (_e *Repository_Expecter) PrepareVoteItem() *Repository_PrepareVoteItem_Call {
	return &Repository_PrepareVoteItem_Call{Call: _e.mock.On("PrepareVoteItem")}
}

func (_c *Repository_PrepareVoteItem_Call) Run(run func()) *Repository_PrepareVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_PrepareVoteItem_Call) Return(_a0 error) *Repository_PrepareVoteItem_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_PrepareVoteItem_Call) RunAndReturn(run func() error) *Repository_PrepareVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateVoteItem provides a mock function with given fields: req
func (_m *Repository) UpdateVoteItem(req *model.UpdateVoteItemRequest) (*model.UpdateVoteItemResponse, error) {
	ret := _m.Called(req)

	var r0 *model.UpdateVoteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.UpdateVoteItemRequest) (*model.UpdateVoteItemResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.UpdateVoteItemRequest) *model.UpdateVoteItemResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UpdateVoteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.UpdateVoteItemRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_UpdateVoteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateVoteItem'
type Repository_UpdateVoteItem_Call struct {
	*mock.Call
}

// UpdateVoteItem is a helper method to define mock.On call
//   - req *model.UpdateVoteItemRequest
func (_e *Repository_Expecter) UpdateVoteItem(req interface{}) *Repository_UpdateVoteItem_Call {
	return &Repository_UpdateVoteItem_Call{Call: _e.mock.On("UpdateVoteItem", req)}
}

func (_c *Repository_UpdateVoteItem_Call) Run(run func(req *model.UpdateVoteItemRequest)) *Repository_UpdateVoteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.UpdateVoteItemRequest))
	})
	return _c
}

func (_c *Repository_UpdateVoteItem_Call) Return(_a0 *model.UpdateVoteItemResponse, _a1 error) *Repository_UpdateVoteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_UpdateVoteItem_Call) RunAndReturn(run func(*model.UpdateVoteItemRequest) (*model.UpdateVoteItemResponse, error)) *Repository_UpdateVoteItem_Call {
	_c.Call.Return(run)
	return _c
}

// Vote provides a mock function with given fields: _a0
func (_m *Repository) Vote(_a0 *model.VoteRequest) (*model.VoteResponse, error) {
	ret := _m.Called(_a0)

	var r0 *model.VoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.VoteRequest) (*model.VoteResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.VoteRequest) *model.VoteResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.VoteRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Vote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Vote'
type Repository_Vote_Call struct {
	*mock.Call
}

// Vote is a helper method to define mock.On call
//   - _a0 *model.VoteRequest
func (_e *Repository_Expecter) Vote(_a0 interface{}) *Repository_Vote_Call {
	return &Repository_Vote_Call{Call: _e.mock.On("Vote", _a0)}
}

func (_c *Repository_Vote_Call) Run(run func(_a0 *model.VoteRequest)) *Repository_Vote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.VoteRequest))
	})
	return _c
}

func (_c *Repository_Vote_Call) Return(_a0 *model.VoteResponse, _a1 error) *Repository_Vote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Vote_Call) RunAndReturn(run func(*model.VoteRequest) (*model.VoteResponse, error)) *Repository_Vote_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
