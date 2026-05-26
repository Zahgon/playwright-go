package playwright

import (
	"errors"
)

var (
	testIdAttributeName    = "data-testid"
	ErrLocatorNotSameFrame = errors.New("inner 'has' or 'hasNot' locator must belong to the same frame")
)

type locatorImpl struct {
	frame       *frameImpl
	selector    string
	options     *LocatorOptions
	err         error
	description *string
}

type LocatorOptions LocatorFilterOptions

func newLocator(frame *frameImpl, selector string, options ...LocatorOptions) *locatorImpl {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) equals(locator Locator) bool { _ = "STUB: not implemented"; return false }

func (l *locatorImpl) Err() error { _ = "STUB: not implemented"; return nil }

func (l *locatorImpl) Describe(description string) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) Description() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (l *locatorImpl) All() ([]Locator, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *locatorImpl) AllInnerTexts() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *locatorImpl) AllTextContents() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *locatorImpl) And(locator Locator) Locator { _ = "STUB: not implemented"; return *new(Locator) }

func (l *locatorImpl) Or(locator Locator) Locator { _ = "STUB: not implemented"; return *new(Locator) }

func (l *locatorImpl) Blur(options ...LocatorBlurOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// default 30s, required in Playwright v1.57+

func (l *locatorImpl) AriaSnapshot(options ...LocatorAriaSnapshotOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *locatorImpl) BoundingBox(options ...LocatorBoundingBoxOptions) (*Rect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *locatorImpl) Check(options ...LocatorCheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) Clear(options ...LocatorClearOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) Click(options ...LocatorClickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) ContentFrame() FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (l *locatorImpl) Count() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *locatorImpl) Dblclick(options ...LocatorDblclickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) DispatchEvent(typ string, eventInit any, options ...LocatorDispatchEventOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) DragTo(target Locator, options ...LocatorDragToOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) ElementHandle(options ...LocatorElementHandleOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (l *locatorImpl) ElementHandles() ([]ElementHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *locatorImpl) Evaluate(expression string, arg any, options ...LocatorEvaluateOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (l *locatorImpl) EvaluateAll(expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (l *locatorImpl) EvaluateHandle(expression string, arg any, options ...LocatorEvaluateHandleOptions) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (l *locatorImpl) Fill(value string, options ...LocatorFillOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) Filter(options ...LocatorFilterOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) First() Locator { _ = "STUB: not implemented"; return *new(Locator) }

func (l *locatorImpl) Focus(options ...LocatorFocusOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) FrameLocator(selector string) FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (l *locatorImpl) GetAttribute(name string, options ...LocatorGetAttributeOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *locatorImpl) GetByAltText(text any, options ...LocatorGetByAltTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) GetByLabel(text any, options ...LocatorGetByLabelOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) GetByPlaceholder(text any, options ...LocatorGetByPlaceholderOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) GetByTestId(testId any) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) GetByText(text any, options ...LocatorGetByTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) GetByTitle(text any, options ...LocatorGetByTitleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) Highlight() error { _ = "STUB: not implemented"; return nil }

func (l *locatorImpl) Hover(options ...LocatorHoverOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) InnerHTML(options ...LocatorInnerHTMLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *locatorImpl) InnerText(options ...LocatorInnerTextOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *locatorImpl) InputValue(options ...LocatorInputValueOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *locatorImpl) IsChecked(options ...LocatorIsCheckedOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *locatorImpl) IsDisabled(options ...LocatorIsDisabledOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *locatorImpl) IsEditable(options ...LocatorIsEditableOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *locatorImpl) IsEnabled(options ...LocatorIsEnabledOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *locatorImpl) IsHidden(options ...LocatorIsHiddenOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *locatorImpl) IsVisible(options ...LocatorIsVisibleOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l *locatorImpl) Last() Locator { _ = "STUB: not implemented"; return *new(Locator) }

func (l *locatorImpl) Locator(selectorOrLocator any, options ...LocatorLocatorOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (l *locatorImpl) Nth(index int) Locator { _ = "STUB: not implemented"; return *new(Locator) }

func (l *locatorImpl) Page() (Page, error) { _ = "STUB: not implemented"; return *new(Page), nil }

func (l *locatorImpl) Press(key string, options ...LocatorPressOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) PressSequentially(text string, options ...LocatorPressSequentiallyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) Screenshot(options ...LocatorScreenshotOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *locatorImpl) ScrollIntoViewIfNeeded(options ...LocatorScrollIntoViewIfNeededOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) SelectOption(values SelectOptionValues, options ...LocatorSelectOptionOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *locatorImpl) SelectText(options ...LocatorSelectTextOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) SetChecked(checked bool, options ...LocatorSetCheckedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) SetInputFiles(files any, options ...LocatorSetInputFilesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) Tap(options ...LocatorTapOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) TextContent(options ...LocatorTextContentOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *locatorImpl) Type(text string, options ...LocatorTypeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) Uncheck(options ...LocatorUncheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) WaitFor(options ...LocatorWaitForOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *locatorImpl) withElement(
	callback func(handle ElementHandle) (any, error),
	options ...FrameWaitForSelectorOptions,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (l *locatorImpl) expect(expression string, options frameExpectOptions) (*frameExpectResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
