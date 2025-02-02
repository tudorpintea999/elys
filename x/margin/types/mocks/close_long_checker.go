// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	ammtypes "github.com/elys-network/elys/x/amm/types"
	margintypes "github.com/elys-network/elys/x/margin/types"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CloseLongChecker is an autogenerated mock type for the CloseLongChecker type
type CloseLongChecker struct {
	mock.Mock
}

type CloseLongChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *CloseLongChecker) EXPECT() *CloseLongChecker_Expecter {
	return &CloseLongChecker_Expecter{mock: &_m.Mock}
}

// EstimateAndRepay provides a mock function with given fields: ctx, mtp, pool, ammPool, collateralAsset, custodyAsset
func (_m *CloseLongChecker) EstimateAndRepay(ctx types.Context, mtp margintypes.MTP, pool margintypes.Pool, ammPool ammtypes.Pool, collateralAsset string, custodyAsset string) (math.Int, error) {
	ret := _m.Called(ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, margintypes.MTP, margintypes.Pool, ammtypes.Pool, string, string) (math.Int, error)); ok {
		return rf(ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)
	}
	if rf, ok := ret.Get(0).(func(types.Context, margintypes.MTP, margintypes.Pool, ammtypes.Pool, string, string) math.Int); ok {
		r0 = rf(ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(types.Context, margintypes.MTP, margintypes.Pool, ammtypes.Pool, string, string) error); ok {
		r1 = rf(ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseLongChecker_EstimateAndRepay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateAndRepay'
type CloseLongChecker_EstimateAndRepay_Call struct {
	*mock.Call
}

// EstimateAndRepay is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp margintypes.MTP
//   - pool margintypes.Pool
//   - ammPool ammtypes.Pool
//   - collateralAsset string
//   - custodyAsset string
func (_e *CloseLongChecker_Expecter) EstimateAndRepay(ctx interface{}, mtp interface{}, pool interface{}, ammPool interface{}, collateralAsset interface{}, custodyAsset interface{}) *CloseLongChecker_EstimateAndRepay_Call {
	return &CloseLongChecker_EstimateAndRepay_Call{Call: _e.mock.On("EstimateAndRepay", ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)}
}

func (_c *CloseLongChecker_EstimateAndRepay_Call) Run(run func(ctx types.Context, mtp margintypes.MTP, pool margintypes.Pool, ammPool ammtypes.Pool, collateralAsset string, custodyAsset string)) *CloseLongChecker_EstimateAndRepay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(margintypes.MTP), args[2].(margintypes.Pool), args[3].(ammtypes.Pool), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *CloseLongChecker_EstimateAndRepay_Call) Return(_a0 math.Int, _a1 error) *CloseLongChecker_EstimateAndRepay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloseLongChecker_EstimateAndRepay_Call) RunAndReturn(run func(types.Context, margintypes.MTP, margintypes.Pool, ammtypes.Pool, string, string) (math.Int, error)) *CloseLongChecker_EstimateAndRepay_Call {
	_c.Call.Return(run)
	return _c
}

// GetAmmPool provides a mock function with given fields: ctx, poolId, tradingAsset
func (_m *CloseLongChecker) GetAmmPool(ctx types.Context, poolId uint64, tradingAsset string) (ammtypes.Pool, error) {
	ret := _m.Called(ctx, poolId, tradingAsset)

	var r0 ammtypes.Pool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) (ammtypes.Pool, error)); ok {
		return rf(ctx, poolId, tradingAsset)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) ammtypes.Pool); ok {
		r0 = rf(ctx, poolId, tradingAsset)
	} else {
		r0 = ret.Get(0).(ammtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64, string) error); ok {
		r1 = rf(ctx, poolId, tradingAsset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseLongChecker_GetAmmPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAmmPool'
type CloseLongChecker_GetAmmPool_Call struct {
	*mock.Call
}

// GetAmmPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
//   - tradingAsset string
func (_e *CloseLongChecker_Expecter) GetAmmPool(ctx interface{}, poolId interface{}, tradingAsset interface{}) *CloseLongChecker_GetAmmPool_Call {
	return &CloseLongChecker_GetAmmPool_Call{Call: _e.mock.On("GetAmmPool", ctx, poolId, tradingAsset)}
}

func (_c *CloseLongChecker_GetAmmPool_Call) Run(run func(ctx types.Context, poolId uint64, tradingAsset string)) *CloseLongChecker_GetAmmPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64), args[2].(string))
	})
	return _c
}

