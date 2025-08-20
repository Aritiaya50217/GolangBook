package main

import "fmt"

type customError struct {
	Message string
}

func (e *customError) Error() string {
	return e.Message
}

func testFunction() error {
	if true != false {
		return &customError{Message: "Something went wrong."}
	}
	return nil
}

func main() {
	err := testFunction()
	if err != nil {
		fmt.Println(err)
	}
}
