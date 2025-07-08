package codecoverage

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := Subtract(3, 3)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
