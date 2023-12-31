// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/adityatresnobudi/library-api/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// BorrowRecordRepository is an autogenerated mock type for the BorrowRecordRepository type
type BorrowRecordRepository struct {
	mock.Mock
}

// CreateBorrowRecord provides a mock function with given fields: ctx, borrow
func (_m *BorrowRecordRepository) CreateBorrowRecord(ctx context.Context, borrow model.BorrowRecords) (model.BorrowRecords, error) {
	ret := _m.Called(ctx, borrow)

	var r0 model.BorrowRecords
	if rf, ok := ret.Get(0).(func(context.Context, model.BorrowRecords) model.BorrowRecords); ok {
		r0 = rf(ctx, borrow)
	} else {
		r0 = ret.Get(0).(model.BorrowRecords)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.BorrowRecords) error); ok {
		r1 = rf(ctx, borrow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBorrowRecordByID provides a mock function with given fields: ctx, id
func (_m *BorrowRecordRepository) FindBorrowRecordByID(ctx context.Context, id int) (model.BorrowRecords, error) {
	ret := _m.Called(ctx, id)

	var r0 model.BorrowRecords
	if rf, ok := ret.Get(0).(func(context.Context, int) model.BorrowRecords); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.BorrowRecords)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBorrowRecord provides a mock function with given fields: ctx, borrow
func (_m *BorrowRecordRepository) UpdateBorrowRecord(ctx context.Context, borrow model.BorrowRecords) (model.BorrowRecords, error) {
	ret := _m.Called(ctx, borrow)

	var r0 model.BorrowRecords
	if rf, ok := ret.Get(0).(func(context.Context, model.BorrowRecords) model.BorrowRecords); ok {
		r0 = rf(ctx, borrow)
	} else {
		r0 = ret.Get(0).(model.BorrowRecords)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.BorrowRecords) error); ok {
		r1 = rf(ctx, borrow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBorrowRecordRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBorrowRecordRepository creates a new instance of BorrowRecordRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBorrowRecordRepository(t mockConstructorTestingTNewBorrowRecordRepository) *BorrowRecordRepository {
	mock := &BorrowRecordRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
