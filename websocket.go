package playwright

type webSocketImpl struct {
	channelOwner
	isClosed bool
	page     *pageImpl
}

func (ws *webSocketImpl) URL() string { _ = "STUB: not implemented"; return "" }

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]any) *webSocketImpl {
	_ = "STUB: not implemented"
	return nil
}

func (ws *webSocketImpl) onFrameSent(opcode float64, data string) {
	_ = "STUB: not implemented"
	return
}

func (ws *webSocketImpl) onFrameReceived(opcode float64, data string) {
	_ = "STUB: not implemented"
	return
}

func (ws *webSocketImpl) ExpectEvent(event string, cb func() error, options ...WebSocketExpectEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (ws *webSocketImpl) WaitForEvent(event string, options ...WebSocketWaitForEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (ws *webSocketImpl) expectEvent(event string, cb func() error, options ...WebSocketExpectEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (ws *webSocketImpl) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (ws *webSocketImpl) OnClose(fn func(WebSocket)) { _ = "STUB: not implemented"; return }

func (ws *webSocketImpl) OnFrameReceived(fn func(payload []byte)) {
	_ = "STUB: not implemented"
	return
}

func (ws *webSocketImpl) OnFrameSent(fn func(payload []byte)) { _ = "STUB: not implemented"; return }

func (ws *webSocketImpl) OnSocketError(fn func(string)) { _ = "STUB: not implemented"; return }
