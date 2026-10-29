package theunitundertest

import "testing"

func TestAdd(t *testing.T) {
	// Arrange
	a := 2
	b := 3

	// Act
	result := Add(a, b)

	// Assert
	if result != 5 {
		t.Errorf("expected 5 , result %d", result)
	}
}
