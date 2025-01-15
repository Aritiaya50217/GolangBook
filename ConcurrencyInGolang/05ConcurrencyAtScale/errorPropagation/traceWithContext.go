package main

import (
	"errors"
	"fmt"
)

func level1() error {
	return errors.New("original error at level1")
}

func level2() error {
	if err := level1(); err != nil {
		return fmt.Errorf("level2 failed: %w", err)
	}
	return nil
}

func level3() error {
	if err := level2(); err != nil {
		return fmt.Errorf("level3 failed: %w", err)
	}
	return nil

}

func traceWithContext() {
	if err := level3(); err != nil {
		fmt.Println("Error Trace:", err)
	}
}

func main() {
	traceWithContext()
}
