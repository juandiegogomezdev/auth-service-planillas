package apperrors

import "fmt"

// This error will be used by all auth services
type SError struct {
	Code string
	Err  error
}

func (e *SError) Error() string {
	return fmt.Sprintf("%s: %v", e.Code, e.Err)
}

func WrapSerror(code string, err error) *SError {
	return &SError{
		Code: code,
		Err:  err,
	}
}
