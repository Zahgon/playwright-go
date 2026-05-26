package playwright

import (
	"sync"
)

type channelOwner struct {
	sync.RWMutex
	eventEmitter
	objectType                 string
	guid                       string
	channel                    *channel
	objects                    map[string]*channelOwner
	eventToSubscriptionMapping map[string]string
	connection                 *connection
	initializer                map[string]any
	parent                     *channelOwner
	wasCollected               bool
	isInternalType             bool
}

func (c *channelOwner) dispose(reason ...string) {
	_ = "STUB: not implemented"
	// Clean up from parent and connection.
	return
}

// Dispose all children.

func (c *channelOwner) adopt(child *channelOwner) { _ = "STUB: not implemented"; return }

func (c *channelOwner) setEventSubscriptionMapping(mapping map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (c *channelOwner) updateSubscription(event string, enabled bool) {
	_ = "STUB: not implemented"
	return
}

func (c *channelOwner) Once(name string, handler any) { _ = "STUB: not implemented"; return }

func (c *channelOwner) On(name string, handler any) { _ = "STUB: not implemented"; return }

func (c *channelOwner) addEvent(name string, handler any, once bool) {
	_ = "STUB: not implemented"
	return
}

func (c *channelOwner) RemoveListener(name string, handler any) { _ = "STUB: not implemented"; return }

func (c *channelOwner) createChannelOwner(self any, parent *channelOwner, objectType string, guid string, initializer map[string]any) {
	_ = "STUB: not implemented"
	return
}

func (c *channelOwner) markAsInternalType() { _ = "STUB: not implemented"; return }

type rootChannelOwner struct {
	channelOwner
}

func (r *rootChannelOwner) initialize() (*Playwright, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GUIDs are now always eagerly resolved in connection.Dispatch

func newRootChannelOwner(connection *connection) *rootChannelOwner {
	_ = "STUB: not implemented"
	return nil
}
