package playwright

type routeImpl struct {
	channelOwner
	handling *chan bool
	context  *browserContextImpl
	didThrow bool
}

func (r *routeImpl) startHandling() chan bool { _ = "STUB: not implemented"; return nil }

func (r *routeImpl) reportHandled(done bool) { _ = "STUB: not implemented"; return }

func (r *routeImpl) checkNotHandled() error { _ = "STUB: not implemented"; return nil }

func (r *routeImpl) Request() Request { _ = "STUB: not implemented"; return *new(Request) }

func unpackOptionalArgument(input any) any { _ = "STUB: not implemented"; return *new(any) }

func (r *routeImpl) Abort(errorCode ...string) error { _ = "STUB: not implemented"; return nil }

func (r *routeImpl) raceWithPageClose(f func() error) error { _ = "STUB: not implemented"; return nil }

// upstream does not throw the err

func (r *routeImpl) Fulfill(options ...RouteFulfillOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *routeImpl) handleRoute(cb func() error) error { _ = "STUB: not implemented"; return nil }

func (r *routeImpl) innerFulfill(options ...RouteFulfillOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *routeImpl) Fallback(options ...RouteFallbackOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *routeImpl) Fetch(options ...RouteFetchOptions) (APIResponse, error) {
	_ = "STUB: not implemented"
	return *new(APIResponse), nil
}

func (r *routeImpl) Continue(options ...RouteContinueOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *routeImpl) internalContinue(isFallback bool) error { _ = "STUB: not implemented"; return nil }

func (r *routeImpl) redirectedNavigationRequest(url string) error {
	_ = "STUB: not implemented"
	return nil
}

func newRoute(parent *channelOwner, objectType string, guid string, initializer map[string]any) *routeImpl {
	_ = "STUB: not implemented"
	return nil
}
