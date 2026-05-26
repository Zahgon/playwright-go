package playwright

type responseImpl struct {
	channelOwner
	request            *requestImpl
	provisionalHeaders *rawHeaders
	rawHeaders         *rawHeaders
	finished           chan error
}

func (r *responseImpl) FromServiceWorker() bool { _ = "STUB: not implemented"; return false }

func (r *responseImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (r *responseImpl) Ok() bool { _ = "STUB: not implemented"; return false }

func (r *responseImpl) Status() int { _ = "STUB: not implemented"; return 0 }

func (r *responseImpl) StatusText() string { _ = "STUB: not implemented"; return "" }

func (r *responseImpl) Headers() map[string]string { _ = "STUB: not implemented"; return nil }

func (r *responseImpl) Finished() error { _ = "STUB: not implemented"; return nil }

func (r *responseImpl) Body() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *responseImpl) Text() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *responseImpl) JSON(v any) error { _ = "STUB: not implemented"; return nil }

func (r *responseImpl) Request() Request { _ = "STUB: not implemented"; return *new(Request) }

func (r *responseImpl) Frame() Frame { _ = "STUB: not implemented"; return *new(Frame) }

func (r *responseImpl) AllHeaders() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *responseImpl) HeadersArray() ([]NameValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *responseImpl) HeaderValue(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *responseImpl) HeaderValues(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *responseImpl) ActualHeaders() (*rawHeaders, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *responseImpl) SecurityDetails() (*ResponseSecurityDetailsResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *responseImpl) ServerAddr() (*ResponseServerAddrResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newResponse(parent *channelOwner, objectType string, guid string, initializer map[string]any) *responseImpl {
	_ = "STUB: not implemented"
	return nil
}
