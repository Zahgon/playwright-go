package playwright

type webErrorImpl struct {
	err  error
	page Page
}

func (e *webErrorImpl) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (e *webErrorImpl) Error() error { _ = "STUB: not implemented"; return nil }

func newWebError(page Page, err error) WebError { _ = "STUB: not implemented"; return *new(WebError) }
