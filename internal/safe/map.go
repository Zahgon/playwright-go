package safe

import (
	"sync"
)

// SyncMap is a thread-safe map
type SyncMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

// NewSyncMap creates a new thread-safe map
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] { _ = "STUB: not implemented"; return nil }

func (m *SyncMap[K, V]) Store(k K, v V) { _ = "STUB: not implemented"; return }

func (m *SyncMap[K, V]) Load(k K) (v V, ok bool) { _ = "STUB: not implemented"; return *new(V), false }

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m *SyncMap[K, V]) LoadOrStore(k K, v V) (actual V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// LoadAndDelete deletes the value for a key, and returns the previous value if any.
func (m *SyncMap[K, V]) LoadAndDelete(k K) (v V, loaded bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *SyncMap[K, V]) Delete(k K) { _ = "STUB: not implemented"; return }

func (m *SyncMap[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (m *SyncMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *SyncMap[K, V]) Clone() map[K]V { _ = "STUB: not implemented"; return nil }

func (m *SyncMap[K, V]) Range(f func(k K, v V) bool) { _ = "STUB: not implemented"; return }
