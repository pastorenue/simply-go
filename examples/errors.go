package main

import (
	"fmt"
	"unicode"
)

func getError(a int) (int, error) {
	if a <= 0 {
		return -1, fmt.Errorf("there was an error with the value: %v", a)
	}
	return a, nil
}

// Custom errors
type notValidNumberError struct {
	errorCode string
	msg string
}

func (e *notValidNumberError) Error() string {
	return fmt.Sprintf("Error code: %v, Message: %v", e.errorCode, e.msg)
}

func getCustomError(a int) (int, error) {
	if !unicode.IsNumber(rune(a)) {
		return -1, &notValidNumberError{errorCode: "1001", msg: "Invalid number"}
	}
	return a, nil
}