package playwright

type BindingCall interface {
	Call(f BindingCallFunction)
}

type bindingCallImpl struct {
	channelOwner
}

// BindingSource is the value passed to a binding call execution
type BindingSource struct {
	Context BrowserContext
	Page    Page
	Frame   Frame
}

// ExposedFunction represents the func signature of an exposed function
type ExposedFunction = func(args ...any) any

// BindingCallFunction represents the func signature of an exposed binding call func
type BindingCallFunction func(source *BindingSource, args ...any) any

func (b *bindingCallImpl) Call(f BindingCallFunction) { _ = "STUB: not implemented"; return }

func serializeError(err error) map[string]any { _ = "STUB: not implemented"; return nil }

// https://github.com/go-stack/stack/issues/27

func newBindingCall(parent *channelOwner, objectType string, guid string, initializer map[string]any) *bindingCallImpl {
	_ = "STUB: not implemented"
	return nil
}
