package playwright

import (
	"sync"
)

type selectorsOwnerImpl struct {
	channelOwner
}

func (s *selectorsOwnerImpl) setTestIdAttributeName(name string) { _ = "STUB: not implemented"; return }

func newSelectorsOwner(parent *channelOwner, objectType string, guid string, initializer map[string]any) *selectorsOwnerImpl {
	_ = "STUB: not implemented"
	return nil
}

type selectorsImpl struct {
	mu            sync.RWMutex // protects registrations slice
	contexts      sync.Map     // map of BrowserContext channels
	registrations []map[string]any
}

func (s *selectorsImpl) Register(name string, script Script, options ...SelectorsRegisterOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Register with all active contexts, ignoring contexts that have been closed

// Continue to next context even if this one failed (e.g., context closed)

func (s *selectorsImpl) SetTestIdAttribute(name string) { _ = "STUB: not implemented"; return }

func (s *selectorsImpl) addChannel(channel *selectorsOwnerImpl) {
	_ = "STUB: not implemented"
	// Legacy support for older Playwright versions with server-side selectors
	return
}

func (s *selectorsImpl) removeChannel(channel *selectorsOwnerImpl) {
	_ = "STUB: not implemented"
	// Legacy support for older Playwright versions with server-side selectors
	return
}

func (s *selectorsImpl) addContext(context *browserContextImpl) { _ = "STUB: not implemented"; return }

func (s *selectorsImpl) removeContext(context *browserContextImpl) {
	_ = "STUB: not implemented"
	return
}

func newSelectorsImpl() *selectorsImpl { _ = "STUB: not implemented"; return nil }
