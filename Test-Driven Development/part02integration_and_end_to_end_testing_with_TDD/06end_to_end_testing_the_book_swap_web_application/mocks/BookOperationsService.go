package mocks

import (
	"test-driven-development/part01the_big_picture/chapter04building_efficient_test_suites/db"

	"github.com/stretchr/testify/mock"
)

type BookOperationsService struct {
	mock.Mock
}

func (m *BookOperationsService) ListByUser(userID string) []db.Book {
	ret := m.Called(userID)

	var books []db.Book

	if rf, ok := ret.Get(0).(func(string) []db.Book); ok {
		books = rf(userID)
	} else {
		if ret.Get(0) != nil {
			books = ret.Get(0).([]db.Book)
		}
	}

	return books
}

type mockConstructorTestingTNewBookOperationsService interface {
	mock.TestingT
	Cleanup(func())
}

func NewBookOperationsService(t mockConstructorTestingTNewBookOperationsService) *BookOperationsService {
	mock := &BookOperationsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
