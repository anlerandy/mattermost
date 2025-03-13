// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// PropertyFieldStore is an autogenerated mock type for the PropertyFieldStore type
type PropertyFieldStore struct {
	mock.Mock
}

// CountForGroup provides a mock function with given fields: groupID, includeDeleted
func (_m *PropertyFieldStore) CountForGroup(groupID string, includeDeleted bool) (int64, error) {
	ret := _m.Called(groupID, includeDeleted)

	if len(ret) == 0 {
		panic("no return value specified for CountForGroup")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (int64, error)); ok {
		return rf(groupID, includeDeleted)
	}
	if rf, ok := ret.Get(0).(func(string, bool) int64); ok {
		r0 = rf(groupID, includeDeleted)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(groupID, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: field
func (_m *PropertyFieldStore) Create(field *model.PropertyField) (*model.PropertyField, error) {
	ret := _m.Called(field)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.PropertyField
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.PropertyField) (*model.PropertyField, error)); ok {
		return rf(field)
	}
	if rf, ok := ret.Get(0).(func(*model.PropertyField) *model.PropertyField); ok {
		r0 = rf(field)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PropertyField)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.PropertyField) error); ok {
		r1 = rf(field)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *PropertyFieldStore) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: groupID, id
func (_m *PropertyFieldStore) Get(groupID string, id string) (*model.PropertyField, error) {
	ret := _m.Called(groupID, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.PropertyField
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.PropertyField, error)); ok {
		return rf(groupID, id)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.PropertyField); ok {
		r0 = rf(groupID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PropertyField)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(groupID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMany provides a mock function with given fields: groupID, ids
func (_m *PropertyFieldStore) GetMany(groupID string, ids []string) ([]*model.PropertyField, error) {
	ret := _m.Called(groupID, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetMany")
	}

	var r0 []*model.PropertyField
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) ([]*model.PropertyField, error)); ok {
		return rf(groupID, ids)
	}
	if rf, ok := ret.Get(0).(func(string, []string) []*model.PropertyField); ok {
		r0 = rf(groupID, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PropertyField)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(groupID, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPropertyFields provides a mock function with given fields: opts
func (_m *PropertyFieldStore) SearchPropertyFields(opts model.PropertyFieldSearchOpts) ([]*model.PropertyField, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for SearchPropertyFields")
	}

	var r0 []*model.PropertyField
	var r1 error
	if rf, ok := ret.Get(0).(func(model.PropertyFieldSearchOpts) ([]*model.PropertyField, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(model.PropertyFieldSearchOpts) []*model.PropertyField); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PropertyField)
		}
	}

	if rf, ok := ret.Get(1).(func(model.PropertyFieldSearchOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: fields
func (_m *PropertyFieldStore) Update(fields []*model.PropertyField) ([]*model.PropertyField, error) {
	ret := _m.Called(fields)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 []*model.PropertyField
	var r1 error
	if rf, ok := ret.Get(0).(func([]*model.PropertyField) ([]*model.PropertyField, error)); ok {
		return rf(fields)
	}
	if rf, ok := ret.Get(0).(func([]*model.PropertyField) []*model.PropertyField); ok {
		r0 = rf(fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PropertyField)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.PropertyField) error); ok {
		r1 = rf(fields)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPropertyFieldStore creates a new instance of PropertyFieldStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPropertyFieldStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *PropertyFieldStore {
	mock := &PropertyFieldStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