func (_c *CloseLongChecker_GetAmmPool_Call) Return(_a0 ammtypes.Pool, _a1 error) *CloseLongChecker_GetAmmPool_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloseLongChecker_GetAmmPool_Call) RunAndReturn(run func(types.Context, uint64, string) (ammtypes.Pool, error)) *CloseLongChecker_GetAmmPool_Call {
	_c.Call.Return(run)
	return _c
}

// GetMTP provides a mock function with given fields: ctx, mtpAddress, id
func (_m *CloseLongChecker) GetMTP(ctx types.Context, mtpAddress string, id uint64) (margintypes.MTP, error) {
	ret := _m.Called(ctx, mtpAddress, id)

	var r0 margintypes.MTP
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, string, uint64) (margintypes.MTP, error)); ok {
		return rf(ctx, mtpAddress, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string, uint64) margintypes.MTP); ok {
		r0 = rf(ctx, mtpAddress, id)
	} else {
		r0 = ret.Get(0).(margintypes.MTP)
	}

	if rf, ok := ret.Get(1).(func(types.Context, string, uint64) error); ok {
		r1 = rf(ctx, mtpAddress, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseLongChecker_GetMTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMTP'
type CloseLongChecker_GetMTP_Call struct {
	*mock.Call
}

// GetMTP is a helper method to define mock.On call
//   - ctx types.Context
//   - mtpAddress string
//   - id uint64
func (_e *CloseLongChecker_Expecter) GetMTP(ctx interface{}, mtpAddress interface{}, id interface{}) *CloseLongChecker_GetMTP_Call {
	return &CloseLongChecker_GetMTP_Call{Call: _e.mock.On("GetMTP", ctx, mtpAddress, id)}
}

func (_c *CloseLongChecker_GetMTP_Call) Run(run func(ctx types.Context, mtpAddress string, id uint64)) *CloseLongChecker_GetMTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *CloseLongChecker_GetMTP_Call) Return(_a0 margintypes.MTP, _a1 error) *CloseLongChecker_GetMTP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CloseLongChecker_GetMTP_Call) RunAndReturn(run func(types.Context, string, uint64) (margintypes.MTP, error)) *CloseLongChecker_GetMTP_Call {
	_c.Call.Return(run)
	return _c
}

