package playwright

type locatorAssertionsImpl struct {
	assertionsBase
}

func newLocatorAssertions(locator Locator, isNot bool, defaultTimeout *float64) *locatorAssertionsImpl {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeAttached(options ...LocatorAssertionsToBeAttachedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeInViewport(options ...LocatorAssertionsToBeInViewportOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToContainClass(expected any, options ...LocatorAssertionsToContainClassOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToContainText(expected any, options ...LocatorAssertionsToContainTextOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveAccessibleDescription(description any, options ...LocatorAssertionsToHaveAccessibleDescriptionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveAccessibleErrorMessage(errorMessage any, options ...LocatorAssertionsToHaveAccessibleErrorMessageOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveAccessibleName(name any, options ...LocatorAssertionsToHaveAccessibleNameOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveAttribute(name string, value any, options ...LocatorAssertionsToHaveAttributeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveClass(expected any, options ...LocatorAssertionsToHaveClassOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveCSS(name string, value any, options ...LocatorAssertionsToHaveCSSOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveId(id any, options ...LocatorAssertionsToHaveIdOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveJSProperty(name string, value any, options ...LocatorAssertionsToHaveJSPropertyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveRole(role AriaRole, options ...LocatorAssertionsToHaveRoleOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveText(expected any, options ...LocatorAssertionsToHaveTextOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveValue(value any, options ...LocatorAssertionsToHaveValueOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToHaveValues(values []any, options ...LocatorAssertionsToHaveValuesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) ToMatchAriaSnapshot(expected string, options ...LocatorAssertionsToMatchAriaSnapshotOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (la *locatorAssertionsImpl) Not() LocatorAssertions {
	_ = "STUB: not implemented"
	return *new(LocatorAssertions)
}
