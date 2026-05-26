package playwright

type localUtilsImpl struct {
	channelOwner
	Devices map[string]*DeviceDescriptor
}

type (
	localUtilsZipOptions struct {
		ZipFile        string `json:"zipFile"`
		Entries        []any  `json:"entries"`
		StacksId       string `json:"stacksId"`
		Mode           string `json:"mode"`
		IncludeSources bool   `json:"includeSources"`
	}

	harLookupOptions struct {
		HarId               string            `json:"harId"`
		URL                 string            `json:"url"`
		Method              string            `json:"method"`
		Headers             map[string]string `json:"headers"`
		IsNavigationRequest bool              `json:"isNavigationRequest"`
		PostData            any               `json:"postData,omitempty"`
	}

	harLookupResult struct {
		Action      string              `json:"action"`
		Message     *string             `json:"message,omitempty"`
		RedirectURL *string             `json:"redirectUrl,omitempty"`
		Status      *int                `json:"status,omitempty"`
		Headers     []map[string]string `json:"headers,omitempty"`
		Body        *string             `json:"body,omitempty"`
	}
)

func (l *localUtilsImpl) Zip(options localUtilsZipOptions) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (l *localUtilsImpl) HarOpen(file string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *localUtilsImpl) HarLookup(option harLookupOptions) (*harLookupResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *localUtilsImpl) HarClose(harId string) error { _ = "STUB: not implemented"; return nil }

func (l *localUtilsImpl) HarUnzip(zipFile, harFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *localUtilsImpl) TracingStarted(traceName string, tracesDir ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (l *localUtilsImpl) TraceDiscarded(stacksId string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *localUtilsImpl) AddStackToTracingNoReply(id uint32, stack []map[string]any) {
	_ = "STUB: not implemented"
	return
}

func newLocalUtils(parent *channelOwner, objectType string, guid string, initializer map[string]any) *localUtilsImpl {
	_ = "STUB: not implemented"
	return nil
}
