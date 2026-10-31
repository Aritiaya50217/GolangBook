package format_test

import (
	"errors"
	"test-driven-development/part01the_big_picture/chapter03mocking_and_assertion_frameworks/02exploring_mocks/format"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	// Arrange
	initialErr := errors.New("error msg")
	expr := "2%3"

	// Act
	wrappedErr := format.Error(expr, initialErr)

	// Assert
	assert.Contains(t, wrappedErr.Error(), initialErr.Error())
	assert.Contains(t, wrappedErr.Error(), expr)
}
