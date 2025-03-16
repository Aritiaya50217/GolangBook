package tests

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 5, 8},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d+%d", test.a, test.b), func(t *testing.T) {
			if got := Add(test.a, test.b); got != test.result {
				t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, got, test.result)
			}
		})
	}
}

func TestAdd_ValidInput(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestAdd_NegativeInput(t *testing.T) {
	result := Add(-2, -3)
	if result != -5 {
		t.Errorf("Expected -5, but got %d", result)
	}
}
