package playwright

import (
	"sync"
)

type EventEmitter interface {
	Emit(name string, payload ...any) bool
	ListenerCount(name string) int
	On(name string, handler any)
	Once(name string, handler any)
	RemoveListener(name string, handler any)
	RemoveListeners(name string)
}

type (
	eventEmitter struct {
		eventsMutex sync.Mutex
		events      map[string]*eventRegister
		hasInit     bool
	}
	eventRegister struct {
		sync.Mutex
		listeners []listener
	}
	listener struct {
		handler any
		once    bool
	}
)

func NewEventEmitter() EventEmitter { _ = "STUB: not implemented"; return *new(EventEmitter) }

func (e *eventEmitter) Emit(name string, payload ...any) (hasListener bool) {
	_ = "STUB: not implemented"
	return false
}

func (e *eventEmitter) Once(name string, handler any) { _ = "STUB: not implemented"; return }

func (e *eventEmitter) On(name string, handler any) { _ = "STUB: not implemented"; return }

func (e *eventEmitter) RemoveListener(name string, handler any) { _ = "STUB: not implemented"; return }

func (e *eventEmitter) RemoveListeners(name string) { _ = "STUB: not implemented"; return }

// ListenerCount count the listeners by name, count all if name is empty
func (e *eventEmitter) ListenerCount(name string) int { _ = "STUB: not implemented"; return 0 }

func (e *eventEmitter) addEvent(name string, handler any, once bool) {
	_ = "STUB: not implemented"
	return
}

func (e *eventEmitter) init() {
	if !e.hasInit {
		e.events = make(map[string]*eventRegister, 0)
		e.hasInit = true
	}
}

func (er *eventRegister) addHandler(handler any, once bool) { _ = "STUB: not implemented"; return }

func (er *eventRegister) count() int { _ = "STUB: not implemented"; return 0 }

func (er *eventRegister) removeHandler(handler any) { _ = "STUB: not implemented"; return }

func (er *eventRegister) callHandlers(payloads ...any) int { _ = "STUB: not implemented"; return 0 }
