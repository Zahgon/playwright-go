package playwright

type serializedFallbackOverrides struct {
	URL            *string
	Method         *string
	Headers        map[string]string
	PostDataBuffer []byte
}

type requestImpl struct {
	channelOwner
	timing             *RequestTiming
	provisionalHeaders *rawHeaders
	allHeaders         *rawHeaders
	redirectedFrom     Request
	redirectedTo       Request
	failureText        string
	fallbackOverrides  *serializedFallbackOverrides
}

func (r *requestImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (r *requestImpl) ResourceType() string { _ = "STUB: not implemented"; return "" }

func (r *requestImpl) Method() string { _ = "STUB: not implemented"; return "" }

func (r *requestImpl) PostData() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *requestImpl) PostDataJSON(v any) error { _ = "STUB: not implemented"; return nil }

func (r *requestImpl) PostDataBuffer() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *requestImpl) Headers() map[string]string { _ = "STUB: not implemented"; return nil }

func (r *requestImpl) Response() (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

// no response

func (r *requestImpl) Frame() Frame { _ = "STUB: not implemented"; return *new(Frame) }

// Service Worker requests do not have an associated frame.

// Frame for this navigation request is not available, because the request
// was issued before the frame is created. You can check whether the request
// is a navigation request by calling IsNavigationRequest() method.

func (r *requestImpl) IsNavigationRequest() bool { _ = "STUB: not implemented"; return false }

func (r *requestImpl) RedirectedFrom() Request { _ = "STUB: not implemented"; return *new(Request) }

func (r *requestImpl) RedirectedTo() Request { _ = "STUB: not implemented"; return *new(Request) }

func (r *requestImpl) Failure() error { _ = "STUB: not implemented"; return nil }

func (r *requestImpl) Timing() *RequestTiming { _ = "STUB: not implemented"; return nil }

func (r *requestImpl) AllHeaders() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *requestImpl) HeadersArray() ([]NameValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *requestImpl) HeaderValue(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *requestImpl) HeaderValues(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *requestImpl) ActualHeaders() (*rawHeaders, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *requestImpl) ServiceWorker() Worker { _ = "STUB: not implemented"; return *new(Worker) }

func (r *requestImpl) Sizes() (*RequestSizesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *requestImpl) applyFallbackOverrides(options RouteFallbackOptions) {
	_ = "STUB: not implemented"
	return
}

func (r *requestImpl) targetClosed() <-chan error { _ = "STUB: not implemented"; return nil }

func (r *requestImpl) setResponseEndTiming(t float64) { _ = "STUB: not implemented"; return }

func (r *requestImpl) safePage() *pageImpl { _ = "STUB: not implemented"; return nil }

func newRequest(parent *channelOwner, objectType string, guid string, initializer map[string]any) *requestImpl {
	_ = "STUB: not implemented"
	return nil
}
