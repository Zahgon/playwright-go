package playwright

type apiRequestImpl struct {
	*Playwright
}

func (r *apiRequestImpl) NewContext(options ...APIRequestNewContextOptions) (APIRequestContext, error) {
	_ = "STUB: not implemented"
	return *new(APIRequestContext), nil
}

func newApiRequestImpl(pw *Playwright) *apiRequestImpl { _ = "STUB: not implemented"; return nil }

type apiRequestContextImpl struct {
	channelOwner
	tracing        *tracingImpl
	closeReason    *string
	defaultTimeout *float64
}

func (r *apiRequestContextImpl) Dispose(options ...APIRequestContextDisposeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *apiRequestContextImpl) Delete(url string, options ...APIRequestContextDeleteOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) Fetch(urlOrRequest any, options ...APIRequestContextFetchOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) innerFetch(url string, request Request, options ...APIRequestContextFetchOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

// only one of them can be specified

// Use context-level timeout as default if no per-request timeout specified

func (r *apiRequestContextImpl) Get(url string, options ...APIRequestContextGetOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) Head(url string, options ...APIRequestContextHeadOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) Patch(url string, options ...APIRequestContextPatchOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) Put(url string, options ...APIRequestContextPutOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) Post(url string, options ...APIRequestContextPostOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *apiRequestContextImpl) StorageState(path ...string) (*StorageState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAPIRequestContext(parent *channelOwner, objectType string, guid string, initializer map[string]any) *apiRequestContextImpl {
	_ = "STUB: not implemented"
	return nil
}

type apiResponseImpl struct {
	request     *apiRequestContextImpl
	initializer map[string]any
	headers     *rawHeaders
}

func (r *apiResponseImpl) Body() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *apiResponseImpl) Dispose() error { _ = "STUB: not implemented"; return nil }

func (r *apiResponseImpl) Headers() map[string]string { _ = "STUB: not implemented"; return nil }

func (r *apiResponseImpl) HeadersArray() []NameValue { _ = "STUB: not implemented"; return nil }

func (r *apiResponseImpl) JSON(v any) error { _ = "STUB: not implemented"; return nil }

func (r *apiResponseImpl) Ok() bool { _ = "STUB: not implemented"; return false }

func (r *apiResponseImpl) Status() int { _ = "STUB: not implemented"; return 0 }

func (r *apiResponseImpl) StatusText() string { _ = "STUB: not implemented"; return "" }

func (r *apiResponseImpl) Text() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *apiResponseImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (r *apiResponseImpl) fetchUid() string { _ = "STUB: not implemented"; return "" }

func (r *apiResponseImpl) fetchLog() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func newAPIResponse(context *apiRequestContextImpl, initializer map[string]any) *apiResponseImpl {
	_ = "STUB: not implemented"
	return nil
}

func countNonNil(args ...any) int { _ = "STUB: not implemented"; return 0 }

func isJsonContentType(headers []map[string]string) bool { _ = "STUB: not implemented"; return false }

func serializeMapToNameValue(data map[string]any) []map[string]string {
	_ = "STUB: not implemented"
	return nil
}
