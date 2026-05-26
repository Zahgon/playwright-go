package playwright

type frameLocatorImpl struct {
	frame         *frameImpl
	frameSelector string
}

func newFrameLocator(frame *frameImpl, frameSelector string) *frameLocatorImpl {
	_ = "STUB: not implemented"
	return nil
}

func (fl *frameLocatorImpl) First() FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (fl *frameLocatorImpl) FrameLocator(selector string) FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (fl *frameLocatorImpl) GetByAltText(text any, options ...FrameLocatorGetByAltTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) GetByLabel(text any, options ...FrameLocatorGetByLabelOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) GetByPlaceholder(text any, options ...FrameLocatorGetByPlaceholderOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) GetByRole(role AriaRole, options ...FrameLocatorGetByRoleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) GetByTestId(testId any) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) GetByText(text any, options ...FrameLocatorGetByTextOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) GetByTitle(text any, options ...FrameLocatorGetByTitleOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) Last() FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (fl *frameLocatorImpl) Locator(selectorOrLocator any, options ...FrameLocatorLocatorOptions) Locator {
	_ = "STUB: not implemented"
	return *new(Locator)
}

func (fl *frameLocatorImpl) Nth(index int) FrameLocator {
	_ = "STUB: not implemented"
	return *new(FrameLocator)
}

func (fl *frameLocatorImpl) Owner() Locator { _ = "STUB: not implemented"; return *new(Locator) }
