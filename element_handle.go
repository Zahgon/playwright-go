package playwright

type elementHandleImpl struct {
	jsHandleImpl
}

func (e *elementHandleImpl) AsElement() ElementHandle {
	_ = "STUB: not implemented"
	return *new(ElementHandle)
}

func (e *elementHandleImpl) OwnerFrame() (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

func (e *elementHandleImpl) ContentFrame() (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

func (e *elementHandleImpl) GetAttribute(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *elementHandleImpl) TextContent() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *elementHandleImpl) InnerText() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *elementHandleImpl) InnerHTML() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *elementHandleImpl) DispatchEvent(typ string, initObjects ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Hover(options ...ElementHandleHoverOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Click(options ...ElementHandleClickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Dblclick(options ...ElementHandleDblclickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) QuerySelector(selector string) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (e *elementHandleImpl) QuerySelectorAll(selector string) ([]ElementHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *elementHandleImpl) EvalOnSelector(selector string, expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (e *elementHandleImpl) EvalOnSelectorAll(selector string, expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (e *elementHandleImpl) ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) SetInputFiles(files any, options ...ElementHandleSetInputFilesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) BoundingBox() (*Rect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *elementHandleImpl) Check(options ...ElementHandleCheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Uncheck(options ...ElementHandleUncheckOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Press(key string, options ...ElementHandlePressOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Fill(value string, options ...ElementHandleFillOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Type(value string, options ...ElementHandleTypeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Focus() error { _ = "STUB: not implemented"; return nil }

func (e *elementHandleImpl) SelectText(options ...ElementHandleSelectTextOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ErrLocatorNotSameFrame

func (e *elementHandleImpl) Tap(options ...ElementHandleTapOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) SelectOption(values SelectOptionValues, options ...ElementHandleSelectOptionOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *elementHandleImpl) IsChecked() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (e *elementHandleImpl) IsDisabled() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *elementHandleImpl) IsEditable() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *elementHandleImpl) IsEnabled() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (e *elementHandleImpl) IsHidden() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (e *elementHandleImpl) IsVisible() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (e *elementHandleImpl) WaitForElementState(state ElementState, options ...ElementHandleWaitForElementStateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *elementHandleImpl) WaitForSelector(selector string, options ...ElementHandleWaitForSelectorOptions) (ElementHandle, error) {
	_ = "STUB: not implemented"
	return *new(ElementHandle), nil
}

func (e *elementHandleImpl) InputValue(options ...ElementHandleInputValueOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *elementHandleImpl) SetChecked(checked bool, options ...ElementHandleSetCheckedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func newElementHandle(parent *channelOwner, objectType string, guid string, initializer map[string]any) *elementHandleImpl {
	_ = "STUB: not implemented"
	return nil
}

func transformToStringList(in any) []string { _ = "STUB: not implemented"; return nil }
