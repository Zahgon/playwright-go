package playwright

import (
	"sync/atomic"
)

type webSocketRouteImpl struct {
	channelOwner
	connected       *atomic.Bool
	server          WebSocketRoute
	onPageMessage   func(any)
	onPageClose     func(code *int, reason *string)
	onServerMessage func(any)
	onServerClose   func(code *int, reason *string)
}

func newWebSocketRoute(parent *channelOwner, objectType string, guid string, initializer map[string]any) *webSocketRouteImpl {
	_ = "STUB: not implemented"
	return nil
}

func (r *webSocketRouteImpl) Close(options ...WebSocketRouteCloseOptions) {
	_ = "STUB: not implemented"
	return
}

func (r *webSocketRouteImpl) ConnectToServer() (WebSocketRoute, error) {
	_ = "STUB: not implemented"
	return *new(WebSocketRoute), nil
}

func (r *webSocketRouteImpl) OnClose(handler func(code *int, reason *string)) {
	_ = "STUB: not implemented"
	return
}

func (r *webSocketRouteImpl) OnMessage(handler func(any)) { _ = "STUB: not implemented"; return }

func (r *webSocketRouteImpl) Send(message any) { _ = "STUB: not implemented"; return }

func (r *webSocketRouteImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (r *webSocketRouteImpl) afterHandle() error { _ = "STUB: not implemented"; return nil }

// Ensure that websocket is "open" and can send messages without an actual server connection.

type serverWebSocketRouteImpl struct {
	webSocketRoute *webSocketRouteImpl
}

func newServerWebSocketRoute(route *webSocketRouteImpl) *serverWebSocketRouteImpl {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverWebSocketRouteImpl) OnMessage(handler func(any)) { _ = "STUB: not implemented"; return }

func (s *serverWebSocketRouteImpl) OnClose(handler func(code *int, reason *string)) {
	_ = "STUB: not implemented"
	return
}

func (s *serverWebSocketRouteImpl) ConnectToServer() (WebSocketRoute, error) {
	_ = "STUB: not implemented"
	return *new(WebSocketRoute), nil
}

func (s *serverWebSocketRouteImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (s *serverWebSocketRouteImpl) Close(options ...WebSocketRouteCloseOptions) {
	_ = "STUB: not implemented"
	return
}

func (s *serverWebSocketRouteImpl) Send(message any) { _ = "STUB: not implemented"; return }

func transformWebSocketMessage(message any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func untransformWebSocketMessage(data map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type webSocketRouteHandler struct {
	matcher *urlMatcher
	handler func(WebSocketRoute)
}

func newWebSocketRouteHandler(matcher *urlMatcher, handler func(WebSocketRoute)) *webSocketRouteHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *webSocketRouteHandler) Handle(route WebSocketRoute) { _ = "STUB: not implemented"; return }

func (h *webSocketRouteHandler) Matches(wsURL string) bool { _ = "STUB: not implemented"; return false }

func prepareWebSocketRouteHandlerInterceptionPatterns(handlers []*webSocketRouteHandler) []map[string]any {
	_ = "STUB: not implemented"
	return nil
}
