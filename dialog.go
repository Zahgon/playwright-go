package playwright

type dialogImpl struct {
	channelOwner
	page Page
}

func (d *dialogImpl) Type() string { _ = "STUB: not implemented"; return "" }

func (d *dialogImpl) Message() string { _ = "STUB: not implemented"; return "" }

func (d *dialogImpl) DefaultValue() string { _ = "STUB: not implemented"; return "" }

func (d *dialogImpl) Accept(promptTextInput ...string) error { _ = "STUB: not implemented"; return nil }

func (d *dialogImpl) Dismiss() error { _ = "STUB: not implemented"; return nil }

func (d *dialogImpl) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func newDialog(parent *channelOwner, objectType string, guid string, initializer map[string]any) *dialogImpl {
	_ = "STUB: not implemented"
	return nil
}
