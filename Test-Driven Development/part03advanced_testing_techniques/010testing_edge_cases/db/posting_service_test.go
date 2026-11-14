package db_test

import (
	"test-driven-development/part02integration_and_end_to_end_testing_with_TDD/07refactoring_in_go/db"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	ps := db.NewPostingService()
	b := db.Book{
		ID: uuid.New().String(),
	}
	err := ps.NewOrder(b)
	assert.Nil(t, err)
}
