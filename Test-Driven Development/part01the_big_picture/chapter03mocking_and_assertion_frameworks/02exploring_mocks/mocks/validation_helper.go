package mocks

import "github.com/stretchr/testify/mock"

type ValidationHelper struct {
	mock.Mock
}

func (m *ValidationHelper) CheckInput(operator string, operands []float64) (err error) {
	ret := m.Called(operator, operands)

	if rf, ok := ret.Get(0).(func(string, []float64) error); ok {
		err = rf(operator, operands)
	} else {
		err = ret.Error(0)
	}

	return err
}

type mockConstructorTestingTNewValidationHelper interface {
	mock.TestingT
	Cleanup(func())
}

func NewValidationHelper(t mockConstructorTestingTNewValidationHelper) *ValidationHelper {
	mock := &ValidationHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}
