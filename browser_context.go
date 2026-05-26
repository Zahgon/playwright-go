package playwright

import (
	"sync/atomic"

	"github.com/playwright-community/playwright-go/internal/safe"
)

type browserContextImpl struct {
	channelOwner
	timeoutSettings *timeoutSettings
	closeWasCalled  atomic.Bool
	options         *BrowserNewContextOptions
	pages           []Page
	routes          []*routeHandlerEntry
	webSocketRoutes []*webSocketRouteHandler
	ownedPage       Page
	browser         *browserImpl
	serviceWorkers  []Worker
	backgroundPages []Page
	bindings        *safe.SyncMap[string, BindingCallFunction]
	tracing         *tracingImpl
	request         *apiRequestContextImpl
	harRecorders    map[string]harRecordingMetadata
	closed          chan struct{}
	closeReason     *string
	harRouters      []*harRouter
	clock           Clock
}

func (b *browserContextImpl) Clock() Clock { _ = "STUB: not implemented"; return *new(Clock) }

func (b *browserContextImpl) SetDefaultNavigationTimeout(timeout float64) {
	_ = "STUB: not implemented"
	return
}

func (b *browserContextImpl) setDefaultNavigationTimeoutImpl(timeout *float64) {
	_ = "STUB: not implemented"
	return
}

func (b *browserContextImpl) SetDefaultTimeout(timeout float64) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) setDefaultTimeoutImpl(timeout *float64) {
	_ = "STUB: not implemented"
	return
}

func (b *browserContextImpl) Pages() []Page { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) Browser() Browser { _ = "STUB: not implemented"; return *new(Browser) }

func (b *browserContextImpl) Tracing() Tracing { _ = "STUB: not implemented"; return *new(Tracing) }

func (b *browserContextImpl) NewCDPSession(page any) (CDPSession, error) {
	_ = "STUB: not implemented"
	return *new(CDPSession), nil
}

func (b *browserContextImpl) NewPage() (Page, error) {
	_ = "STUB: not implemented"
	return *new(Page), nil
}

func (b *browserContextImpl) Cookies(urls ...string) ([]Cookie, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *browserContextImpl) AddCookies(cookies []OptionalCookie) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) ClearCookies(options ...BrowserContextClearCookiesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) ClearPermissions() error { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) SetGeolocation(geolocation *Geolocation) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) ResetGeolocation() error { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) SetExtraHTTPHeaders(headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) SetOffline(offline bool) error { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) AddInitScript(script Script) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) ExposeFunction(name string, binding ExposedFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) Route(url any, handler routeHandler, times ...int) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) Unroute(url any, handlers ...routeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) unrouteInternal(removed []*routeHandlerEntry, remaining []*routeHandlerEntry, behavior *UnrouteBehavior) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) UnrouteAll(options ...BrowserContextUnrouteAllOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) disposeHarRouters() { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) Request() APIRequestContext {
	_ = "STUB: not implemented"
	return *new(APIRequestContext)
}

func (b *browserContextImpl) RouteFromHAR(har string, options ...BrowserContextRouteFromHAROptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) WaitForEvent(event string, options ...BrowserContextWaitForEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (b *browserContextImpl) waiterForEvent(event string, options ...BrowserContextWaitForEventOptions) *waiter {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) ExpectConsoleMessage(cb func() error, options ...BrowserContextExpectConsoleMessageOptions) (ConsoleMessage, error) {
	_ = "STUB: not implemented"
	return *new(ConsoleMessage), nil
}

func (b *browserContextImpl) ExpectEvent(event string, cb func() error, options ...BrowserContextExpectEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (b *browserContextImpl) ExpectPage(cb func() error, options ...BrowserContextExpectPageOptions) (Page, error) {
	_ = "STUB: not implemented"
	return *new(Page), nil
}

func (b *browserContextImpl) Close(options ...BrowserContextCloseOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Server side will compress artifact if content is attach or if file is .zip.

type browserContextRecordIntoHarOptions struct {
	Page          Page
	URL           any
	UpdateContent *HarContentPolicy
	UpdateMode    *HarMode
}

func (b *browserContextImpl) recordIntoHar(har string, options ...browserContextRecordIntoHarOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) StorageState(paths ...string) (*StorageState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *browserContextImpl) onBinding(binding *bindingCallImpl) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) onClose() { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) onPage(page Page) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) onRoute(route *routeImpl) { _ = "STUB: not implemented"; return }

// If the page or the context was closed we stall all requests right away.

// If the page is closed or unrouteAll() was called without waiting and interception disabled,
// the method will throw an error - silence it.

func (b *browserContextImpl) updateInterceptionPatterns() error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) pause() <-chan error { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) onBackgroundPage(ev map[string]any) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) onServiceWorker(worker *workerImpl) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) setOptions(options *BrowserNewContextOptions, tracesDir *string) {
	_ = "STUB: not implemented"
	return
}

// initializeHarFromOptions starts HAR recording if RecordHarPath is set in options.
// This must be called after context creation to properly register the HAR recorder on the server.
func (b *browserContextImpl) initializeHarFromOptions() error {
	_ = "STUB: not implemented"
	return nil
}

// Determine default content policy based on file extension

func (b *browserContextImpl) BackgroundPages() []Page { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) ServiceWorkers() []Worker { _ = "STUB: not implemented"; return nil }

func (b *browserContextImpl) OnBackgroundPage(fn func(Page)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnClose(fn func(BrowserContext)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnConsole(fn func(ConsoleMessage)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnDialog(fn func(Dialog)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnPage(fn func(Page)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnRequest(fn func(Request)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnRequestFailed(fn func(Request)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnRequestFinished(fn func(Request)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnResponse(fn func(Response)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) OnWebError(fn func(WebError)) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) RouteWebSocket(url any, handler func(WebSocketRoute)) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) onWebSocketRoute(wr WebSocketRoute) { _ = "STUB: not implemented"; return }

func (b *browserContextImpl) updateWebSocketInterceptionPatterns() error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserContextImpl) effectiveCloseReason() *string { _ = "STUB: not implemented"; return nil }

func newBrowserContext(parent *channelOwner, objectType string, guid string, initializer map[string]any) *browserContextImpl {
	_ = "STUB: not implemented"
	return nil
}

// Register this context with the selectors manager for custom selector engines

// Unregister this context from the selectors manager

// Although we do similar handling on the server side, we still need this logic
// on the client side due to a possible race condition between two async calls:
// a) removing "dialog" listener subscription (client->server)
// b) actual "dialog" event (server->client)
