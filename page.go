package playwright

import (
	"sync/atomic"

	"github.com/playwright-community/playwright-go/internal/safe"
)

type pageImpl struct {
	channelOwner
	isClosed        bool
	closedOrCrashed chan error
	video           *videoImpl
	mouse           *mouseImpl
	keyboard        *keyboardImpl
	touchscreen     *touchscreenImpl
	timeoutSettings *timeoutSettings
	browserContext  *browserContextImpl
	frames          []Frame
	workers         []Worker
	mainFrame       Frame
	routes          []*routeHandlerEntry
	webSocketRoutes []*webSocketRouteHandler
	viewportSize    *Size
	ownedContext    BrowserContext
	bindings        *safe.SyncMap[string, BindingCallFunction]
	closeReason     *string
	closeWasCalled  atomic.Bool
	harRouters      []*harRouter
	locatorHandlers map[float64]*locatorHandlerEntry
}

type locatorHandlerEntry struct {
	locator *locatorImpl
	handler func(Locator)
	times   *int
}

func (p *pageImpl) AddLocatorHandler(locator Locator, handler func(Locator), options ...PageAddLocatorHandlerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) onLocatorHandlerTriggered(uid float64) { _ = "STUB: not implemented"; return }

func (p *pageImpl) RemoveLocatorHandler(locator Locator) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Context() BrowserContext { _ = "STUB: not implemented"; return *new(BrowserContext) }

func (b *pageImpl) Clock() Clock { _ = "STUB: not implemented"; return *new(Clock) }

func (p *pageImpl) Close(options ...PageCloseOptions) error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) InnerText(selector string, options ...PageInnerTextOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *pageImpl) InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *pageImpl) Opener() (Page, error) { _ = "STUB: not implemented"; return *new(Page), nil }

// not popup page or opener has been closed

func (p *pageImpl) MainFrame() Frame { _ = "STUB: not implemented"; return *new(Frame) }

func (p *pageImpl) Frame(options ...PageFrameOptions) Frame {
	_ = "STUB: not implemented"
	return *new(Frame)
}

func (p *pageImpl) Frames() []Frame { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) SetDefaultNavigationTimeout(timeout float64) { _ = "STUB: not implemented"; return }

func (p *pageImpl) SetDefaultTimeout(timeout float64) { _ = "STUB: not implemented"; return }

