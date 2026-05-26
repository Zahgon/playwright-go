package playwright

import (
	"sync"
	"sync/atomic"
)

type (
	waiter struct {
		mu        sync.Mutex
		timeout   float64
		fulfilled atomic.Bool
		listeners []eventListener
		errChan   chan error
		waitFunc  func() (any, error)
	}
	eventListener struct {
		emitter EventEmitter
		event   string
		handler any
	}
)

// RejectOnEvent sets the Waiter to return an error when an event occurs (and the predicate returns true)
func (w *waiter) RejectOnEvent(emitter EventEmitter, event string, err error, predicates ...any) *waiter {
	_ = "STUB: not implemented"
	return nil
}

// WithTimeout sets timeout, in milliseconds, for the waiter. 0 means no timeout.
func (w *waiter) WithTimeout(timeout float64) *waiter { _ = "STUB: not implemented"; return nil }

// WaitForEvent sets the Waiter to return when an event occurs (and the predicate returns true)
func (w *waiter) WaitForEvent(emitter EventEmitter, event string, predicate any) *waiter {
	_ = "STUB: not implemented"
	return nil
}

// Wait waits for the waiter to return. It needs to call WaitForEvent once first.
func (w *waiter) Wait() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// RunAndWait waits for the waiter to return after calls func.
func (w *waiter) RunAndWait(cb func() error) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (w *waiter) createHandler(evChan chan<- any, predicate any) func(...any) {
	_ = "STUB: not implemented"
	return nil
}

func (w *waiter) reject(err error) { _ = "STUB: not implemented"; return }

func newWaiter() *waiter {
	_ = "STUB: not implemented"

	// receive both event timeout err and callback err
	// but just return event timeout err
	return nil
}
