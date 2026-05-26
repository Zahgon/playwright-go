package playwright

type tracingImpl struct {
	channelOwner
	includeSources bool
	isTracing      bool
	stacksId       string
	tracesDir      string
}

func (t *tracingImpl) Start(options ...TracingStartOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tracingImpl) StartChunk(options ...TracingStartChunkOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tracingImpl) StopChunk(path ...string) error { _ = "STUB: not implemented"; return nil }

func (t *tracingImpl) Stop(path ...string) error { _ = "STUB: not implemented"; return nil }

func (t *tracingImpl) doStopChunk(filePath string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Not interested in artifacts.

// Save trace to the final local file.

// The artifact may be missing if the browser closed while stopping tracing.

func (t *tracingImpl) startCollectingStacks(name string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *tracingImpl) Group(name string, options ...TracingGroupOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tracingImpl) GroupEnd() error { _ = "STUB: not implemented"; return nil }

func newTracing(parent *channelOwner, objectType string, guid string, initializer map[string]any) *tracingImpl {
	_ = "STUB: not implemented"
	return nil
}
