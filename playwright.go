// Package playwright is a library to automate Chromium, Firefox and WebKit with
// a single API. Playwright is built to enable cross-browser web automation that
// is ever-green, capable, reliable and fast.
package playwright

// DeviceDescriptor represents a single device
type DeviceDescriptor struct {
	UserAgent          string  `json:"userAgent"`
	Viewport           *Size   `json:"viewport"`
	Screen             *Size   `json:"screen"`
	DeviceScaleFactor  float64 `json:"deviceScaleFactor"`
	IsMobile           bool    `json:"isMobile"`
	HasTouch           bool    `json:"hasTouch"`
	DefaultBrowserType string  `json:"defaultBrowserType"`
}

// Playwright represents a Playwright instance
type Playwright struct {
	channelOwner
	Selectors Selectors
	Chromium  BrowserType
	Firefox   BrowserType
	WebKit    BrowserType
	Request   APIRequest
	Devices   map[string]*DeviceDescriptor
}

// Stop stops the Playwright instance
func (p *Playwright) Stop() error { _ = "STUB: not implemented"; return nil }

// Pid returns the process ID of the Playwright driver process, or 0 if not available
func (p *Playwright) Pid() int { _ = "STUB: not implemented"; return 0 }

func (p *Playwright) setSelectors(selectors Selectors) {
	_ = "STUB: not implemented"
	// Selectors has been moved to client-side only in Playwright v1.57+
	return
}

func newPlaywright(parent *channelOwner, objectType string, guid string, initializer map[string]any) *Playwright {
	_ = "STUB: not implemented"
	return nil
}

// Selectors has been moved to client-side only in Playwright v1.57+
// Only set up channel if selectors is in the initializer (older protocol)

//go:generate bash scripts/generate-api.sh
