package playwright

type consoleMessageImpl struct {
	event  map[string]any
	page   Page
	worker Worker
}

func (c *consoleMessageImpl) Type() string { _ = "STUB: not implemented"; return "" }

func (c *consoleMessageImpl) Text() string { _ = "STUB: not implemented"; return "" }

func (c *consoleMessageImpl) String() string { _ = "STUB: not implemented"; return "" }

func (c *consoleMessageImpl) Args() []JSHandle { _ = "STUB: not implemented"; return nil }

func (c *consoleMessageImpl) Location() *ConsoleMessageLocation {
	_ = "STUB: not implemented"
	return nil
}

func (c *consoleMessageImpl) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (c *consoleMessageImpl) Worker() (Worker, error) {
	_ = "STUB: not implemented"
	return *new(Worker), nil
}

func newConsoleMessage(event map[string]any) *consoleMessageImpl {
	_ = "STUB: not implemented"
	return nil
}
