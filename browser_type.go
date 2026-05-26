package playwright

type browserTypeImpl struct {
	channelOwner
	playwright *Playwright
}

func (b *browserTypeImpl) Name() string { _ = "STUB: not implemented"; return "" }

func (b *browserTypeImpl) ExecutablePath() string { _ = "STUB: not implemented"; return "" }

func (b *browserTypeImpl) Launch(options ...BrowserTypeLaunchOptions) (Browser, error) {
	_ = "STUB: not implemented"
	return *new(Browser), nil
}

// timeout is required in Playwright v1.57+ protocol

// default 30s

func (b *browserTypeImpl) LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error) {
	_ = "STUB: not implemented"
	return *new(BrowserContext), nil
}

// timeout is required in Playwright v1.57+ protocol

// default 30s

func (b *browserTypeImpl) Connect(wsEndpoint string, options ...BrowserTypeConnectOptions) (Browser, error) {
	_ = "STUB: not implemented"
	return *new(Browser), nil
}

// timeout is required in Playwright v1.57+ protocol

// default no timeout

func (b *browserTypeImpl) ConnectOverCDP(endpointURL string, options ...BrowserTypeConnectOverCDPOptions) (Browser, error) {
	_ = "STUB: not implemented"
	return *new(Browser), nil
}

// timeout is required in Playwright v1.57+ protocol

// default 30s

func (b *browserTypeImpl) didCreateContext(context *browserContextImpl, contextOptions *BrowserNewContextOptions, tracesDir *string) {
	_ = "STUB: not implemented"
	return
}

func (b *browserTypeImpl) didLaunchBrowser(browser *browserImpl) { _ = "STUB: not implemented"; return }

func newBrowserType(parent *channelOwner, objectType string, guid string, initializer map[string]any) *browserTypeImpl {
	_ = "STUB: not implemented"
	return nil
}
