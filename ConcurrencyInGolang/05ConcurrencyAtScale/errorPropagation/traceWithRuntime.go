package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func generateError() error {
	return errors.New("an unexpected error occurred")
}

func wrappedError() error {
	if err := generateError(); err != nil {
		return fmt.Errorf("wrappedError: %w\n%s", err, debug.Stack())
	}
	return nil
}

func main() {
	if err := wrappedError(); err != nil {
		fmt.Println("Error Trace : ")
		fmt.Println(err)
	}
}