func (p *pageImpl) QuerySelector(selector string, options ...PageQuerySelectorOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (p *pageImpl) QuerySelectorAll(selector string) ([]ElementHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pageImpl) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (p *pageImpl) DispatchEvent(selector string, typ string, eventInit any, options ...PageDispatchEventOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Evaluate(expression string, arg ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pageImpl) EvaluateHandle(expression string, arg ...any) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (p *pageImpl) EvalOnSelector(selector string, expression string, arg any, options ...PageEvalOnSelectorOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pageImpl) EvalOnSelectorAll(selector string, expression string, arg ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pageImpl) AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (p *pageImpl) AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (p *pageImpl) SetExtraHTTPHeaders(headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (p *pageImpl) Unroute(url any, handlers ...routeHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) unrouteInternal(removed []*routeHandlerEntry, remaining []*routeHandlerEntry, behavior *UnrouteBehavior) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) disposeHarRouters() { _ = "STUB: not implemented"; return }

func (p *pageImpl) UnrouteAll(options ...PageUnrouteAllOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Content() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *pageImpl) SetContent(html string, options ...PageSetContentOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Goto(url string, options ...PageGotoOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (p *pageImpl) Reload(options ...PageReloadOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (p *pageImpl) WaitForLoadState(options ...PageWaitForLoadStateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) GoBack(options ...PageGoBackOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

// can not go back

func (p *pageImpl) GoForward(options ...PageGoForwardOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

// can not go forward

func (p *pageImpl) EmulateMedia(options ...PageEmulateMediaOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) SetViewportSize(width, height int) error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) ViewportSize() *Size { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) BringToFront() error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) Type(selector, text string, options ...PageTypeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Fill(selector, text string, options ...PageFillOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Press(selector, key string, options ...PagePressOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Title() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *pageImpl) Workers() []Worker { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) Request() APIRequestContext {
	_ = "STUB: not implemented"
	return *new(APIRequestContext)
}

func (p *pageImpl) Screenshot(options ...PageScreenshotOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ErrLocatorNotSameFrame

func (p *pageImpl) PDF(options ...PagePdfOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pageImpl) Click(selector string, options ...PageClickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) WaitForEvent(event string, options ...PageWaitForEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pageImpl) waiterForEvent(event string, options ...PageWaitForEventOptions) *waiter {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) waiterForRequest(url any, options ...PageExpectRequestOptions) *waiter {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) waiterForResponse(url any, options ...PageExpectResponseOptions) *waiter {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) ExpectEvent(event string, cb func() error, options ...PageExpectEventOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pageImpl) ExpectNavigation(cb func() error, options ...PageExpectNavigationOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (p *pageImpl) ExpectConsoleMessage(cb func() error, options ...PageExpectConsoleMessageOptions) (ConsoleMessage, error) {
	_ = "STUB: not implemented"
	return *new(ConsoleMessage), nil
}

func (p *pageImpl) ExpectDownload(cb func() error, options ...PageExpectDownloadOptions) (Download, error) {
	_ = "STUB: not implemented"
	return *new(Download), nil
}

func (p *pageImpl) ExpectFileChooser(cb func() error, options ...PageExpectFileChooserOptions) (FileChooser, error) {
	_ = "STUB: not implemented"
	return *new(FileChooser), nil
}

func (p *pageImpl) ExpectPopup(cb func() error, options ...PageExpectPopupOptions) (Page, error) {
	_ = "STUB: not implemented"
	return *new(Page), nil
}

func (p *pageImpl) ExpectResponse(url any, cb func() error, options ...PageExpectResponseOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (p *pageImpl) ExpectRequest(url any, cb func() error, options ...PageExpectRequestOptions) (Request, error) {
	_ = "STUB: not implemented"
	return *new(Request), nil
}

func (p *pageImpl) ExpectRequestFinished(cb func() error, options ...PageExpectRequestFinishedOptions) (Request, error) {
	_ = "STUB: not implemented"
	return *new(Request), nil
}

func (p *pageImpl) ExpectWebSocket(cb func() error, options ...PageExpectWebSocketOptions) (WebSocket, error) {
	_ = "STUB: not implemented"
	return *new(WebSocket), nil
}

func (p *pageImpl) ExpectWorker(cb func() error, options ...PageExpectWorkerOptions) (Worker, error) {
	_ = "STUB: not implemented"
	return *new(Worker), nil
}

func (p *pageImpl) Route(url any, handler routeHandler, times ...int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *pageImpl) Hover(selector string, options ...PageHoverOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (p *pageImpl) AddInitScript(script Script) error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) Keyboard() Keyboard { _ = "STUB: not implemented"; return *new(Keyboard) }

func (p *pageImpl) ConsoleMessages() ([]ConsoleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pageImpl) Requests() ([]Request, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pageImpl) Mouse() Mouse { _ = "STUB: not implemented"; return *new(Mouse) }

func (p *pageImpl) RouteFromHAR(har string, options ...PageRouteFromHAROptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Touchscreen() Touchscreen { _ = "STUB: not implemented"; return *new(Touchscreen) }

func newPage(parent *channelOwner, objectType string, guid string, initializer map[string]any) *pageImpl {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) closeErrorWithReason() error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) onBinding(binding *bindingCallImpl) { _ = "STUB: not implemented"; return }

func (p *pageImpl) onFrameAttached(frame *frameImpl) { _ = "STUB: not implemented"; return }

func (p *pageImpl) onFrameDetached(frame *frameImpl) { _ = "STUB: not implemented"; return }

func (p *pageImpl) onRoute(route *routeImpl) { _ = "STUB: not implemented"; return }

// If the page was closed we stall all requests right away.

func (p *pageImpl) updateInterceptionPatterns() error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) onWorker(worker *workerImpl) { _ = "STUB: not implemented"; return }

func (p *pageImpl) onClose() { _ = "STUB: not implemented"; return }

func (p *pageImpl) SetInputFiles(selector string, files any, options ...PageSetInputFilesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Check(selector string, options ...PageCheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Uncheck(selector string, options ...PageUncheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) WaitForTimeout(timeout float64) { _ = "STUB: not implemented"; return }

func (p *pageImpl) WaitForFunction(expression string, arg any, options ...PageWaitForFunctionOptions) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (p *pageImpl) Dblclick(expression string, options ...PageDblclickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Focus(expression string, options ...PageFocusOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) TextContent(selector string, options ...PageTextContentOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *pageImpl) Video() Video { _ = "STUB: not implemented"; return *new(Video) }

func (p *pageImpl) Tap(selector string, options ...PageTapOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) ExposeFunction(name string, binding ExposedFunction) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) SelectOption(selector string, values SelectOptionValues, options ...PageSelectOptionOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pageImpl) IsChecked(selector string, options ...PageIsCheckedOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pageImpl) IsDisabled(selector string, options ...PageIsDisabledOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pageImpl) IsEditable(selector string, options ...PageIsEditableOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pageImpl) IsEnabled(selector string, options ...PageIsEnabledOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pageImpl) IsHidden(selector string, options ...PageIsHiddenOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pageImpl) IsVisible(selector string, options ...PageIsVisibleOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *pageImpl) DragAndDrop(source, target string, options ...PageDragAndDropOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Pause() (err error) { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) InputValue(selector string, options ...PageInputValueOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *pageImpl) WaitForURL(url any, options ...PageWaitForURLOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) SetChecked(selector string, checked bool, options ...PageSetCheckedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) Locator(selector string, options ...PageLocatorOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) GetByAltText(text any, options ...PageGetByAltTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) GetByLabel(text any, options ...PageGetByLabelOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) GetByPlaceholder(text any, options ...PageGetByPlaceholderOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) GetByRole(role AriaRole, options ...PageGetByRoleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) GetByTestId(testId any) Locator { _ = "STUB: not implemented"; return *new(Locator) }

func (p *pageImpl) GetByText(text any, options ...PageGetByTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) GetByTitle(text any, options ...PageGetByTitleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (p *pageImpl) FrameLocator(selector string) FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (p *pageImpl) OnClose(fn func(Page)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnConsole(fn func(ConsoleMessage)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnCrash(fn func(Page)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnDialog(fn func(Dialog)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnDOMContentLoaded(fn func(Page)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnDownload(fn func(Download)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnFileChooser(fn func(FileChooser)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnFrameAttached(fn func(Frame)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnFrameDetached(fn func(Frame)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnFrameNavigated(fn func(Frame)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnLoad(fn func(Page)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnPageError(fn func(error)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnPopup(fn func(Page)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnRequest(fn func(Request)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnRequestFailed(fn func(Request)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnRequestFinished(fn func(Request)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnResponse(fn func(Response)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnWebSocket(fn func(WebSocket)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) OnWorker(fn func(Worker)) { _ = "STUB: not implemented"; return }

func (p *pageImpl) RequestGC() error { _ = "STUB: not implemented"; return nil }

func (p *pageImpl) RouteWebSocket(url any, handler func(WebSocketRoute)) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pageImpl) onWebSocketRoute(wr WebSocketRoute) { _ = "STUB: not implemented"; return }

func (p *pageImpl) updateWebSocketInterceptionPatterns() error {
	_ = "STUB: not implemented"
	return nil
}
