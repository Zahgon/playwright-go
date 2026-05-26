package playwright

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type frameImpl struct {
	channelOwner
	detached    bool
	page        *pageImpl
	name        string
	url         string
	parentFrame Frame
	childFrames []Frame
	loadStates  mapset.Set[string]
}

func newFrame(parent *channelOwner, objectType string, guid string, initializer map[string]any) *frameImpl {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (f *frameImpl) Name() string { _ = "STUB: not implemented"; return "" }

func (f *frameImpl) SetContent(content string, options ...FrameSetContentOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// timeout is required in Playwright v1.57+ protocol

func (f *frameImpl) Content() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *frameImpl) Goto(url string, options ...FrameGotoOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

// timeout is required in Playwright v1.57+ protocol

// navigation to about:blank or navigation to the same URL with a different hash

func (f *frameImpl) AddScriptTag(options FrameAddScriptTagOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (f *frameImpl) AddStyleTag(options FrameAddStyleTagOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (f *frameImpl) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (f *frameImpl) WaitForLoadState(options ...FrameWaitForLoadStateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) waitForLoadStateImpl(state string, timeout *float64, cb func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) WaitForURL(url any, options ...FrameWaitForURLOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) ExpectNavigation(cb func() error, options ...FrameExpectNavigationOptions) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

// Any failed navigation results in a rejection.

func (f *frameImpl) setNavigationWaiter(timeout *float64) (*waiter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *frameImpl) onFrameNavigated(ev map[string]any) { _ = "STUB: not implemented"; return }

func (f *frameImpl) onLoadState(ev map[string]any) { _ = "STUB: not implemented"; return }

func (f *frameImpl) QuerySelector(selector string, options ...FrameQuerySelectorOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (f *frameImpl) QuerySelectorAll(selector string) ([]ElementHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *frameImpl) Evaluate(expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *frameImpl) EvalOnSelector(selector string, expression string, arg any, options ...FrameEvalOnSelectorOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *frameImpl) EvalOnSelectorAll(selector string, expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (f *frameImpl) EvaluateHandle(expression string, options ...any) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (f *frameImpl) Click(selector string, options ...FrameClickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) WaitForSelector(selector string, options ...FrameWaitForSelectorOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (f *frameImpl) DispatchEvent(selector, typ string, eventInit any, options ...FrameDispatchEventOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) InnerText(selector string, options ...FrameInnerTextOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *frameImpl) InnerHTML(selector string, options ...FrameInnerHTMLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *frameImpl) GetAttribute(selector string, name string, options ...FrameGetAttributeOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *frameImpl) Hover(selector string, options ...FrameHoverOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) SetInputFiles(selector string, files any, options ...FrameSetInputFilesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Type(selector, text string, options ...FrameTypeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Press(selector, key string, options ...FramePressOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Check(selector string, options ...FrameCheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Uncheck(selector string, options ...FrameUncheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) WaitForTimeout(timeout float64) { _ = "STUB: not implemented"; return }

func (f *frameImpl) WaitForFunction(expression string, arg any, options ...FrameWaitForFunctionOptions) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

// timeout is required in Playwright v1.57+ protocol

func (f *frameImpl) Title() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *frameImpl) ChildFrames() []Frame { _ = "STUB: not implemented"; return nil }

func (f *frameImpl) Dblclick(selector string, options ...FrameDblclickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Fill(selector string, value string, options ...FrameFillOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Focus(selector string, options ...FrameFocusOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) FrameElement() (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (f *frameImpl) IsDetached() bool { _ = "STUB: not implemented"; return false }

func (f *frameImpl) ParentFrame() Frame { _ = "STUB: not implemented"; return *new(Frame) }

func (f *frameImpl) TextContent(selector string, options ...FrameTextContentOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *frameImpl) Tap(selector string, options ...FrameTapOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *frameImpl) IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *frameImpl) IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *frameImpl) IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *frameImpl) IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *frameImpl) IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *frameImpl) IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *frameImpl) InputValue(selector string, options ...FrameInputValueOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *frameImpl) DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameImpl) Locator(selector string, options ...FrameLocatorOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByAltText(text any, options ...FrameGetByAltTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByLabel(text any, options ...FrameGetByLabelOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByPlaceholder(text any, options ...FrameGetByPlaceholderOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByRole(role AriaRole, options ...FrameGetByRoleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByTestId(testId any) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByText(text any, options ...FrameGetByTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) GetByTitle(text any, options ...FrameGetByTitleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (f *frameImpl) FrameLocator(selector string) FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (f *frameImpl) highlight(selector string) error { _ = "STUB: not implemented"; return nil }

func (f *frameImpl) queryCount(selector string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
