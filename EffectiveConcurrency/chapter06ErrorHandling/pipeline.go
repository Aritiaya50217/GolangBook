package main

import (
	"errors"
	"fmt"
	"strings"
)

func step1(input string) (string, error) {
	if input == "" {
		return "", errors.New("step1 : input is empty")
	}
	return strings.TrimSpace(input), nil
}

func step2(input string) (string, error) {
	if len(input) < 3 {
		return "", errors.New("step2 : input too short")
	}
	return strings.ToUpper(input), nil
}

func step3(input string) (string, error) {
	if strings.Contains(input, "BAD") {
		return "", errors.New("step3 : input contains banned word")
	}
	return input + " ", nil
}

func pipe(input string, steps ...func(string) (string, error)) (string, error) {
	var err error
	for _, step := range steps {
		input, err = step(input)
		if err != nil {
			return "", err
		}
	}
	return input, nil
}

func main() {
	input := " golang pipeline "
	result, err := pipe(input, step1, step2, step3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Final Result:", result)
}
