package deferredfunctions

import (
	"log"
	"testing"
)

type Engine struct{}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

func TestAdd(t *testing.T) {

	defer func() {
		log.Println("Deferred tearing down. ")
	}()

	// Arrange
	e := Engine{}
	x, y := 2.5, 3.5
	want := 6.0

	// Act
	got := e.Add(x, y)

	// Assert
	if got != want {
		t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}

}
