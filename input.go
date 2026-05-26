package playwright

type mouseImpl struct {
	channel *channel
}

func newMouse(channel *channel) *mouseImpl { _ = "STUB: not implemented"; return nil }

func (m *mouseImpl) Move(x float64, y float64, options ...MouseMoveOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mouseImpl) Down(options ...MouseDownOptions) error { _ = "STUB: not implemented"; return nil }

func (m *mouseImpl) Up(options ...MouseUpOptions) error { _ = "STUB: not implemented"; return nil }

func (m *mouseImpl) Click(x, y float64, options ...MouseClickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mouseImpl) Dblclick(x, y float64, options ...MouseDblclickOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mouseImpl) Wheel(deltaX, deltaY float64) error { _ = "STUB: not implemented"; return nil }

type keyboardImpl struct {
	channel *channel
}

func newKeyboard(channel *channel) *keyboardImpl { _ = "STUB: not implemented"; return nil }

func (m *keyboardImpl) Down(key string) error { _ = "STUB: not implemented"; return nil }

func (m *keyboardImpl) Up(key string) error { _ = "STUB: not implemented"; return nil }

func (m *keyboardImpl) InsertText(text string) error { _ = "STUB: not implemented"; return nil }

func (m *keyboardImpl) Type(text string, options ...KeyboardTypeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *keyboardImpl) Press(key string, options ...KeyboardPressOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type touchscreenImpl struct {
	channel *channel
}

func newTouchscreen(channel *channel) *touchscreenImpl { _ = "STUB: not implemented"; return nil }

func (t *touchscreenImpl) Tap(x int, y int) error { _ = "STUB: not implemented"; return nil }