// GetPool provides a mock function with given fields: ctx, poolId
func (_m *CloseLongChecker) GetPool(ctx types.Context, poolId uint64) (margintypes.Pool, bool) {
	ret := _m.Called(ctx, poolId)

	var r0 margintypes.Pool
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (margintypes.Pool, bool)); ok {
		return rf(ctx, poolId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) margintypes.Pool); ok {
		r0 = rf(ctx, poolId)
	} else {
		r0 = ret.Get(0).(margintypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) bool); ok {
		r1 = rf(ctx, poolId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// CloseLongChecker_GetPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPool'
type CloseLongChecker_GetPool_Call struct {
	*mock.Call
}

// GetPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
func (_e *CloseLongChecker_Expecter) GetPool(ctx interface{}, poolId interface{}) *CloseLongChecker_GetPool_Call {
	return &CloseLongChecker_GetPool_Call{Call: _e.mock.On("GetPool", ctx, poolId)}
}

func (_c *CloseLongChecker_GetPool_Call) Run(run func(ctx types.Context, poolId uint64)) *CloseLongChecker_GetPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *CloseLongChecker_GetPool_Call) Return(val margintypes.Pool, found bool) *CloseLongChecker_GetPool_Call {
	_c.Call.Return(val, found)
	return _c
}

func (_c *CloseLongChecker_GetPool_Call) RunAndReturn(run func(types.Context, uint64) (margintypes.Pool, bool)) *CloseLongChecker_GetPool_Call {
	_c.Call.Return(run)
	return _c
}

// HandleInterest provides a mock function with given fields: ctx, mtp, pool, ammPool, collateralAsset, custodyAsset
func (_m *CloseLongChecker) HandleInterest(ctx types.Context, mtp *margintypes.MTP, pool *margintypes.Pool, ammPool ammtypes.Pool, collateralAsset string, custodyAsset string) error {
	ret := _m.Called(ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *margintypes.MTP, *margintypes.Pool, ammtypes.Pool, string, string) error); ok {
		r0 = rf(ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseLongChecker_HandleInterest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleInterest'
type CloseLongChecker_HandleInterest_Call struct {
	*mock.Call
}

// HandleInterest is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp *margintypes.MTP
//   - pool *margintypes.Pool
//   - ammPool ammtypes.Pool
//   - collateralAsset string
//   - custodyAsset string
func (_e *CloseLongChecker_Expecter) HandleInterest(ctx interface{}, mtp interface{}, pool interface{}, ammPool interface{}, collateralAsset interface{}, custodyAsset interface{}) *CloseLongChecker_HandleInterest_Call {
	return &CloseLongChecker_HandleInterest_Call{Call: _e.mock.On("HandleInterest", ctx, mtp, pool, ammPool, collateralAsset, custodyAsset)}
}

func (_c *CloseLongChecker_HandleInterest_Call) Run(run func(ctx types.Context, mtp *margintypes.MTP, pool *margintypes.Pool, ammPool ammtypes.Pool, collateralAsset string, custodyAsset string)) *CloseLongChecker_HandleInterest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*margintypes.MTP), args[2].(*margintypes.Pool), args[3].(ammtypes.Pool), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *CloseLongChecker_HandleInterest_Call) Return(_a0 error) *CloseLongChecker_HandleInterest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloseLongChecker_HandleInterest_Call) RunAndReturn(run func(types.Context, *margintypes.MTP, *margintypes.Pool, ammtypes.Pool, string, string) error) *CloseLongChecker_HandleInterest_Call {
	_c.Call.Return(run)
	return _c
}

// TakeOutCustody provides a mock function with given fields: ctx, mtp, pool, custodyAsset
func (_m *CloseLongChecker) TakeOutCustody(ctx types.Context, mtp margintypes.MTP, pool *margintypes.Pool, custodyAsset string) error {
	ret := _m.Called(ctx, mtp, pool, custodyAsset)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, margintypes.MTP, *margintypes.Pool, string) error); ok {
		r0 = rf(ctx, mtp, pool, custodyAsset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseLongChecker_TakeOutCustody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TakeOutCustody'
type CloseLongChecker_TakeOutCustody_Call struct {
	*mock.Call
}

// TakeOutCustody is a helper method to define mock.On call
//   - ctx types.Context
//   - mtp margintypes.MTP
//   - pool *margintypes.Pool
//   - custodyAsset string
func (_e *CloseLongChecker_Expecter) TakeOutCustody(ctx interface{}, mtp interface{}, pool interface{}, custodyAsset interface{}) *CloseLongChecker_TakeOutCustody_Call {
	return &CloseLongChecker_TakeOutCustody_Call{Call: _e.mock.On("TakeOutCustody", ctx, mtp, pool, custodyAsset)}
}

func (_c *CloseLongChecker_TakeOutCustody_Call) Run(run func(ctx types.Context, mtp margintypes.MTP, pool *margintypes.Pool, custodyAsset string)) *CloseLongChecker_TakeOutCustody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(margintypes.MTP), args[2].(*margintypes.Pool), args[3].(string))
	})
	return _c
}

func (_c *CloseLongChecker_TakeOutCustody_Call) Return(_a0 error) *CloseLongChecker_TakeOutCustody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CloseLongChecker_TakeOutCustody_Call) RunAndReturn(run func(types.Context, margintypes.MTP, *margintypes.Pool, string) error) *CloseLongChecker_TakeOutCustody_Call {
	_c.Call.Return(run)
	return _c
}

// NewCloseLongChecker creates a new instance of CloseLongChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCloseLongChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *CloseLongChecker {
	mock := &CloseLongChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
