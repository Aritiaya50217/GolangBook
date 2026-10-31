package format_test

import (
	"fmt"
	"test-driven-development/part01the_big_picture/chapter03mocking_and_assertion_frameworks/02exploring_mocks/format"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	// Arrange
	result := 5.55
	expr := "2+3"

	// Act
	wrappedResult := format.Result(expr, result)

	// Assert
	assert.Contains(t, wrappedResult, expr)
	assert.Contains(t, wrappedResult, fmt.Sprint(result))
}
