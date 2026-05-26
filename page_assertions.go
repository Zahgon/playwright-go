package playwright

type pageAssertionsImpl struct {
	assertionsBase
	actualPage Page
}

func newPageAssertions(page Page, isNot bool, defaultTimeout *float64) *pageAssertionsImpl {
	_ = "STUB: not implemented"
	return nil
}

// expectOnFrame calls the frame's expect method directly without a selector.
// This is needed for page-level assertions like ToHaveTitle and ToHaveURL
// which should not be bound to a specific element.
func (pa *pageAssertionsImpl) expectOnFrame(
	expression string,
	options frameExpectOptions,
	expected any,
	message string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pa *pageAssertionsImpl) ToHaveTitle(titleOrRegExp any, options ...PageAssertionsToHaveTitleOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (pa *pageAssertionsImpl) ToHaveURL(urlOrRegExp any, options ...PageAssertionsToHaveURLOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (pa *pageAssertionsImpl) Not() PageAssertions {
	_ = "STUB: not implemented"
	return *new(PageAssertions)
}
