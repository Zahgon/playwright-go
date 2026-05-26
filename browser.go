package playwright

type browserImpl struct {
	channelOwner
	isConnected                  bool
	shouldCloseConnectionOnClose bool
	contexts                     []BrowserContext
	browserType                  BrowserType
	chromiumTracingPath          *string
	closeReason                  *string
}

func (b *browserImpl) BrowserType() BrowserType {
	_ = "STUB: not implemented"
	return *new(BrowserType)
}

func (b *browserImpl) IsConnected() bool { _ = "STUB: not implemented"; return false }

func (b *browserImpl) NewContext(options ...BrowserNewContextOptions) (BrowserContext, error) {
	_ = "STUB: not implemented"
	return *new(BrowserContext), nil
}

func (b *browserImpl) NewPage(options ...BrowserNewPageOptions) (Page, error) {
	_ = "STUB: not implemented"
	return *new(Page), nil
}

func (b *browserImpl) NewBrowserCDPSession() (CDPSession, error) {
	_ = "STUB: not implemented"
	return *new(CDPSession), nil
}

func (b *browserImpl) Contexts() []BrowserContext { _ = "STUB: not implemented"; return nil }

func (b *browserImpl) Close(options ...BrowserCloseOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserImpl) Version() string { _ = "STUB: not implemented"; return "" }

func (b *browserImpl) StartTracing(options ...BrowserStartTracingOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *browserImpl) StopTracing() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *browserImpl) onClose() { _ = "STUB: not implemented"; return }

func (b *browserImpl) OnDisconnected(fn func(Browser)) { _ = "STUB: not implemented"; return }

func newBrowser(parent *channelOwner, objectType string, guid string, initializer map[string]any) *browserImpl {
	_ = "STUB: not implemented"
	return nil
}

// convert parent to *browserTypeImpl

func transformClientCertificate(clientCertificates []ClientCertificate) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
