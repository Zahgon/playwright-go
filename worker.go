package playwright

type workerImpl struct {
	channelOwner
	page    *pageImpl
	context *browserContextImpl
}

func (w *workerImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (w *workerImpl) Evaluate(expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (w *workerImpl) EvaluateHandle(expression string, options ...any) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (w *workerImpl) onClose() { _ = "STUB: not implemented"; return }

func (w *workerImpl) OnClose(fn func(Worker)) { _ = "STUB: not implemented"; return }

func (w *workerImpl) OnConsole(fn func(ConsoleMessage)) { _ = "STUB: not implemented"; return }

func newWorker(parent *channelOwner, objectType string, guid string, initializer map[string]any) *workerImpl {
	_ = "STUB: not implemented"
	return nil
}
