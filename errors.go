package playwright

import (
	"errors"
)

var (
	// ErrPlaywright wraps all Playwright errors.
	//   - Use errors.Is to check if the error is a Playwright error.
	//   - Use errors.As to cast an error to [Error] if you want to access "Stack".
	ErrPlaywright = errors.New("playwright")
	// ErrTargetClosed usually wraps a reason.
	ErrTargetClosed = errors.New("target closed")
	// ErrTimeout wraps timeout errors. It can be either Playwright TimeoutError or client timeout.
	ErrTimeout = errors.New("timeout")
)

// Error represents a Playwright error
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }

// same name and not normal error

func parseError(err Error) error { _ = "STUB: not implemented"; return nil }

func targetClosedError(reason *string) error { _ = "STUB: not implemented"; return nil }
