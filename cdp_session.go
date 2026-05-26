package playwright

type cdpSessionImpl struct {
	channelOwner
}

func (c *cdpSessionImpl) Detach() error { _ = "STUB: not implemented"; return nil }

func (c *cdpSessionImpl) Send(method string, params map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *cdpSessionImpl) onEvent(params map[string]any) { _ = "STUB: not implemented"; return }

func newCDPSession(parent *channelOwner, objectType string, guid string, initializer map[string]any) *cdpSessionImpl {
	_ = "STUB: not implemented"
	return nil
}
